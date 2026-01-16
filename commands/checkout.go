package commands

import (
	"fmt"

	"github.com/pavlovic265/265-gt/components"
	helpers "github.com/pavlovic265/265-gt/helpers"
	"github.com/pavlovic265/265-gt/runner"
	"github.com/pavlovic265/265-gt/utils/log"
	"github.com/spf13/cobra"
)

type checkoutCommand struct {
	runner    runner.Runner
	gitHelper helpers.GitHelper
}

func NewCheckoutCommand(
	runner runner.Runner,
	gitHelper helpers.GitHelper,
) checkoutCommand {
	return checkoutCommand{
		runner:    runner,
		gitHelper: gitHelper,
	}
}

func (svc checkoutCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:     "checkout",
		Aliases: []string{"co"},
		Short:   "checkout branch",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := svc.gitHelper.EnsureGitRepository(); err != nil {
				return err
			}

			if len(args) > 0 {
				err := svc.checkoutBranch(args[0])
				if err != nil {
					return err
				}
			} else {
				branches, err := svc.gitHelper.GetBranches()
				if err != nil {
					return err
				}
				return svc.selectAndCheckoutBranch(branches)
			}
			return nil
		},
	}
}

func (svc checkoutCommand) checkoutBranch(
	branch string,
) error {
	if err := svc.runner.Git("checkout", branch); err != nil {
		return log.Error(fmt.Sprintf("failed to checkout branch '%s'", branch), err)
	}

	log.Success(fmt.Sprintf("Switched to branch '%s'", branch))
	return nil
}

func (svc checkoutCommand) selectAndCheckoutBranch(choices []string) error {
	if len(choices) == 0 {
		return log.ErrorMsg("no branches available to checkout")
	}

	selected, err := components.SelectString(choices)
	if err != nil {
		return log.Error("failed to display branch selection", err)
	}
	if selected == "" {
		return log.ErrorMsg("no branch selected")
	}

	return svc.checkoutBranch(selected)
}

package commands

import (
	"github.com/pavlovic265/265-gt/components"
	helpers "github.com/pavlovic265/265-gt/helpers"
	"github.com/pavlovic265/265-gt/runner"
	"github.com/pavlovic265/265-gt/utils/log"
	"github.com/spf13/cobra"
)

type upCommand struct {
	runner    runner.Runner
	gitHelper helpers.GitHelper
}

func NewUpCommand(
	runner runner.Runner,
	gitHelper helpers.GitHelper,
) upCommand {
	return upCommand{
		runner:    runner,
		gitHelper: gitHelper,
	}
}

func (svc upCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:   "up",
		Short: "move to branch up in stack",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := svc.gitHelper.EnsureGitRepository(); err != nil {
				return err
			}

			currentBranch, err := svc.gitHelper.GetCurrentBranch()
			if err != nil {
				return log.Error("failed to get current branch name", err)
			}

			children := svc.gitHelper.GetChildren(currentBranch)

			if len(children) == 0 {
				return log.ErrorMsg("cannot move up: no child branches available")
			}

			if len(children) == 1 {
				err := svc.checkoutBranch(children[0])
				if err != nil {
					return err
				}
			} else {
				return svc.selectAndCheckoutBranch(children)
			}

			return nil
		},
	}
}

func (svc upCommand) checkoutBranch(
	branch string,
) error {
	if err := svc.runner.Git("checkout", branch); err != nil {
		return log.Error("failed to checkout branch", err)
	}
	log.Successf("Moved up to branch '%s'", branch)
	return nil
}

func (svc upCommand) selectAndCheckoutBranch(choices []string) error {
	selected, err := components.SelectString(choices)
	if err != nil {
		return log.Error("failed to display branch selection menu", err)
	}
	if selected == "" {
		return log.ErrorMsg("no branch selected for checkout")
	}

	return svc.checkoutBranch(selected)
}

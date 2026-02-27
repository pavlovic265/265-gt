package branch

import (
	helpers "github.com/pavlovic265/265-gt/helpers"
	"github.com/pavlovic265/265-gt/runner"
	"github.com/pavlovic265/265-gt/ui/components"
	"github.com/pavlovic265/265-gt/utils/log"
	"github.com/spf13/cobra"
)

type moveCommand struct {
	runner    runner.Runner
	gitHelper helpers.GitHelper
}

func NewMoveCommand(
	runner runner.Runner,
	gitHelper helpers.GitHelper,
) moveCommand {
	return moveCommand{
		runner:    runner,
		gitHelper: gitHelper,
	}
}

func (svc moveCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:     "move",
		Aliases: []string{"mo"},
		Short:   "rebase branch onto other branch",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := svc.gitHelper.EnsureGitRepository(); err != nil {
				return err
			}

			if svc.gitHelper.IsRebaseInProgress() {
				return log.ErrorMsg("a rebase is already in progress; resolve it, then run `gt cont` or abort")
			}

			branch, err := svc.gitHelper.GetCurrentBranch()
			if err != nil {
				return log.Error("failed to get current branch name", err)
			}

			if len(args) > 0 {
				parent := args[0]
				if err := svc.gitHelper.RebaseBranch(branch, parent); err != nil {
					return err
				}
			} else {
				branches, err := svc.gitHelper.GetBranches()
				if err != nil {
					return log.Error("failed to get branch list", err)
				}

				if err := svc.rebaseBranch(branch, branches); err != nil {
					return err
				}
			}

			return nil
		},
	}
}

func (svc moveCommand) rebaseBranch(branch string, choices []string) error {
	selected, err := components.SelectString(choices)
	if err != nil {
		return log.Error("failed to display branch selection menu", err)
	}
	if selected == "" {
		return log.ErrorMsg("no target branch selected for rebase")
	}

	return svc.gitHelper.RebaseBranch(branch, selected)
}

package branch

import (
	helpers "github.com/pavlovic265/265-gt/helpers"
	"github.com/pavlovic265/265-gt/runner"
	"github.com/pavlovic265/265-gt/utils/log"
	"github.com/spf13/cobra"
)

type downCommand struct {
	runner    runner.Runner
	gitHelper helpers.GitHelper
}

func NewDownCommand(
	runner runner.Runner,
	gitHelper helpers.GitHelper,
) downCommand {
	return downCommand{
		runner:    runner,
		gitHelper: gitHelper,
	}
}

func (svc downCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:   "down",
		Short: "move to branch down in stack",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := svc.gitHelper.EnsureGitRepository(); err != nil {
				return err
			}

			currentBranch, err := svc.gitHelper.GetCurrentBranch()
			if err != nil {
				return log.Error("failed to get current branch name", err)
			}

			parent, err := svc.gitHelper.GetParent(currentBranch)
			if err != nil {
				return log.Error("failed to get parent branch", err)
			}

			if parent == "" {
				return log.ErrorMsg("cannot move down: no parent branch available")
			}

			if err := svc.runner.Git("checkout", parent); err != nil {
				return log.Error("failed to checkout parent branch", err)
			}

			log.Successf("Moved down to branch '%s'", parent)
			return nil
		},
	}
}

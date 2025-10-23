package commands

import (
	"github.com/pavlovic265/265-gt/executor"
	"github.com/pavlovic265/265-gt/helpers"
	"github.com/pavlovic265/265-gt/utils/log"
	pointer "github.com/pavlovic265/265-gt/utils/pointer"
	"github.com/spf13/cobra"
)

type downCommand struct {
	exe       executor.Executor
	gitHelper helpers.GitHelper
}

func NewDownCommand(
	exe executor.Executor,
	gitHelper helpers.GitHelper,
) downCommand {
	return downCommand{
		exe:       exe,
		gitHelper: gitHelper,
	}
}

func (svc downCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:   "down",
		Short: "move to brunch down in stack",
		RunE: func(cmd *cobra.Command, args []string) error {
			branch, err := svc.gitHelper.GetCurrentBranchName()
			if err != nil {
				return log.Error("Failed to get current branch name", err)
			}

			currentBranch := pointer.Deref(branch)
			parent, err := svc.gitHelper.GetParent(currentBranch)
			if err != nil {
				return log.Error("Failed to get parent branch", err)
			}

			if parent == "" {
				return log.ErrorMsg("Cannot move down - no parent branch available")
			}

			exeArgs := []string{"checkout", parent}
			err = svc.exe.WithGit().WithArgs(exeArgs).Run()
			if err != nil {
				return log.Error("Failed to checkout parent branch", err)
			}

			log.Success("Moved down to branch '" + parent + "'")
			return nil
		},
	}
}

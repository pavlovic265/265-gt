package commands

import (
	"github.com/pavlovic265/265-gt/executor"
	"github.com/pavlovic265/265-gt/helpers"
	"github.com/pavlovic265/265-gt/utils/log"
	"github.com/spf13/cobra"
)

type pullCommand struct {
	exe       executor.Executor
	gitHelper helpers.GitHelper
}

func NewPullCommand(
	exe executor.Executor,
	gitHelper helpers.GitHelper,
) pullCommand {
	return pullCommand{
		exe:       exe,
		gitHelper: gitHelper,
	}
}

func (svc pullCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:     "pull",
		Aliases: []string{"pl"},
		Short:   "pull branch",
		RunE: func(cmd *cobra.Command, args []string) error {
			currentBranchName, err := svc.gitHelper.GetCurrentBranch()
			if err != nil {
				return log.Error("Failed to get current branch name", err)
			}

			exeArgs := []string{"pull", "origin", currentBranchName}
			err = svc.exe.WithGit().WithArgs(exeArgs).Run()
			if err != nil {
				return log.Error("Failed to pull branch from remote", err)
			}

			log.Success("Branch '" + currentBranchName + "' pulled successfully")
			return nil
		},
	}
}

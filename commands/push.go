package commands

import (
	"github.com/pavlovic265/265-gt/executor"
	"github.com/pavlovic265/265-gt/helpers"
	"github.com/pavlovic265/265-gt/utils/log"
	"github.com/pavlovic265/265-gt/utils/pointer"
	"github.com/spf13/cobra"
)

type pushCommand struct {
	exe       executor.Executor
	gitHelper helpers.GitHelper
}

func NewPushCommand(
	exe executor.Executor,
	gitHelper helpers.GitHelper,
) pushCommand {
	return pushCommand{
		exe:       exe,
		gitHelper: gitHelper,
	}
}

func (svc pushCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:     "push",
		Aliases: []string{"pu"},
		Short:   "push branch always froce",
		RunE: func(cmd *cobra.Command, args []string) error {
			currentBranch, err := svc.gitHelper.GetCurrentBranchName()
			if err != nil {
				return log.Error("Failed to get current branch name", err)
			}

			currentBranchName := pointer.Deref(currentBranch)
			log.Warning("Using force push - this will overwrite remote changes")

			exeArgs := []string{"push", "--force", "origin", currentBranchName}
			err = svc.exe.WithGit().WithArgs(exeArgs).Run()
			if err != nil {
				return log.Error("Failed to push branch to remote", err)
			}

			log.Success("Branch '" + currentBranchName + "' pushed successfully")
			return nil
		},
	}
}

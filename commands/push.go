package commands

import (
	"fmt"
	"os"

	"github.com/pavlovic265/265-gt/executor"
	helpers "github.com/pavlovic265/265-gt/git_helpers"
	"github.com/pavlovic265/265-gt/utils/log"
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
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if err := svc.gitHelper.EnsureGitRepository(); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			currentBranchName, err := svc.gitHelper.GetCurrentBranch()
			if err != nil {
				return log.Error("Failed to get current branch name", err)
			}

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

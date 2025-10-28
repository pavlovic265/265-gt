package commands

import (
	"os"

	"github.com/pavlovic265/265-gt/executor"
	helpers "github.com/pavlovic265/265-gt/git_helpers"
	"github.com/pavlovic265/265-gt/utils/log"
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
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if err := svc.gitHelper.EnsureGitRepository(); err != nil {
				_ = log.Error("Not in a git repository", err)
				os.Exit(1)
			}
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			currentBranch, err := svc.gitHelper.GetCurrentBranch()
			if err != nil {
				return log.Error("Failed to get current branch name", err)
			}

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

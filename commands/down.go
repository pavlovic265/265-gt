package commands

import (
	"os"

	helpers "github.com/pavlovic265/265-gt/git_helpers"
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

			if err := svc.runner.Git("checkout", parent); err != nil {
				return log.Error("Failed to checkout parent branch", err)
			}

			log.Success("Moved down to branch '" + parent + "'")
			return nil
		},
	}
}

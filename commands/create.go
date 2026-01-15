package commands

import (
	"os"

	helpers "github.com/pavlovic265/265-gt/helpers"
	"github.com/pavlovic265/265-gt/runner"
	"github.com/pavlovic265/265-gt/utils/log"
	"github.com/spf13/cobra"
)

type createCommand struct {
	runner    runner.Runner
	gitHelper helpers.GitHelper
}

func NewCreateCommand(
	runner runner.Runner,
	gitHelper helpers.GitHelper,
) createCommand {
	return createCommand{
		runner:    runner,
		gitHelper: gitHelper,
	}
}

func (svc createCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:     "create",
		Aliases: []string{"c"},
		Short:   "create branch",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if err := svc.gitHelper.EnsureGitRepository(); err != nil {
				_ = log.Error("Not in a git repository", err)
				os.Exit(1)
			}
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return log.ErrorMsg("Branch name is required")
			}
			branch := args[0]
			parent, err := svc.gitHelper.GetCurrentBranch()
			if err != nil {
				return log.Error("Failed to get current branch name", err)
			}

			if err := svc.runner.Git("checkout", "-b", branch); err != nil {
				return log.Error("Failed to create and checkout branch", err)
			}

			if err := svc.gitHelper.SetParent(parent, branch); err != nil {
				return log.Error("Failed to set parent branch relationship", err)
			}

			log.Success("Branch '" + branch + "' created and switched to successfully")
			return nil
		},
	}
}

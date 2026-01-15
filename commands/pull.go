package commands

import (
	helpers "github.com/pavlovic265/265-gt/helpers"
	"github.com/pavlovic265/265-gt/runner"
	"github.com/pavlovic265/265-gt/utils/log"
	"github.com/spf13/cobra"
)

type pullCommand struct {
	runner    runner.Runner
	gitHelper helpers.GitHelper
}

func NewPullCommand(
	runner runner.Runner,
	gitHelper helpers.GitHelper,
) pullCommand {
	return pullCommand{
		runner:    runner,
		gitHelper: gitHelper,
	}
}

func (svc pullCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:     "pull",
		Aliases: []string{"pl"},
		Short:   "pull branch",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := svc.gitHelper.EnsureGitRepository(); err != nil {
				return err
			}

			currentBranchName, err := svc.gitHelper.GetCurrentBranch()
			if err != nil {
				return log.Error("Failed to get current branch name", err)
			}

			if err := svc.runner.Git("pull", "origin", currentBranchName); err != nil {
				return log.Error("Failed to pull branch from remote", err)
			}

			log.Successf("Branch '%s' pulled successfully", currentBranchName)
			return nil
		},
	}
}

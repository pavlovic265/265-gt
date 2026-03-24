package remote

import (
	"github.com/pavlovic265/265-gt/client"
	helpers "github.com/pavlovic265/265-gt/helpers"
	"github.com/pavlovic265/265-gt/runner"
	"github.com/pavlovic265/265-gt/utils/log"
	"github.com/spf13/cobra"
)

type pushCommand struct {
	runner    runner.Runner
	gitHelper helpers.GitHelper
	cliClient client.CliClient
}

func NewPushCommand(
	runner runner.Runner,
	gitHelper helpers.GitHelper,
	cliClient client.CliClient,
) pushCommand {
	return pushCommand{
		runner:    runner,
		gitHelper: gitHelper,
		cliClient: cliClient,
	}
}

func (svc pushCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:     "push",
		Aliases: []string{"pu"},
		Short:   "push branch always force",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := svc.gitHelper.EnsureGitRepository(); err != nil {
				return err
			}

			currentBranchName, err := svc.gitHelper.GetCurrentBranch()
			if err != nil {
				return log.Error("failed to get current branch name", err)
			}

			log.Warning("Using force push - this will overwrite remote changes")

			if err := svc.runner.Git("push", "--force", "origin", currentBranchName); err != nil {
				return log.Error("failed to push branch to remote", err)
			}

			hasOpenPR, err := svc.cliClient.HasOpenPullRequestForBranch(cmd.Context(), currentBranchName)
			if err != nil {
				return log.Error("failed to check for open pull request", err)
			}
			if hasOpenPR {
				if err := svc.cliClient.UpdatePullRequestBaseBranch(cmd.Context(), currentBranchName); err != nil {
					return log.Error("failed to update pull request base branch", err)
				}
			}

			log.Successf("Branch '%s' pushed successfully", currentBranchName)
			return nil
		},
	}
}

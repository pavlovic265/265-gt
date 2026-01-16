package commands

import (
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
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := svc.gitHelper.EnsureGitRepository(); err != nil {
				return err
			}

			if len(args) == 0 {
				return log.ErrorMsg("branch name is required")
			}
			branch := args[0]

			if err := svc.gitHelper.ValidateBranchName(branch); err != nil {
				return log.Error("invalid branch name", err)
			}
			parent, err := svc.gitHelper.GetCurrentBranch()
			if err != nil {
				return log.Error("failed to get current branch name", err)
			}

			if err := svc.runner.Git("checkout", "-b", branch); err != nil {
				return log.Error("failed to create and checkout branch", err)
			}

			if err := svc.gitHelper.SetParent(parent, branch); err != nil {
				return log.Error("failed to set parent branch relationship", err)
			}

			log.Successf("Branch '%s' created and switched to successfully", branch)
			return nil
		},
	}
}

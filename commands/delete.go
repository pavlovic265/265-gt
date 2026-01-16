package commands

import (
	"strings"

	"github.com/pavlovic265/265-gt/components"
	helpers "github.com/pavlovic265/265-gt/helpers"
	"github.com/pavlovic265/265-gt/runner"
	"github.com/pavlovic265/265-gt/utils/log"
	"github.com/spf13/cobra"
)

type deleteCommand struct {
	runner    runner.Runner
	gitHelper helpers.GitHelper
}

func NewDeleteCommand(
	runner runner.Runner,
	gitHelper helpers.GitHelper,
) deleteCommand {
	return deleteCommand{
		runner:    runner,
		gitHelper: gitHelper,
	}
}

func (svc deleteCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:                "delete",
		Aliases:            []string{"dl"},
		Short:              "delete branch",
		DisableFlagParsing: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := svc.gitHelper.EnsureGitRepository(); err != nil {
				return err
			}

			ctx := cmd.Context()
			currentBranchName, err := svc.gitHelper.GetCurrentBranch()
			if err != nil {
				return log.Error("failed to get current branch name", err)
			}

			if len(args) > 0 {
				branch := strings.TrimSpace(args[0])
				if currentBranchName == branch {
					return log.ErrorMsg("cannot delete the branch you are currently on")
				}

				if svc.gitHelper.IsProtectedBranch(ctx, branch) {
					return log.ErrorMsg("cannot delete protected branch: " + branch)
				}

				err := svc.deleteBranch(branch)
				if err != nil {
					return err
				}
			} else {
				branches, err := svc.gitHelper.GetBranches()
				if err != nil {
					return log.Error("failed to get branch list", err)
				}
				var branchesWithoutCurrent []string
				for _, branch := range branches {
					if branch != currentBranchName && !svc.gitHelper.IsProtectedBranch(ctx, branch) {
						branchesWithoutCurrent = append(branchesWithoutCurrent, branch)
					}
				}

				if len(branchesWithoutCurrent) == 0 {
					return log.ErrorMsg("no branches available for deletion")
				}

				return svc.selectAndDeleteBranch(branchesWithoutCurrent)
			}
			return nil
		},
	}
}

func (svc deleteCommand) deleteBranch(
	branch string,
) error {
	parent, err := svc.gitHelper.GetParent(branch)
	if err != nil {
		parent = ""
	}
	branchChildren := svc.gitHelper.GetChildren(branch)

	if err := svc.runner.Git("branch", "-D", branch); err != nil {
		return log.Error("failed to delete branch", err)
	}

	err = svc.gitHelper.RelinkParentChildren(parent, branchChildren)
	if err != nil {
		return log.Error("failed to update branch relationships", err)
	}

	log.Successf("Branch '%s' deleted successfully", branch)
	return nil
}

func (svc deleteCommand) selectAndDeleteBranch(choices []string) error {
	selected, err := components.SelectString(choices)
	if err != nil {
		return log.Error("failed to display branch selection menu", err)
	}
	if selected == "" {
		return log.ErrorMsg("no branch selected for deletion")
	}

	return svc.deleteBranch(selected)
}

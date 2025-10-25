package commands

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/pavlovic265/265-gt/components"
	"github.com/pavlovic265/265-gt/executor"
	helpers "github.com/pavlovic265/265-gt/git_helpers"
	"github.com/pavlovic265/265-gt/utils/log"
	"github.com/spf13/cobra"
)

type deleteCommand struct {
	exe       executor.Executor
	gitHelper helpers.GitHelper
}

func NewDeleteCommand(
	exe executor.Executor,
	gitHelper helpers.GitHelper,
) deleteCommand {
	return deleteCommand{
		exe:       exe,
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
			currentBranchName, err := svc.gitHelper.GetCurrentBranch()
			if err != nil {
				return log.Error("Failed to get current branch name", err)
			}

			if len(args) > 0 {
				branch := strings.TrimSpace(args[0])
				if currentBranchName == branch {
					return log.ErrorMsg("Cannot delete the branch you are currently on")
				}

				if svc.gitHelper.IsProtectedBranch(branch) {
					return log.ErrorMsg("Cannot delete protected branch: " + branch)
				}

				err := svc.deleteBranch(branch)
				if err != nil {
					return err
				}
			} else {
				branches, err := svc.gitHelper.GetBranches()
				if err != nil {
					return log.Error("Failed to get branch list", err)
				}
				var branchesWithoutCurrent []string
				for _, branch := range branches {
					if branch != currentBranchName && !svc.gitHelper.IsProtectedBranch(branch) {
						branchesWithoutCurrent = append(branchesWithoutCurrent, branch)
					}
				}

				if len(branchesWithoutCurrent) == 0 {
					return log.ErrorMsg("No branches available for deletion")
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
		// If we can't get parent, just set it to empty string
		parent = ""
	}
	branchChildren := svc.gitHelper.GetChildren(branch)

	exeArgs := []string{"branch", "-D", branch}
	err = svc.exe.WithGit().WithArgs(exeArgs).Run()
	if err != nil {
		return log.Error("Failed to delete branch", err)
	}

	err = svc.gitHelper.RelinkParentChildren(parent, branchChildren)
	if err != nil {
		return log.Error("Failed to update branch relationships", err)
	}

	log.Success("Branch '" + branch + "' deleted successfully")
	return nil
}

func (svc deleteCommand) selectAndDeleteBranch(
	choices []string,
) error {
	initialModel := components.ListModel{
		AllChoices: choices,
		Choices:    choices,
		Cursor:     0,
		Query:      "",
	}

	program := tea.NewProgram(initialModel)

	if finalModel, err := program.Run(); err == nil {
		if m, ok := finalModel.(components.ListModel); ok && m.Selected != "" {
			err := svc.deleteBranch(m.Selected)
			if err != nil {
				return err
			}
		} else {
			return log.ErrorMsg("No branch selected for deletion")
		}
	} else {
		return log.Error("Failed to display branch selection menu", err)
	}
	return nil
}

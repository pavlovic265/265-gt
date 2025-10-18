package commands

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/pavlovic265/265-gt/components"
	"github.com/pavlovic265/265-gt/executor"
	"github.com/pavlovic265/265-gt/helpers"
	"github.com/pavlovic265/265-gt/utils/log"
	pointer "github.com/pavlovic265/265-gt/utils/pointer"
	"github.com/spf13/cobra"
)

type moveCommand struct {
	exe       executor.Executor
	gitHelper helpers.GitHelper
}

func NewMoveCommand(
	exe executor.Executor,
	gitHelper helpers.GitHelper,
) moveCommand {
	return moveCommand{
		exe:       exe,
		gitHelper: gitHelper,
	}
}

func (svc moveCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:     "move",
		Aliases: []string{"mo"},
		Short:   "rebase branch onto other branch",
		RunE: func(cmd *cobra.Command, args []string) error {
			currentBranch, err := svc.gitHelper.GetCurrentBranchName()
			if err != nil {
				return log.Error("Failed to get current branch name", err)
			}

			currentBranchName := pointer.Deref(currentBranch)

			if len(args) > 0 {
				parentBranch := args[0]
				if err := svc.rebaseBranchOnto(parentBranch, currentBranchName); err != nil {
					return err
				}
			} else {
				branches, err := svc.gitHelper.GetBranches()
				if err != nil {
					return log.Error("Failed to get branch list", err)
				}
				if err := svc.rebaseBranch(currentBranchName, branches); err != nil {
					return err
				}
			}

			return nil
		},
	}
}

func (svc moveCommand) rebaseBranchOnto(parentBranch, currentBranch string) error {
	exeArgs := []string{"checkout", currentBranch}
	err := svc.exe.WithGit().WithArgs(exeArgs).Run()
	if err != nil {
		return log.Error("Failed to checkout current branch", err)
	}

	exeArgs = []string{"rebase", parentBranch}
	err = svc.exe.WithGit().WithArgs(exeArgs).Run()
	if err != nil {
		return log.Error("Failed to rebase branch", err)
	}

	if err := svc.gitHelper.SetParent(parentBranch, currentBranch); err != nil {
		return log.Error("Failed to set parent branch relationship", err)
	}

	if err := svc.setChildrenBranch(parentBranch, currentBranch); err != nil {
		return log.Error("Failed to update parent's children list", err)
	}

	log.Success("Branch '" + currentBranch + "' rebased onto '" + parentBranch + "' successfully")
	return nil
}

func (svc moveCommand) setChildrenBranch(parent, child string) error {
	parentChildren := svc.gitHelper.GetChildren(parent)

	var children string
	if len(parentChildren) > 0 {
		splitedParentChildren := strings.Split(parentChildren, " ")
		splitedParentChildren = append(splitedParentChildren, child)
		children = strings.Join(splitedParentChildren, " ")

	} else {
		children = child
	}

	if err := svc.gitHelper.SetChildren(children, parent); err != nil {
		return err
	}
	return nil
}

func (svc moveCommand) rebaseBranch(
	currentBranch string,
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
			if err := svc.rebaseBranchOnto(m.Selected, currentBranch); err != nil {
				return err
			}
		} else {
			return log.ErrorMsg("No target branch selected for rebase")
		}
	} else {
		return log.Error("Failed to display branch selection menu", err)
	}
	return nil
}

package commands

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/pavlovic265/265-gt/components"
	"github.com/pavlovic265/265-gt/executor"
	"github.com/pavlovic265/265-gt/utils"
	"github.com/spf13/cobra"
)

type moveCommand struct {
	exe executor.Executor
}

func NewMoveCommand(
	exe executor.Executor,
) moveCommand {
	return moveCommand{
		exe: exe,
	}
}

func (svc *moveCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:     "move",
		Aliases: []string{"mo"},
		Short:   "rebase branch onto other branch",
		RunE: func(cmd *cobra.Command, args []string) error {
			currentBranch, err := utils.GetCurrentBranchName(svc.exe)
			if err != nil {
				return err
			}

			if len(args) > 0 {
				parentBranch := args[0]
				if err := svc.rebaseBranchOnto(parentBranch, *currentBranch); err != nil {
					return err
				}
			} else {
				branchs, err := utils.GetBranches(svc.exe)
				if err != nil {
					return err
				}
				if err := svc.rebaseBranch(*currentBranch, branchs); err != nil {
					return err
				}
			}

			return nil
		},
	}
}

func (svc *moveCommand) rebaseBranchOnto(parentBranch, currentBranch string) error {
	exeArgs := []string{"checkout", parentBranch}
	if err := svc.exe.Execute("git", exeArgs...); err != nil {
		return err
	}

	exeArgs = []string{"rebase", "--onto", parentBranch, currentBranch + "~1", currentBranch}
	if err := svc.exe.Execute("git", exeArgs...); err != nil {
		return err
	}

	return nil
}

func (svc *moveCommand) rebaseBranch(
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
				return fmt.Errorf("error checking out branch: %w", err)
			}
		}
	} else {
		return fmt.Errorf("error running program: %w", err)
	}
	return nil
}

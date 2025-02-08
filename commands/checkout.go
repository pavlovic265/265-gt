package commands

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/pavlovic265/265-gt/components"
	"github.com/pavlovic265/265-gt/executor"
	"github.com/pavlovic265/265-gt/utils"
	"github.com/spf13/cobra"
)

type checkoutCommand struct {
	exe executor.Executor
}

func NewCheckoutCommand(
	exe executor.Executor,
) checkoutCommand {
	return checkoutCommand{
		exe: exe,
	}
}

func (svc checkoutCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:     "checkout",
		Aliases: []string{"co"},
		Short:   "checkout branch",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				exeArgs := []string{"checkout", args[0]}
				_, err := svc.exe.Execute("git", exeArgs...)
				if err != nil {
					return err
				}
			} else {
				branches, err := utils.GetBranches(svc.exe)
				if err != nil {
					return err
				}
				return svc.checkoutBranch(branches)
			}
			return nil
		},
	}
}

func (svc checkoutCommand) checkoutBranch(
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
			fmt.Printf("Checking out branch '%s'...\n", m.Selected)
			exeArgs := []string{"checkout", m.Selected}
			_, err := svc.exe.Execute("git", exeArgs...)
			if err != nil {
				return err
			}
		}
	} else {
		return err
	}
	return nil
}

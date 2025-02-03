package commands

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/pavlovic265/265-gt/components"
	"github.com/pavlovic265/265-gt/executor"
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

func (svc *checkoutCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:     "checkout",
		Aliases: []string{"co"},
		Short:   "checkout branch",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				exeArgs := []string{"checkout", args[0]}
				if err := svc.exe.Execute("git", exeArgs...); err != nil {
					return fmt.Errorf("error checking out branch: %w", err)
				}
			} else {
				branches, err := svc.getBranches()
				if err != nil {
					return err
				}
				return svc.checkoutBranch(branches)
			}
			return nil
		},
	}
}

func (svc *checkoutCommand) checkoutBranch(
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
			if err := svc.exe.Execute("git", exeArgs...); err != nil {
				return fmt.Errorf("error checking out branch: %w", err)
			}

		}
	} else {
		return fmt.Errorf("error running program: %w", err)
	}
	return nil
}

func (svc *checkoutCommand) getBranches() ([]string, error) {
	exeArgs := []string{"branch", "--list"}
	output, err := svc.exe.ExecuteWithOutput("git", exeArgs...)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(output), "\n")
	var branches []string
	for _, line := range lines {
		branch := strings.TrimSpace(line)
		if branch != "" {
			branches = append(branches, strings.TrimPrefix(branch, "* "))
		}
	}
	return branches, nil
}

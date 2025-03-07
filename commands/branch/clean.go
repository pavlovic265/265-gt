package branch

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/pavlovic265/265-gt/components"
	"github.com/pavlovic265/265-gt/executor"
	"github.com/pavlovic265/265-gt/utils"
	"github.com/spf13/cobra"
)

type cleanCommand struct {
	exe executor.Executor
}

func NewCleanCommand(
	exe executor.Executor,
) cleanCommand {
	return cleanCommand{
		exe: exe,
	}
}

func (svc cleanCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:                "clean",
		Aliases:            []string{"cl"},
		Short:              "git clean",
		DisableFlagParsing: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			currentBranch, err := utils.GetCurrentBranchName(svc.exe)
			if err != nil {
				return err
			}

			branches, err := utils.GetBranches(svc.exe)
			if err != nil {
				return err
			}

			for _, branch := range branches {
				if branch != *currentBranch {
					svc.deleteBranch(branch)
				}
			}

			return nil
		},
	}
}

func (svc cleanCommand) deleteBranch(
	branch string,
) error {
	initialModel := components.NewYesNoPrompt(
		fmt.Sprintf("Do you want to delete %s? (Y/n)", branch),
	)

	program := tea.NewProgram(initialModel)

	m, err := program.Run()
	if err != nil {
		return err
	}

	if model, ok := m.(components.YesNoPrompt); ok {
		if model.Quitting {
			fmt.Println("Operation canceled")
		}

		if model.IsYes() {
			exeArgs := []string{"branch", "-D", branch}
			err := svc.exe.WithGit().WithArgs(exeArgs).Run()
			if err != nil {
				return err
			}
			utils.DeleteParent(svc.exe, branch)
		}

	}

	return nil
}

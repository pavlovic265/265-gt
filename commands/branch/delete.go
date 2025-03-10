package branch

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/pavlovic265/265-gt/components"
	"github.com/pavlovic265/265-gt/executor"
	"github.com/pavlovic265/265-gt/utils"
	"github.com/spf13/cobra"
)

type deleteCommand struct {
	exe executor.Executor
}

func NewDeleteCommand(
	exe executor.Executor,
) deleteCommand {
	return deleteCommand{
		exe: exe,
	}
}

func (svc deleteCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:                "delete",
		Aliases:            []string{"dl"},
		Short:              "delete branch",
		DisableFlagParsing: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			currentBranch, err := utils.GetCurrentBranchName(svc.exe)
			if err != nil {
				return err
			}
			if len(args) > 0 {
				branch := strings.TrimSpace(args[0])
				if *currentBranch == branch {
					return fmt.Errorf("cant delete branch you are on")
				}
				err := svc.deleteBranch(branch)
				if err != nil {
					return err
				}
			} else {
				branches, err := utils.GetBranches(svc.exe)
				if err != nil {
					return err
				}
				var branchesWithoutCurrent []string
				for _, branch := range branches {
					if branch != *currentBranch {
						branchesWithoutCurrent = append(branchesWithoutCurrent, branch)
					}
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
	parent := utils.GetParent(svc.exe, branch)
	utils.RelinkParentChildren(svc.exe, parent, branch)

	exeArgs := []string{"branch", "-D", branch}
	err := svc.exe.WithGit().WithArgs(exeArgs).Run()
	if err != nil {
		return err
	}

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
		}
	} else {
		return err
	}
	return nil
}

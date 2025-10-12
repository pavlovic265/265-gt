package commands

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/pavlovic265/265-gt/components"
	"github.com/pavlovic265/265-gt/executor"
	"github.com/pavlovic265/265-gt/helpers"
	pointer "github.com/pavlovic265/265-gt/utils/pointer"
	"github.com/spf13/cobra"
)

type upCommand struct {
	exe       executor.Executor
	gitHelper helpers.GitHelper
}

func NewUpCommand(
	exe executor.Executor,
	gitHelper helpers.GitHelper,
) upCommand {
	return upCommand{
		exe:       exe,
		gitHelper: gitHelper,
	}
}

func (svc upCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:   "up",
		Short: "move to brunch up in stack",
		RunE: func(cmd *cobra.Command, args []string) error {
			branch, err := svc.gitHelper.GetCurrentBranchName(svc.exe)
			if err != nil {
				return err
			}
			childrenStr := svc.gitHelper.GetChildren(svc.exe, pointer.Deref(branch))
			children := svc.gitHelper.UnmarshalChildren(childrenStr)

			if len(children) == 1 {
				err := svc.checkoutBranch(children[0])
				if err != nil {
					return err
				}
			} else {
				return svc.selectAndCheckoutBranch(children)
			}

			return nil
		},
	}
}

func (svc upCommand) checkoutBranch(
	branch string,
) error {
	exeArgs := []string{"checkout", branch}
	err := svc.exe.WithGit().WithArgs(exeArgs).Run()
	if err != nil {
		return err
	}
	fmt.Println("Moved up to branch '" + branch + "'")
	return nil
}

func (svc upCommand) selectAndCheckoutBranch(
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
			err := svc.checkoutBranch(m.Selected)
			if err != nil {
				return err
			}
		}
	} else {
		return err
	}
	return nil
}

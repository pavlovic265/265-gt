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
		Short:              "clean branches one by one",
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
					shouldBrake, err := svc.deleteBranch(branch)
					if err != nil {
						fmt.Println(err.Error())
					}
					if shouldBrake {
						break
					}
				}
			}

			return nil
		},
	}
}

func (svc cleanCommand) deleteBranch(
	branch string,
) (bool, error) {
	initialModel := components.NewYesNoPrompt(
		fmt.Sprintf("Do you want to delete %s? (Y/n)", branch),
	)

	program := tea.NewProgram(initialModel)

	m, err := program.Run()
	if err != nil {
		return false, err
	}

	if model, ok := m.(components.YesNoPrompt); ok {
		if model.Quitting {
			fmt.Println("Operation canceled")
			return true, nil
		}

		if model.IsYes() {
			parent := utils.GetParent(svc.exe, branch)
			exeArgs := []string{"branch", "-D", branch}
			err := svc.exe.WithGit().WithArgs(exeArgs).Run()
			if err != nil {
				return false, err
			}

			svc.cleanBranchs(parent, branch)
		}

	}

	return false, nil
}

func (svc cleanCommand) cleanBranchs(parent, branch string) {
	fmt.Println("0 ", parent)
	children := utils.GetChildren(svc.exe, parent)
	fmt.Println("1 ", children)
	splitChldren := strings.Split(children, " ")
	fmt.Println("2 ", splitChldren)

	if len(splitChldren) > 0 {
		fmt.Println("3 ", len(splitChldren))
		var newChildren []string

		for _, child := range splitChldren {
			fmt.Println("4 ", child, branch)
			if child != branch {
				newChildren = append(newChildren, child)
			}
		}
		fmt.Println("4.5 ", newChildren)
		if len(newChildren) > 0 {

			joinChildren := strings.TrimSpace(strings.Join(newChildren, " "))
			fmt.Println("5 ", joinChildren)

			utils.SetChildren(svc.exe, parent, joinChildren)
		} else {
			utils.DeleteChildren(svc.exe, branch)
		}
		return
	}
	utils.DeleteParent(svc.exe, branch)
}

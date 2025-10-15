package commands

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/pavlovic265/265-gt/components"
	"github.com/pavlovic265/265-gt/constants"
	"github.com/pavlovic265/265-gt/executor"
	"github.com/pavlovic265/265-gt/helpers"
	pointer "github.com/pavlovic265/265-gt/utils/pointer"
	"github.com/spf13/cobra"
)

type cleanCommand struct {
	exe       executor.Executor
	gitHelper helpers.GitHelper
}

// Styling definitions for clean command
var (
	// Header styles
	headerStyle = lipgloss.NewStyle().
			Foreground(constants.Blue).
			Bold(true)

	// Info styles
	infoStyle = lipgloss.NewStyle().
			Foreground(constants.Cyan)

	// Success styles
	successStyle = lipgloss.NewStyle().
			Foreground(constants.White)

	successIconStyle = lipgloss.NewStyle().
				Foreground(constants.Green)

	// Error styles
	errorStyle = lipgloss.NewStyle().
			Foreground(constants.Red)

	// Branch info styles
	branchStyle = lipgloss.NewStyle().
			Foreground(constants.Magenta)

	parentStyle = lipgloss.NewStyle().
			Foreground(constants.BrightBlack)
)

func NewCleanCommand(
	exe executor.Executor,
	gitHelper helpers.GitHelper,
) cleanCommand {
	return cleanCommand{
		exe:       exe,
		gitHelper: gitHelper,
	}
}

func (svc cleanCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:     "clean",
		Aliases: []string{"cl"},
		Short:   "Clean up branches interactively",
		Long:    "Clean up branches one by one with confirmation. Protected branches and current branch are skipped.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return svc.cleanBranches()
		},
	}
}

func (svc cleanCommand) cleanBranches() error {
	currentBranch, err := svc.gitHelper.GetCurrentBranchName()
	if err != nil {
		return fmt.Errorf("failed to get current branch: %w", err)
	}

	branches, err := svc.gitHelper.GetBranches()
	if err != nil {
		return fmt.Errorf("failed to get branches: %w", err)
	}

	// Styled header
	fmt.Println(headerStyle.Render("Branch Cleanup"))
	fmt.Println()

	cleanableCount := 0
	for _, branch := range branches {
		if branch != pointer.Deref(currentBranch) && !svc.gitHelper.IsProtectedBranch(branch) {
			cleanableCount++
		}
	}

	if cleanableCount == 0 {
		fmt.Println(infoStyle.Render("No branches to clean up!"))
		return nil
	}

	deletedCount := 0
	for _, branch := range branches {
		if branch == pointer.Deref(currentBranch) || svc.gitHelper.IsProtectedBranch(branch) {
			continue
		}

		shouldBreak, err := svc.deleteBranch(branch)
		if err != nil {
			fmt.Printf("%s %s\n",
				constants.ErrorIcon,
				errorStyle.Render(fmt.Sprintf("Error: %v", err)))
			continue
		}

		if shouldBreak {
			break
		}

		deletedCount++
	}

	if deletedCount > 0 {
		fmt.Printf("%s %s",
			successIconStyle.Render(constants.SuccessIcon),
			successStyle.Render(fmt.Sprintf("Cleaned up %d branches", deletedCount)))
	}
	return nil
}

func (svc cleanCommand) deleteBranch(branch string) (bool, error) {
	parent := svc.gitHelper.GetParent(branch)

	// Create styled prompt message
	var promptMsg strings.Builder
	promptMsg.WriteString("Delete branch ")
	promptMsg.WriteString(branchStyle.Render("'" + branch + "'"))
	promptMsg.WriteString("?")

	if parent != "" {
		promptMsg.WriteString(" (")
		promptMsg.WriteString(parentStyle.Render("parent: "))
		promptMsg.WriteString(branchStyle.Render(parent))
		promptMsg.WriteString(")")
	}

	initialModel := components.NewYesNoPrompt(promptMsg.String())
	program := tea.NewProgram(initialModel)

	m, err := program.Run()
	if err != nil {
		return false, err
	}

	if model, ok := m.(components.YesNoPrompt); ok {
		if model.Quitting {
			return true, nil
		}

		if model.IsYes() {
			parentChildren := svc.gitHelper.GetChildren(parent)
			branchChildren := svc.gitHelper.GetChildren(branch)

			exeArgs := []string{"branch", "-D", branch}
			output, err := svc.exe.WithGit().WithArgs(exeArgs).RunWithOutput()
			if err != nil {
				return false, err
			}

			err = svc.gitHelper.RelinkParentChildren(parent, parentChildren, branch, branchChildren)
			if err != nil {
				return false, err
			}

			gitOutput := strings.TrimSpace(output.String())
			fmt.Printf("   %s %s\n\n",
				successIconStyle.Render(constants.SuccessIcon),
				successStyle.Render(gitOutput))
		}
	}

	return false, nil
}

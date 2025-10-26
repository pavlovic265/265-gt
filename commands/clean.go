package commands

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/pavlovic265/265-gt/components"
	"github.com/pavlovic265/265-gt/constants"
	"github.com/pavlovic265/265-gt/executor"
	helpers "github.com/pavlovic265/265-gt/git_helpers"
	"github.com/pavlovic265/265-gt/utils/log"
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
	currentBranch, err := svc.gitHelper.GetCurrentBranch()
	if err != nil {
		return log.Error("Failed to get current branch", err)
	}

	branches, err := svc.gitHelper.GetBranches()
	if err != nil {
		return log.Error("Failed to get branches", err)
	}

	// Styled header
	fmt.Println(headerStyle.Render("Branch Cleanup"))
	log.Infof("Current branch: %s", branchStyle.Render(currentBranch))
	fmt.Println()

	cleanableCount := 0
	protectedCount := 0
	for _, branch := range branches {
		if branch == currentBranch {
			continue
		}
		if svc.gitHelper.IsProtectedBranch(branch) {
			protectedCount++
			continue
		}
		cleanableCount++
	}

	log.Infof("Found %d branches (%d protected, %d cleanable)",
		len(branches)-1, // -1 for current branch
		protectedCount,
		cleanableCount)
	fmt.Println()

	if cleanableCount == 0 {
		log.Info("No branches to clean up!")
		return nil
	}

	deletedCount := 0
	skippedCount := 0
	for _, branch := range branches {
		if branch == currentBranch || svc.gitHelper.IsProtectedBranch(branch) {
			continue
		}

		shouldBreak, deleted, err := svc.deleteBranch(branch)
		if err != nil {
			_ = log.Errorf("Failed to delete branch: %v", err)
			continue
		}

		if shouldBreak {
			break
		}

		if deleted {
			deletedCount++
		} else {
			skippedCount++
		}
	}

	// Summary
	fmt.Println()
	if deletedCount > 0 {
		log.Successf("Deleted %d branches", deletedCount)
	}
	if skippedCount > 0 {
		log.Infof("Skipped %d branches", skippedCount)
	}
	return nil
}

func (svc cleanCommand) deleteBranch(branch string) (shouldBreak bool, deleted bool, err error) {
	parent, err := svc.gitHelper.GetParent(branch)
	if err != nil {
		// If we can't get parent, just set it to empty string
		parent = ""
	}

	branchChildren := svc.gitHelper.GetChildren(branch)

	// Create styled prompt message
	var promptMsg strings.Builder
	promptMsg.WriteString("Delete branch ")
	promptMsg.WriteString(branchStyle.Render("'" + branch + "'"))
	promptMsg.WriteString("?")

	if parent != "" {
		promptMsg.WriteString(" ")
		promptMsg.WriteString(parentStyle.Render("(parent: " + parent + ")"))
	}

	if len(branchChildren) > 0 {
		promptMsg.WriteString(" ")
		promptMsg.WriteString(infoStyle.Render(fmt.Sprintf("(children: %d)", len(branchChildren))))
	}

	initialModel := components.NewYesNoPrompt(promptMsg.String())
	program := tea.NewProgram(initialModel)

	m, err := program.Run()
	if err != nil {
		return false, false, err
	}

	if model, ok := m.(components.YesNoPrompt); ok {
		if model.Quitting {
			return true, false, nil
		}

		if model.IsYes() {
			// Delete the branch
			exeArgs := []string{"branch", "-D", branch}
			output, err := svc.exe.WithGit().WithArgs(exeArgs).RunWithOutput()
			if err != nil {
				return false, false, err
			}

			// Relink children to parent
			if len(branchChildren) > 0 {
				err = svc.gitHelper.RelinkParentChildren(parent, branchChildren)
				if err != nil {
					return false, false, err
				}
				fmt.Printf("   → Relinked %d children to %s\n",
					len(branchChildren),
					branchStyle.Render(parent))
			}

			gitOutput := strings.TrimSpace(output.String())
			fmt.Printf("   ")
			log.Success(gitOutput)
			fmt.Println()

			return false, true, nil
		}
	}

	return false, false, nil
}

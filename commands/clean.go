package commands

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/pavlovic265/265-gt/components"
	"github.com/pavlovic265/265-gt/config"
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
	currentBranch, err := utils.GetCurrentBranchName(svc.exe)
	if err != nil {
		return fmt.Errorf("failed to get current branch: %w", err)
	}

	branches, err := utils.GetBranches(svc.exe)
	if err != nil {
		return fmt.Errorf("failed to get branches: %w", err)
	}

	fmt.Println(config.TitleStyle.Render("🧹 Branch Cleanup"))
	fmt.Println()

	cleanableCount := 0
	for _, branch := range branches {
		if branch != *currentBranch && !utils.IsProtectedBranch(branch) {
			cleanableCount++
		}
	}

	if cleanableCount == 0 {
		fmt.Println(config.SuccessStyle.Render("✨ No branches to clean up!"))
		return nil
	}

	// Show options once at the top
			fmt.Printf("   %s [Y] Yes  [N] No  [Ctrl+Q] Cancel\n", config.InfoStyle.Render("Options:"))
		fmt.Printf("   %s Default: Yes (press Enter)\n", config.HighlightStyle.Render("💡"))
	fmt.Println()

	deletedCount := 0
	for _, branch := range branches {
		if branch == *currentBranch || utils.IsProtectedBranch(branch) {
			continue
		}

		shouldBreak, err := svc.deleteBranch(branch)
		if err != nil {
			fmt.Printf("%s Error: %v\n", config.ErrorStyle.Render("❌"), err)
			continue
		}
		
		if shouldBreak {
			break
		}
		
		deletedCount++
	}

	fmt.Printf("\n%s Cleaned up %d branches\n", config.SuccessStyle.Render("✨"), deletedCount)
	return nil
}

func (svc cleanCommand) deleteBranch(branch string) (bool, error) {
	parent := utils.GetParent(svc.exe, branch)
	
	promptMsg := fmt.Sprintf("🗑️  Delete branch '%s'?", config.InfoStyle.Render(branch))
	if parent != "" {
		promptMsg += fmt.Sprintf(" (parent: %s)", config.DebugStyle.Render(parent))
	}

	initialModel := components.NewYesNoPrompt(promptMsg)
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
			parentChildren := utils.GetChildren(svc.exe, parent)
			branchChildren := utils.GetChildren(svc.exe, branch)

			exeArgs := []string{"branch", "-D", branch}
			output, err := svc.exe.WithGit().WithArgs(exeArgs).RunWithOutput()
			if err != nil {
				return false, err
			}

			err = utils.RelinkParentChildren(svc.exe, parent, parentChildren, branch, branchChildren)
			if err != nil {
				return false, err
			}

			gitOutput := strings.TrimSpace(output.String())
			fmt.Printf("   %s %s\n", config.SuccessStyle.Render("✅"), gitOutput)
		}
	}

	return false, nil
}

package commands

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/pavlovic265/265-gt/components"
	"github.com/pavlovic265/265-gt/executor"
	"github.com/pavlovic265/265-gt/helpers"
	pointer "github.com/pavlovic265/265-gt/utils/pointer"
	"github.com/spf13/cobra"
)

type cleanCommand struct {
	exe       executor.Executor
	gitHelper helpers.GitHelper
}

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
	currentBranch, err := svc.gitHelper.GetCurrentBranchName(svc.exe)
	if err != nil {
		return fmt.Errorf("failed to get current branch: %w", err)
	}

	branches, err := svc.gitHelper.GetBranches(svc.exe)
	if err != nil {
		return fmt.Errorf("failed to get branches: %w", err)
	}

	fmt.Println("🧹 Branch Cleanup")
	fmt.Println()

	cleanableCount := 0
	for _, branch := range branches {
		if branch != pointer.Deref(currentBranch) && !svc.gitHelper.IsProtectedBranch(branch) {
			cleanableCount++
		}
	}

	if cleanableCount == 0 {
		fmt.Println("No branches to clean up!")
		return nil
	}

	// Show options once at the top
	fmt.Printf("   %s [Y] Yes  [N] No  [Ctrl+Q] Cancel\n", "Options:")
	fmt.Printf("   %s Default: Yes (press Enter)\n", "💡")
	fmt.Println()

	deletedCount := 0
	for _, branch := range branches {
		if branch == pointer.Deref(currentBranch) || svc.gitHelper.IsProtectedBranch(branch) {
			continue
		}

		shouldBreak, err := svc.deleteBranch(branch)
		if err != nil {
			fmt.Printf("✗ Error: %v\n", err)
			continue
		}

		if shouldBreak {
			break
		}

		deletedCount++
	}

	fmt.Printf("\n✓ Cleaned up %d branches\n", deletedCount)
	return nil
}

func (svc cleanCommand) deleteBranch(branch string) (bool, error) {
	parent := svc.gitHelper.GetParent(svc.exe, branch)

	promptMsg := fmt.Sprintf("🗑️  Delete branch '%s'?", branch)
	if parent != "" {
		promptMsg += fmt.Sprintf(" (parent: %s)", parent)
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
			parentChildren := svc.gitHelper.GetChildren(svc.exe, parent)
			branchChildren := svc.gitHelper.GetChildren(svc.exe, branch)

			exeArgs := []string{"branch", "-D", branch}
			output, err := svc.exe.WithGit().WithArgs(exeArgs).RunWithOutput()
			if err != nil {
				return false, err
			}

			err = svc.gitHelper.RelinkParentChildren(svc.exe, parent, parentChildren, branch, branchChildren)
			if err != nil {
				return false, err
			}

			gitOutput := strings.TrimSpace(output.String())
			fmt.Printf("   ✓ %s\n", gitOutput)
		}
	}

	return false, nil
}

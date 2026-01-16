package commands

import (
	"context"
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/pavlovic265/265-gt/components"
	"github.com/pavlovic265/265-gt/constants"
	helpers "github.com/pavlovic265/265-gt/helpers"
	"github.com/pavlovic265/265-gt/runner"
	"github.com/pavlovic265/265-gt/utils/log"
	"github.com/spf13/cobra"
)

type cleanCommand struct {
	runner    runner.Runner
	gitHelper helpers.GitHelper
}

var (
	headerStyle = lipgloss.NewStyle().
			Foreground(constants.Blue).
			Bold(true)

	infoStyle = lipgloss.NewStyle().
			Foreground(constants.Cyan)

	branchStyle = lipgloss.NewStyle().
			Foreground(constants.Magenta)

	parentStyle = lipgloss.NewStyle().
			Foreground(constants.BrightBlack)
)

func NewCleanCommand(
	runner runner.Runner,
	gitHelper helpers.GitHelper,
) cleanCommand {
	return cleanCommand{
		runner:    runner,
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
			if err := svc.gitHelper.EnsureGitRepository(); err != nil {
				return err
			}

			return svc.cleanBranches(cmd.Context())
		},
	}
}

func (svc cleanCommand) cleanBranches(ctx context.Context) error {
	currentBranch, err := svc.gitHelper.GetCurrentBranch()
	if err != nil {
		return log.Error("failed to get current branch", err)
	}

	branches, err := svc.gitHelper.GetBranches()
	if err != nil {
		return log.Error("failed to get branches", err)
	}

	fmt.Println(headerStyle.Render("Branch Cleanup"))
	log.Infof("Current branch: %s", branchStyle.Render(currentBranch))
	fmt.Println()

	cleanableCount := 0
	protectedCount := 0
	for _, branch := range branches {
		if branch == currentBranch {
			continue
		}
		if svc.gitHelper.IsProtectedBranch(ctx, branch) {
			protectedCount++
			continue
		}
		cleanableCount++
	}

	log.Infof("Found %d branches (%d protected, %d cleanable)",
		len(branches)-1,
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
		if branch == currentBranch || svc.gitHelper.IsProtectedBranch(ctx, branch) {
			continue
		}

		shouldBreak, deleted, err := svc.deleteBranch(branch)
		if err != nil {
			_ = log.Errorf("failed to delete branch: %v", err)
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
		parent = ""
	}
	branchChildren := svc.gitHelper.GetChildren(branch)

	promptMsg := svc.buildDeletePrompt(branch, parent, branchChildren)
	model, err := svc.showDeletePrompt(promptMsg)
	if err != nil {
		return false, false, err
	}

	if model.Quitting {
		return true, false, nil
	}

	if model.IsYes() {
		if err := svc.performDelete(branch, parent, branchChildren); err != nil {
			return false, false, err
		}
		return false, true, nil
	}

	return false, false, nil
}

func (svc cleanCommand) buildDeletePrompt(branch, parent string, children []string) string {
	var msg strings.Builder
	msg.WriteString("Delete branch ")
	msg.WriteString(branchStyle.Render("'" + branch + "'"))
	msg.WriteString("?")

	if parent != "" {
		msg.WriteString(" ")
		msg.WriteString(parentStyle.Render("(parent: " + parent + ")"))
	}

	if len(children) > 0 {
		msg.WriteString(" ")
		msg.WriteString(infoStyle.Render(fmt.Sprintf("(children: %d)", len(children))))
	}

	return msg.String()
}

func (svc cleanCommand) showDeletePrompt(promptMsg string) (components.YesNoPrompt, error) {
	initialModel := components.NewYesNoPrompt(promptMsg)
	program := tea.NewProgram(initialModel)

	m, err := program.Run()
	if err != nil {
		return components.YesNoPrompt{}, err
	}

	if model, ok := m.(components.YesNoPrompt); ok {
		return model, nil
	}

	return components.YesNoPrompt{}, nil
}

func (svc cleanCommand) performDelete(branch, parent string, children []string) error {
	output, err := svc.runner.GitOutput("branch", "-D", branch)
	if err != nil {
		return err
	}

	if len(children) > 0 {
		if err := svc.gitHelper.RelinkParentChildren(parent, children); err != nil {
			return err
		}
		fmt.Printf("   → Relinked %d children to %s\n", len(children), branchStyle.Render(parent))
	}

	fmt.Printf("   ")
	log.Success(output)
	fmt.Println()

	return nil
}

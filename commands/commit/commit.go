package commit

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/pavlovic265/265-gt/constants"
	"github.com/pavlovic265/265-gt/executor"
	timeutils "github.com/pavlovic265/265-gt/utils/timeutils"
	"github.com/spf13/cobra"
)

// Styling definitions for commit command
var (
	// Success styles
	successIconStyle = lipgloss.NewStyle().
				Foreground(constants.Green)

	// Error styles
	errorIconStyle = lipgloss.NewStyle().
			Foreground(constants.Red)

	// Message styles
	messageStyle = lipgloss.NewStyle().
			Foreground(constants.White)
)

type commitCommand struct {
	exe executor.Executor
}

func NewCommitCommand(
	exe executor.Executor,
) commitCommand {
	return commitCommand{
		exe: exe,
	}
}

func (svc commitCommand) Command() *cobra.Command {
	var empty bool

	commitCmd := &cobra.Command{
		Use:     "commit",
		Aliases: []string{"cm"},
		Short:   "create commit",
		RunE: func(cmd *cobra.Command, args []string) error {
			if empty {
				return svc.handleEmptyCommit()
			}

			if len(args) == 0 {
				return fmt.Errorf("%s %s",
					errorIconStyle.Render(constants.ErrorIcon),
					messageStyle.Render("No commit message provided"))
			}
			message := string(args[0])
			return svc.handleCommit(message)

		},
	}

	commitCmd.Flags().BoolVarP(&empty, "empty", "e", false, "create an empty commit")

	return commitCmd
}

func (svc commitCommand) handleEmptyCommit() error {
	message := timeutils.Now().Format(timeutils.LayoutUserFriendly)
	exeArgs := []string{"commit", "--allow-empty", "-m", message}
	err := svc.exe.WithGit().WithArgs(exeArgs).Run()
	if err != nil {
		return fmt.Errorf("%s %s",
			errorIconStyle.Render(constants.ErrorIcon),
			messageStyle.Render(fmt.Sprintf("Failed to create empty commit: %v", err)))
	}

	fmt.Printf("%s %s\n",
		successIconStyle.Render(constants.SuccessIcon),
		messageStyle.Render(fmt.Sprintf("Empty commit created '%s'", message)))
	return nil
}

func (svc commitCommand) handleCommit(message string) error {
	exeArgs := []string{"commit", "-m", message}
	err := svc.exe.WithGit().WithArgs(exeArgs).Run()
	if err != nil {
		return fmt.Errorf("%s %s",
			errorIconStyle.Render(constants.ErrorIcon),
			messageStyle.Render(fmt.Sprintf("Failed to create commit: %v", err)))
	}

	fmt.Printf("%s %s\n",
		successIconStyle.Render(constants.SuccessIcon),
		messageStyle.Render(fmt.Sprintf("Commit created '%s'", message)))
	return nil
}

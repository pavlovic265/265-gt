package commands

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/pavlovic265/265-gt/constants"
	"github.com/pavlovic265/265-gt/executor"
	"github.com/spf13/cobra"
)

// Styling definitions for unstage command
var (
	// Success styles
	unstageSuccessIconStyle = lipgloss.NewStyle().
				Foreground(constants.Green)

	// Error styles
	unstageErrorIconStyle = lipgloss.NewStyle().
				Foreground(constants.Red)

	// Message styles
	unstageMessageStyle = lipgloss.NewStyle().
				Foreground(constants.White)
)

type unstageCommand struct {
	exe executor.Executor
}

func NewUnstageCommand(
	exe executor.Executor,
) unstageCommand {
	return unstageCommand{
		exe: exe,
	}
}

func (svc unstageCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:     "unstage",
		Aliases: []string{"us"},
		Short:   "unstage ",
		RunE: func(cmd *cobra.Command, args []string) error {
			exeArgs := append([]string{"restore", "--staged"}, args...)
			err := svc.exe.WithGit().WithArgs(exeArgs).Run()
			if err != nil {
				return fmt.Errorf("%s %s",
					unstageErrorIconStyle.Render(constants.ErrorIcon),
					unstageMessageStyle.Render(fmt.Sprintf("Failed to unstage files: %v", err)))
			}

			if len(args) == 0 {
				fmt.Printf("%s %s\n",
					unstageSuccessIconStyle.Render(constants.SuccessIcon),
					unstageMessageStyle.Render("All staged changes unstaged"))
			} else {
				fmt.Printf("%s %s\n",
					unstageSuccessIconStyle.Render(constants.SuccessIcon),
					unstageMessageStyle.Render("Files unstaged successfully"))
			}
			return nil
		},
	}
}

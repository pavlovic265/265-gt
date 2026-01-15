package commands

import (
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/pavlovic265/265-gt/constants"
	helpers "github.com/pavlovic265/265-gt/helpers"
	"github.com/pavlovic265/265-gt/runner"
	"github.com/pavlovic265/265-gt/utils/log"
	"github.com/spf13/cobra"
)

var (
	unstageSuccessIconStyle = lipgloss.NewStyle().
				Foreground(constants.Green)

	unstageErrorIconStyle = lipgloss.NewStyle().
				Foreground(constants.Red)

	unstageMessageStyle = lipgloss.NewStyle().
				Foreground(constants.White)
)

type unstageCommand struct {
	runner    runner.Runner
	gitHelper helpers.GitHelper
}

func NewUnstageCommand(
	runner runner.Runner,
	gitHelper helpers.GitHelper,
) unstageCommand {
	return unstageCommand{
		runner:    runner,
		gitHelper: gitHelper,
	}
}

func (svc unstageCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:     "unstage",
		Aliases: []string{"us"},
		Short:   "unstage ",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if err := svc.gitHelper.EnsureGitRepository(); err != nil {
				_ = log.Error("Not in a git repository", err)
				os.Exit(1)
			}
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			gitArgs := append([]string{"restore", "--staged"}, args...)
			if err := svc.runner.Git(gitArgs...); err != nil {
				return fmt.Errorf("%s %s",
					unstageErrorIconStyle.Render(constants.CrossIcon),
					unstageMessageStyle.Render(fmt.Sprintf("Failed to unstage files: %v", err)))
			}

			if len(args) == 0 {
				fmt.Printf("%s %s\n",
					unstageSuccessIconStyle.Render(constants.CheckIcon),
					unstageMessageStyle.Render("All staged changes unstaged"))
			} else {
				fmt.Printf("%s %s\n",
					unstageSuccessIconStyle.Render(constants.CheckIcon),
					unstageMessageStyle.Render("Files unstaged successfully"))
			}
			return nil
		},
	}
}

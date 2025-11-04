package commands

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/pavlovic265/265-gt/components"
	"github.com/pavlovic265/265-gt/executor"
	helpers "github.com/pavlovic265/265-gt/git_helpers"
	"github.com/pavlovic265/265-gt/utils/log"
	"github.com/spf13/cobra"
)

type trackCommand struct {
	exe       executor.Executor
	gitHelper helpers.GitHelper
}

func NewTrackCommand(
	exe executor.Executor,
	gitHelper helpers.GitHelper,
) trackCommand {
	return trackCommand{
		exe:       exe,
		gitHelper: gitHelper,
	}
}

func (svc trackCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:     "track",
		Aliases: []string{"tr"},
		Short:   "track existing branch",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if err := svc.gitHelper.EnsureGitRepository(); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			branchs, err := svc.gitHelper.GetBranches()
			if err != nil {
				return log.Error("Failed to get branches", err)
			}

			initialModel := components.ListModel[string]{
				AllChoices: branchs,
				Choices:    branchs,
				Cursor:     0,
				Query:      "",
				Formatter:  func(s string) string { return s },
				Matcher:    func(s, query string) bool { return strings.Contains(s, query) },
			}

			program := tea.NewProgram(initialModel)

			if finalModel, err := program.Run(); err == nil {
				if m, ok := finalModel.(components.ListModel[string]); ok && m.Selected != "" {
					branch, err := svc.gitHelper.GetCurrentBranch()
					if err != nil {
						return log.Error("Failed to get current branch name", err)
					}

					if err := svc.gitHelper.RebaseBranch(branch, m.Selected); err != nil {
						return err
					}
				} else {
					// User cancelled or no selection made
					return log.ErrorMsg("No branch selected")
				}
			} else {
				return log.Error("Failed to display branch selection", err)
			}

			log.Success("tracked to previous branch")
			return nil
		},
	}
}

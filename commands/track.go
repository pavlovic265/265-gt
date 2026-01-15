package commands

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/pavlovic265/265-gt/components"
	helpers "github.com/pavlovic265/265-gt/git_helpers"
	"github.com/pavlovic265/265-gt/runner"
	"github.com/pavlovic265/265-gt/utils/log"
	"github.com/spf13/cobra"
)

type trackCommand struct {
	runner    runner.Runner
	gitHelper helpers.GitHelper
}

func NewTrackCommand(
	runner runner.Runner,
	gitHelper helpers.GitHelper,
) trackCommand {
	return trackCommand{
		runner:    runner,
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
			branch, err := svc.gitHelper.GetCurrentBranch()
			if err != nil {
				return log.Error("Failed to get current branch name", err)
			}

			branchs, err := svc.gitHelper.GetBranches()
			if err != nil {
				return log.Error("Failed to get branches", err)
			}
			branchesWithoutCurrent := make([]string, len(branchs)-2)
			for _, b := range branchs {
				if b != branch {
					branchesWithoutCurrent = append(branchesWithoutCurrent, b)
				}
			}

			initialModel := components.ListModel[string]{
				AllChoices: branchesWithoutCurrent,
				Choices:    branchesWithoutCurrent,
				Cursor:     0,
				Query:      "",
				Formatter:  func(s string) string { return s },
				Matcher:    func(s, query string) bool { return strings.Contains(s, query) },
			}

			program := tea.NewProgram(initialModel)

			if finalModel, err := program.Run(); err == nil {
				if m, ok := finalModel.(components.ListModel[string]); ok && m.Selected != "" {
					if err := svc.gitHelper.SetParent(m.Selected, branch); err != nil {
						return log.Error("Faild to set parent", err)
					}
				} else {
					// User cancelled or no selection made
					return log.ErrorMsg("No branch selected")
				}
			} else {
				return log.Error("Failed to display branch selection", err)
			}

			log.Successf("successfully tracking %s", branch)
			return nil
		},
	}
}

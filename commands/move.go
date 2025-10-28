package commands

import (
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/pavlovic265/265-gt/components"
	"github.com/pavlovic265/265-gt/executor"
	helpers "github.com/pavlovic265/265-gt/git_helpers"
	"github.com/pavlovic265/265-gt/utils/log"
	"github.com/spf13/cobra"
)

type moveCommand struct {
	exe       executor.Executor
	gitHelper helpers.GitHelper
}

func NewMoveCommand(
	exe executor.Executor,
	gitHelper helpers.GitHelper,
) moveCommand {
	return moveCommand{
		exe:       exe,
		gitHelper: gitHelper,
	}
}

func (svc moveCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:     "move",
		Aliases: []string{"mo"},
		Short:   "rebase branch onto other branch",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if err := svc.gitHelper.EnsureGitRepository(); err != nil {
				_ = log.Error("Not in a git repository", err)
				os.Exit(1)
			}
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if svc.gitHelper.IsRebaseInProgress() {
				return log.ErrorMsg("A rebase is already in progress. Resolve it, then run `gt cont` or abort.")
			}

			branch, err := svc.gitHelper.GetCurrentBranch()
			if err != nil {
				return log.Error("Failed to get current branch name", err)
			}

			if len(args) > 0 {
				parent := args[0]
				if err := svc.gitHelper.RebaseBranch(branch, parent); err != nil {
					return err
				}
			} else {
				branches, err := svc.gitHelper.GetBranches()
				if err != nil {
					return log.Error("Failed to get branch list", err)
				}
				if err := svc.rebaseBranch(branch, branches); err != nil {
					return err
				}
			}

			return nil
		},
	}
}

func (svc moveCommand) rebaseBranch(
	branch string,
	choices []string,
) error {
	initialModel := components.ListModel{
		AllChoices: choices,
		Choices:    choices,
		Cursor:     0,
		Query:      "",
	}

	program := tea.NewProgram(initialModel)

	if finalModel, err := program.Run(); err == nil {
		if m, ok := finalModel.(components.ListModel); ok && m.Selected != "" {
			if err := svc.gitHelper.RebaseBranch(branch, m.Selected); err != nil {
				return err
			}
		} else {
			return log.ErrorMsg("No target branch selected for rebase")
		}
	} else {
		return log.Error("Failed to display branch selection menu", err)
	}
	return nil
}

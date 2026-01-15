package commands

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/pavlovic265/265-gt/components"
	helpers "github.com/pavlovic265/265-gt/helpers"
	"github.com/pavlovic265/265-gt/runner"
	"github.com/pavlovic265/265-gt/utils/log"
	"github.com/spf13/cobra"
)

type upCommand struct {
	runner    runner.Runner
	gitHelper helpers.GitHelper
}

func NewUpCommand(
	runner runner.Runner,
	gitHelper helpers.GitHelper,
) upCommand {
	return upCommand{
		runner:    runner,
		gitHelper: gitHelper,
	}
}

func (svc upCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:   "up",
		Short: "move to branch up in stack",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := svc.gitHelper.EnsureGitRepository(); err != nil {
				return err
			}

			currentBranch, err := svc.gitHelper.GetCurrentBranch()
			if err != nil {
				return log.Error("Failed to get current branch name", err)
			}

			children := svc.gitHelper.GetChildren(currentBranch)

			if len(children) == 0 {
				return log.ErrorMsg("Cannot move up - no child branches available")
			}

			if len(children) == 1 {
				err := svc.checkoutBranch(children[0])
				if err != nil {
					return err
				}
			} else {
				return svc.selectAndCheckoutBranch(children)
			}

			return nil
		},
	}
}

func (svc upCommand) checkoutBranch(
	branch string,
) error {
	if err := svc.runner.Git("checkout", branch); err != nil {
		return log.Error("Failed to checkout branch", err)
	}
	log.Success("Moved up to branch '" + branch + "'")
	return nil
}

func (svc upCommand) selectAndCheckoutBranch(
	choices []string,
) error {
	initialModel := components.ListModel[string]{
		AllChoices: choices,
		Choices:    choices,
		Cursor:     0,
		Query:      "",
		Formatter:  func(s string) string { return s },
		Matcher:    func(s, query string) bool { return strings.Contains(s, query) },
	}

	program := tea.NewProgram(initialModel)

	if finalModel, err := program.Run(); err == nil {
		if m, ok := finalModel.(components.ListModel[string]); ok && m.Selected != "" {
			err := svc.checkoutBranch(m.Selected)
			if err != nil {
				return err
			}
		} else {
			return log.ErrorMsg("No branch selected for checkout")
		}
	} else {
		return log.Error("Failed to display branch selection menu", err)
	}
	return nil
}

package pullrequests

import (
	"fmt"
	"os/exec"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/pavlovic265/265-gt/client"
	"github.com/pavlovic265/265-gt/components"
	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/constants"
	"github.com/pavlovic265/265-gt/executor"
	"github.com/pavlovic265/265-gt/utils/log"
	"github.com/spf13/cobra"
)

type listCommand struct {
	exe           executor.Executor
	configManager config.ConfigManager
}

func NewListCommand(
	exe executor.Executor,
	configManager config.ConfigManager,
) listCommand {
	return listCommand{
		exe:           exe,
		configManager: configManager,
	}
}

func (svc listCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:     "list",
		Short:   "show list of pull request",
		Aliases: []string{"li"},
		RunE: func(cmd *cobra.Command, args []string) error {
			if !svc.configManager.HasActiveAccount() {
				return log.ErrorMsg("No active account found")
			}
			account := svc.configManager.GetActiveAccount()

			prs, err := client.Client[account.Platform].ListPullRequests(args)
			if err != nil {
				return log.Error("Failed to list pull requests", err)
			}
			return svc.selectPullRequest(prs)
		},
	}
}

type pullRequest struct {
	number int
	title  string
	url    string
}

func (svc listCommand) selectPullRequest(
	prs []client.PullRequest,
) error {
	var pullRequests []pullRequest
	for _, pr := range prs {
		ciStatus := ""
		ciStatusColor := constants.White
		switch pr.StatusState {
		case "SUCCESS":
			ciStatus = "✓ "
			ciStatusColor = constants.Green
		case "FAILURE", "ERROR":
			ciStatus = "✗ "
			ciStatusColor = constants.Red
		case "PENDING", "IN_PROGRESS":
			ciStatus = "* "
			ciStatusColor = constants.Yellow
		}

		// Mergeable status indicator
		mergeableStatus := " ●"
		mergeableColor := constants.Yellow
		switch pr.Mergeable {
		case "MERGEABLE":
			mergeableStatus = " ●"
			mergeableColor = constants.Green
		case "CONFLICTING":
			mergeableStatus = " ●"
			mergeableColor = constants.Red
		}

		// Style each component
		styledCiStatus := lipgloss.NewStyle().Foreground(ciStatusColor).Render(ciStatus)
		styledNumber := lipgloss.NewStyle().Foreground(constants.White).Render(fmt.Sprintf("%d", pr.Number))
		styledTitle := lipgloss.NewStyle().Foreground(constants.White).Render(pr.Title)
		_ = lipgloss.NewStyle().Foreground(mergeableColor).Render(mergeableStatus)

		pullRequests = append(pullRequests, pullRequest{
			number: pr.Number,
			title:  fmt.Sprintf("%s%s: %s", styledCiStatus, styledNumber, styledTitle),
			url:    pr.URL,
		})
	}

	initialCursor := 0

	initialModel := components.ListModel[pullRequest]{
		AllChoices:    pullRequests,
		Choices:       pullRequests,
		Cursor:        initialCursor,
		Query:         "",
		EnableYank:    true,
		EnableMerge:   true,
		EnableUpdate:  true,
		EnableRefresh: true,
		Formatter:     func(pr pullRequest) string { return pr.title },
		Matcher:       func(pr pullRequest, query string) bool { return strings.Contains(pr.title, query) },
	}

	program := tea.NewProgram(initialModel)

	if finalModel, err := program.Run(); err == nil {
		if m, ok := finalModel.(components.ListModel[pullRequest]); ok {
			if m.RefreshAction {
				// Refresh the PR list
				account := svc.configManager.GetActiveAccount()
				updatedPrs, err := client.Client[account.Platform].ListPullRequests([]string{})
				if err != nil {
					return log.Error("Failed to refresh pull requests", err)
				}
				log.Info("Pull requests refreshed")
				return svc.selectPullRequest(updatedPrs)
			}

			if m.YankAction {
				svc.yankToClipboard(m.Selected.url)
				log.Success("URL yanked to clipboard: " + m.Selected.url)
				return nil
			}

			if m.MergeAction {
				account := svc.configManager.GetActiveAccount()
				err := client.Client[account.Platform].MergePullRequest(m.Selected.number)
				if err != nil {
					return log.Error("Failed to merge pull request", err)
				}
				log.Success(fmt.Sprintf("Successfully merged PR #%d", m.Selected.number))
				return nil
			}

			if m.UpdateAction {
				account := svc.configManager.GetActiveAccount()
				err := client.Client[account.Platform].UpdatePullRequestBranch(m.Selected.number)
				if err != nil {
					return log.Error("Failed to update pull request branch", err)
				}
				log.Success(fmt.Sprintf("Successfully updated PR #%d branch", m.Selected.number))

				// Refresh the PR list
				updatedPrs, err := client.Client[account.Platform].ListPullRequests([]string{})
				if err != nil {
					return log.Error("Failed to refresh pull requests", err)
				}
				return svc.selectPullRequest(updatedPrs)
			}

			for _, pr := range prs {
				if m.Selected.number == pr.Number {
					_ = exec.Command("open", pr.URL).Start() // Ignore errors when opening URL
				}
			}
		}
	} else {
		return log.Error("Failed to display pull request selection menu", err)
	}
	return nil
}

func (svc listCommand) yankToClipboard(url string) {
	commands := [][]string{
		{"pbcopy"},                           // macOS
		{"xclip", "-selection", "clipboard"}, // Linux with xclip
		{"xsel", "--clipboard", "--input"},   // Linux with xsel
		{"clip"},                             // Windows
	}

	for _, cmd := range commands {
		command := exec.Command(cmd[0], cmd[1:]...)
		command.Stdin = strings.NewReader(url)
		if err := command.Run(); err == nil {
			return
		}
	}
}

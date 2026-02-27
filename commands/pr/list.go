package pr

import (
	"context"
	"fmt"
	"os/exec"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/pavlovic265/265-gt/client"
	"github.com/pavlovic265/265-gt/config"
	helpers "github.com/pavlovic265/265-gt/helpers"
	"github.com/pavlovic265/265-gt/runner"
	"github.com/pavlovic265/265-gt/ui/components"
	"github.com/pavlovic265/265-gt/ui/theme"
	"github.com/pavlovic265/265-gt/utils/log"
	"github.com/spf13/cobra"
)

type listCommand struct {
	runner        runner.Runner
	configManager config.ConfigManager
	gitHelper     helpers.GitHelper
	cliClient     client.CliClient
	ctx           context.Context
}

func NewListCommand(
	runner runner.Runner,
	configManager config.ConfigManager,
	gitHelper helpers.GitHelper,
	cliClient client.CliClient,
) *listCommand {
	return &listCommand{
		runner:        runner,
		configManager: configManager,
		gitHelper:     gitHelper,
		cliClient:     cliClient,
	}
}

func (svc *listCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:     "list",
		Short:   "show list of pull requests",
		Aliases: []string{"li"},
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := svc.gitHelper.EnsureGitRepository(); err != nil {
				return err
			}

			cfg, err := config.RequireGlobal(cmd.Context())
			if err != nil {
				return err
			}

			if cfg.Global.ActiveAccount == nil || cfg.Global.ActiveAccount.User == "" {
				return log.ErrorMsg("no active account found")
			}
			svc.ctx = cmd.Context()

			prs, err := svc.cliClient.ListPullRequests(svc.ctx, args)
			if err != nil {
				return log.Error("failed to list pull requests", err)
			}
			return svc.selectPullRequest(prs)
		},
	}
}

type PullRequestItem struct {
	Number int
	Title  string
	URL    string
	Branch string
}

func (svc *listCommand) FormatPullRequest(pr client.PullRequest) PullRequestItem {
	ciStatus := ""
	ciStatusColor := theme.White
	switch pr.StatusState {
	case client.StatusStateTypeSuccess:
		ciStatus = "✓ "
		ciStatusColor = theme.Green
	case client.StatusStateTypeFailure:
		ciStatus = "✗ "
		ciStatusColor = theme.Red
	case client.StatusStateTypePending:
		ciStatus = "* "
		ciStatusColor = theme.Yellow
	}

	// Review status indicator
	reviewStatus := " ●"
	reviewColor := theme.Yellow
	switch pr.ReviewState {
	case client.ReviewStateApproved:
		reviewColor = theme.Green
	case client.ReviewStateChangesRequested:
		reviewColor = theme.Red
	}

	// Conflict indicator (only shown when conflicting)
	conflictStatus := ""
	if pr.Mergeable == "CONFLICTING" {
		conflictStatus = " ⚠"
	}

	// Merge queue indicator
	mergeQueueStatus := ""
	if pr.MergeQueued {
		mergeQueueStatus = " ⧗"
	}

	// Style each component
	styledCiStatus := lipgloss.NewStyle().Foreground(ciStatusColor).Render(ciStatus)
	styledNumber := lipgloss.NewStyle().Foreground(theme.White).Render(fmt.Sprintf("%d", pr.Number))
	styledTitle := lipgloss.NewStyle().Foreground(theme.White).Render(pr.Title)
	styledReview := lipgloss.NewStyle().Foreground(reviewColor).Render(reviewStatus)
	styledConflict := lipgloss.NewStyle().Foreground(theme.Red).Render(conflictStatus)
	styledMergeQueue := lipgloss.NewStyle().Foreground(theme.BrightYellow).Render(mergeQueueStatus)

	return PullRequestItem{
		Number: pr.Number,
		Title: fmt.Sprintf("%s%s: %s%s%s%s",
			styledCiStatus, styledNumber, styledTitle, styledReview, styledConflict, styledMergeQueue),
		URL:    pr.URL,
		Branch: pr.Branch,
	}
}

func (svc *listCommand) refreshFunc() tea.Msg {
	updatedPrs, err := svc.cliClient.ListPullRequests(svc.ctx, []string{})
	if err != nil {
		return components.RefreshCompleteMsg[PullRequestItem]{Err: err}
	}

	var refreshedPullRequests []PullRequestItem
	for _, pr := range updatedPrs {
		refreshedPullRequests = append(refreshedPullRequests, svc.FormatPullRequest(pr))
	}

	return components.RefreshCompleteMsg[PullRequestItem]{
		Choices: refreshedPullRequests,
		Err:     nil,
	}
}

func (svc *listCommand) selectPullRequest(
	prs []client.PullRequest,
	selectedPRNumber ...int,
) error {
	var pullRequestItems []PullRequestItem
	for _, pr := range prs {
		pullRequestItems = append(pullRequestItems, svc.FormatPullRequest(pr))
	}

	initialCursor := 0
	// If a previously selected PR number is provided, find its index in the new list
	if len(selectedPRNumber) > 0 {
		for i, pr := range pullRequestItems {
			if pr.Number == selectedPRNumber[0] {
				initialCursor = i
				break
			}
		}
	}

	initialModel := components.ListModel[PullRequestItem]{
		AllChoices:    pullRequestItems,
		Choices:       pullRequestItems,
		Cursor:        initialCursor,
		Query:         "",
		EnableYank:    true,
		EnableMerge:   true,
		EnableRefresh: true,
		Formatter:     func(pr PullRequestItem) string { return pr.Title },
		Matcher:       func(pr PullRequestItem, query string) bool { return strings.Contains(pr.Title, query) },
		RefreshFunc:   svc.refreshFunc,
	}

	program := tea.NewProgram(initialModel)

	if finalModel, err := program.Run(); err == nil {
		if m, ok := finalModel.(components.ListModel[PullRequestItem]); ok {
			if m.YankAction {
				svc.yankToClipboard(m.Selected.URL)
				log.Successf("URL yanked to clipboard: %s", m.Selected.URL)
				return nil
			}

			if m.MergeAction {
				err := svc.cliClient.MergePullRequest(svc.ctx, m.Selected.Number)
				if err != nil {
					return log.Error("failed to merge pull request", err)
				}
				err = svc.gitHelper.DeleteParent(m.Selected.Branch)
				if err != nil {
					return log.Error("failed to delete parent connection", err)
				}
				log.Success(fmt.Sprintf("Successfully merged PR #%d", m.Selected.Number))
				return nil
			}

			for _, pr := range prs {
				if m.Selected.Number == pr.Number {
					_ = exec.Command("open", pr.URL).Start() // Ignore errors when opening URL
				}
			}
		}
	} else {
		return log.Error("failed to display pull request selection menu", err)
	}
	return nil
}

func (svc *listCommand) yankToClipboard(url string) {
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

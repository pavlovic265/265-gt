package pullrequests

import (
	"context"
	"fmt"
	"os/exec"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/pavlovic265/265-gt/client"
	"github.com/pavlovic265/265-gt/components"
	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/constants"
	helpers "github.com/pavlovic265/265-gt/git_helpers"
	"github.com/pavlovic265/265-gt/runner"
	"github.com/pavlovic265/265-gt/utils/log"
	"github.com/spf13/cobra"
)

type listCommand struct {
	runner        runner.Runner
	configManager config.ConfigManager
	gitHelper     helpers.GitHelper
	account       *config.Account
	ctx           context.Context
}

func NewListCommand(
	runner runner.Runner,
	configManager config.ConfigManager,
	gitHelper helpers.GitHelper,
) *listCommand {
	return &listCommand{
		runner:        runner,
		configManager: configManager,
		gitHelper:     gitHelper,
	}
}

func (svc *listCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:     "list",
		Short:   "show list of pull request",
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
				return log.ErrorMsg("No active account found")
			}
			svc.account = cfg.Global.ActiveAccount
			svc.ctx = cmd.Context()

			prs, err := client.Client[svc.account.Platform].ListPullRequests(svc.ctx, args)
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
	branch string
}

func (svc *listCommand) formatPullRequest(pr client.PullRequest) pullRequest {
	ciStatus := ""
	ciStatusColor := constants.White
	switch pr.StatusState {
	case "SUCCESS":
		ciStatus = "✓ "
		ciStatusColor = constants.Green
	case "FAILURE":
		ciStatus = "✗ "
		ciStatusColor = constants.Red
	case "PENDING":
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

	return pullRequest{
		number: pr.Number,
		title:  fmt.Sprintf("%s%s: %s", styledCiStatus, styledNumber, styledTitle),
		url:    pr.URL,
		branch: pr.Branch,
	}
}

func (svc *listCommand) refreshFunc() tea.Msg {
	updatedPrs, err := client.Client[svc.account.Platform].ListPullRequests(svc.ctx, []string{})
	if err != nil {
		return components.RefreshCompleteMsg[pullRequest]{Err: err}
	}

	var refreshedPullRequests []pullRequest
	for _, pr := range updatedPrs {
		refreshedPullRequests = append(refreshedPullRequests, svc.formatPullRequest(pr))
	}

	return components.RefreshCompleteMsg[pullRequest]{
		Choices: refreshedPullRequests,
		Err:     nil,
	}
}

func (svc *listCommand) selectPullRequest(
	prs []client.PullRequest,
	selectedPRNumber ...int,
) error {
	var pullRequests []pullRequest
	for _, pr := range prs {
		pullRequests = append(pullRequests, svc.formatPullRequest(pr))
	}

	initialCursor := 0
	// If a previously selected PR number is provided, find its index in the new list
	if len(selectedPRNumber) > 0 {
		for i, pr := range pullRequests {
			if pr.number == selectedPRNumber[0] {
				initialCursor = i
				break
			}
		}
	}

	initialModel := components.ListModel[pullRequest]{
		AllChoices:    pullRequests,
		Choices:       pullRequests,
		Cursor:        initialCursor,
		Query:         "",
		EnableYank:    true,
		EnableMerge:   true,
		EnableRefresh: true,
		Formatter:     func(pr pullRequest) string { return pr.title },
		Matcher:       func(pr pullRequest, query string) bool { return strings.Contains(pr.title, query) },
		RefreshFunc:   svc.refreshFunc,
	}

	program := tea.NewProgram(initialModel)

	if finalModel, err := program.Run(); err == nil {
		if m, ok := finalModel.(components.ListModel[pullRequest]); ok {
			if m.YankAction {
				svc.yankToClipboard(m.Selected.url)
				log.Success("URL yanked to clipboard: " + m.Selected.url)
				return nil
			}

			if m.MergeAction {
				err := client.Client[svc.account.Platform].MergePullRequest(svc.ctx, m.Selected.number)
				if err != nil {
					return log.Error("Failed to merge pull request", err)
				}
				err = svc.gitHelper.DeleteParent(m.Selected.branch)
				if err != nil {
					return log.Error("Failed to delete parent connection", err)
				}
				log.Success(fmt.Sprintf("Successfully merged PR #%d", m.Selected.number))
				return nil
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

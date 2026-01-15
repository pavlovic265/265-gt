package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/pavlovic265/265-gt/config"
	helpers "github.com/pavlovic265/265-gt/git_helpers"
	"github.com/pavlovic265/265-gt/utils/pointer"
)

const githubAPIBase = "https://api.github.com"

var githubHTTPClient = &http.Client{Timeout: 30 * time.Second}

type gitHubClient struct {
	gitHelper helpers.GitHelper
}

func NewGitHubClient(gitHelper helpers.GitHelper) CliClient {
	return &gitHubClient{gitHelper: gitHelper}
}

func (c *gitHubClient) getRepoInfo(ctx context.Context) (*RepoInfo, *config.Account, error) {
	cfg, ok := config.GetConfig(ctx)
	if !ok {
		return nil, nil, fmt.Errorf("config not loaded")
	}

	if cfg.Global.ActiveAccount == nil {
		return nil, nil, fmt.Errorf("no active account")
	}

	remoteURL, err := c.gitHelper.GetRemoteURL("origin")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get remote URL: %w", err)
	}

	repoInfo, err := ParseRemoteURL(remoteURL)
	if err != nil {
		return nil, nil, err
	}

	return repoInfo, cfg.Global.ActiveAccount, nil
}

func (c *gitHubClient) doRequest(
	ctx context.Context, method, url string, body any, token string,
) (*http.Response, error) {
	var reqBody *bytes.Buffer
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		reqBody = bytes.NewBuffer(jsonBody)
	} else {
		reqBody = bytes.NewBuffer(nil)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, reqBody)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	return githubHTTPClient.Do(req)
}

func (c *gitHubClient) AuthStatus(ctx context.Context) error {
	cfg, ok := config.GetConfig(ctx)
	if !ok {
		return fmt.Errorf("config not loaded")
	}

	account := cfg.Global.ActiveAccount
	if account == nil {
		return fmt.Errorf("no active account")
	}

	resp, err := c.doRequest(ctx, "GET", githubAPIBase+"/user", nil, account.Token)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("authentication failed: %s", resp.Status)
	}

	var user struct {
		Login string `json:"login"`
		Name  string `json:"name"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return err
	}

	fmt.Printf("Logged in as %s (%s)\n", user.Login, account.Platform)
	if user.Name != "" {
		fmt.Printf("Name: %s\n", user.Name)
	}

	return nil
}

func (c *gitHubClient) AuthLogin(ctx context.Context, user string) error {
	cfg, ok := config.GetConfig(ctx)
	if !ok {
		return fmt.Errorf("config not loaded")
	}

	for _, acc := range cfg.Global.Accounts {
		if acc.User == user {
			cfg.Global.ActiveAccount = pointer.From(acc)
			cfg.MarkDirty()
			return nil
		}
	}

	return fmt.Errorf("account not found: %s", user)
}

func (c *gitHubClient) AuthLogout(ctx context.Context, user string) error {
	cfg, ok := config.GetConfig(ctx)
	if !ok {
		return fmt.Errorf("config not loaded")
	}

	cfg.Global.ActiveAccount = nil
	cfg.MarkDirty()
	return nil
}

func (c *gitHubClient) CreatePullRequest(ctx context.Context, args []string) error {
	repoInfo, account, err := c.getRepoInfo(ctx)
	if err != nil {
		return err
	}

	branch, err := c.gitHelper.GetCurrentBranch()
	if err != nil {
		return err
	}

	parent, err := c.gitHelper.GetParent(branch)
	if err != nil {
		return err
	}

	isDraft := false
	for _, arg := range args {
		if arg == "--draft" || arg == "-d" {
			isDraft = true
		}
	}

	// Get last commit message for PR title
	title := branch

	payload := map[string]any{
		"title": title,
		"head":  branch,
		"base":  parent,
		"draft": isDraft,
	}

	url := fmt.Sprintf("%s/repos/%s/%s/pulls", githubAPIBase, repoInfo.Owner, repoInfo.Repo)
	resp, err := c.doRequest(ctx, "POST", url, payload, account.Token)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 201 {
		var errResp struct {
			Message string `json:"message"`
			Errors  []struct {
				Message string `json:"message"`
			} `json:"errors"`
		}
		_ = json.NewDecoder(resp.Body).Decode(&errResp)
		if len(errResp.Errors) > 0 {
			return fmt.Errorf("failed to create PR: %s", errResp.Errors[0].Message)
		}
		return fmt.Errorf("failed to create PR: %s", errResp.Message)
	}

	var pr struct {
		HTMLURL string `json:"html_url"`
		Number  int    `json:"number"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&pr); err != nil {
		return err
	}

	fmt.Printf("Created PR #%d: %s\n", pr.Number, pr.HTMLURL)

	// Add assignee
	assigneeURL := fmt.Sprintf(
		"%s/repos/%s/%s/issues/%d/assignees", githubAPIBase, repoInfo.Owner, repoInfo.Repo, pr.Number)
	assigneePayload := map[string]any{
		"assignees": []string{account.User},
	}
	assignResp, err := c.doRequest(ctx, "POST", assigneeURL, assigneePayload, account.Token)
	if err == nil {
		assignResp.Body.Close()
	}

	return nil
}

func (c *gitHubClient) ListPullRequests(ctx context.Context, args []string) ([]PullRequest, error) {
	repoInfo, account, err := c.getRepoInfo(ctx)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/repos/%s/%s/pulls?state=open", githubAPIBase, repoInfo.Owner, repoInfo.Repo)
	resp, err := c.doRequest(ctx, "GET", url, nil, account.Token)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("failed to list PRs: %s", resp.Status)
	}

	var ghPRs []struct {
		Number  int    `json:"number"`
		Title   string `json:"title"`
		HTMLURL string `json:"html_url"`
		User    struct {
			Login string `json:"login"`
		} `json:"user"`
		Head struct {
			Ref string `json:"ref"`
			SHA string `json:"sha"`
		} `json:"head"`
		Mergeable *bool `json:"mergeable"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&ghPRs); err != nil {
		return nil, err
	}

	var prs []PullRequest
	for _, pr := range ghPRs {
		if pr.User.Login != account.User {
			continue
		}

		mergeable := "UNKNOWN"
		if pr.Mergeable != nil {
			if *pr.Mergeable {
				mergeable = "MERGEABLE"
			} else {
				mergeable = "CONFLICTING"
			}
		}

		statusState := c.getCheckRunStatus(ctx, repoInfo, account.Token, pr.Head.SHA)

		prs = append(prs, PullRequest{
			Number:      pr.Number,
			Title:       pr.Title,
			URL:         pr.HTMLURL,
			Author:      pr.User.Login,
			Branch:      pr.Head.Ref,
			Mergeable:   mergeable,
			StatusState: statusState,
		})
	}

	return prs, nil
}

func (c *gitHubClient) MergePullRequest(ctx context.Context, prNumber int) error {
	repoInfo, account, err := c.getRepoInfo(ctx)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%s/repos/%s/%s/pulls/%d/merge", githubAPIBase, repoInfo.Owner, repoInfo.Repo, prNumber)
	resp, err := c.doRequest(ctx, "PUT", url, map[string]string{}, account.Token)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		var errResp struct {
			Message string `json:"message"`
		}
		_ = json.NewDecoder(resp.Body).Decode(&errResp)
		return fmt.Errorf("failed to merge PR: %s", errResp.Message)
	}

	return nil
}

func (c *gitHubClient) UpdatePullRequestBranch(ctx context.Context, prNumber int) error {
	repoInfo, account, err := c.getRepoInfo(ctx)
	if err != nil {
		return err
	}

	payload := map[string]string{
		"body": "@dependabot rebase",
	}

	url := fmt.Sprintf("%s/repos/%s/%s/issues/%d/comments", githubAPIBase, repoInfo.Owner, repoInfo.Repo, prNumber)
	resp, err := c.doRequest(ctx, "POST", url, payload, account.Token)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 201 {
		return fmt.Errorf("failed to add comment: status %d", resp.StatusCode)
	}

	return nil
}

func (c *gitHubClient) getCheckRunStatus(ctx context.Context, repoInfo *RepoInfo, token, sha string) StatusStateType {
	url := fmt.Sprintf("%s/repos/%s/%s/commits/%s/check-runs", githubAPIBase, repoInfo.Owner, repoInfo.Repo, sha)

	resp, err := c.doRequest(ctx, "GET", url, nil, token)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return ""
	}

	var result struct {
		CheckRuns []struct {
			Status     string  `json:"status"`
			Conclusion *string `json:"conclusion"`
		} `json:"check_runs"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return ""
	}

	hasFailure := false
	hasPending := false
	hasSuccess := false

	for _, check := range result.CheckRuns {
		if check.Status != "completed" {
			hasPending = true
			continue
		}
		if check.Conclusion != nil {
			switch *check.Conclusion {
			case "success", "skipped":
				hasSuccess = true
			case "failure", "timed_out", "cancelled":
				hasFailure = true
			}
		}
	}

	if hasFailure {
		return StatusStateTypeFailure
	} else if hasPending {
		return StatusStateTypePending
	} else if hasSuccess {
		return StatusStateTypeSuccess
	}
	return ""
}

type PullRequest struct {
	Number      int             `json:"number"`
	Title       string          `json:"title"`
	URL         string          `json:"url"`
	Author      string          `json:"author"`
	Mergeable   string          `json:"mergeable"`
	Branch      string          `json:"headRefName"`
	StatusState StatusStateType `json:"statusState"`
}

type StatusStateType string

const (
	StatusStateTypeSuccess StatusStateType = "SUCCESS"
	StatusStateTypeFailure StatusStateType = "FAILURE"
	StatusStateTypePending StatusStateType = "PENDING"
)

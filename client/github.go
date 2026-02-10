package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/pavlovic265/265-gt/config"
	helpers "github.com/pavlovic265/265-gt/helpers"
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
		return nil, nil, ErrConfigNotLoaded
	}

	if cfg.Global.ActiveAccount == nil {
		return nil, nil, ErrNoActiveAccount
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
		return ErrConfigNotLoaded
	}

	account := cfg.Global.ActiveAccount
	if account == nil {
		return ErrNoActiveAccount
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
		return ErrConfigNotLoaded
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
		return ErrConfigNotLoaded
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
		body, _ := io.ReadAll(resp.Body)
		var errResp struct {
			Message string `json:"message"`
			Errors  []struct {
				Resource string `json:"resource"`
				Code     string `json:"code"`
				Field    string `json:"field"`
				Message  string `json:"message"`
			} `json:"errors"`
		}
		_ = json.Unmarshal(body, &errResp)
		if len(errResp.Errors) > 0 {
			e := errResp.Errors[0]
			if e.Message != "" {
				return fmt.Errorf("failed to create PR: %s", e.Message)
			}
			return fmt.Errorf("failed to create PR: %s %s (%s)", e.Resource, e.Field, e.Code)
		}
		if errResp.Message != "" {
			return fmt.Errorf("failed to create PR: %s", errResp.Message)
		}
		return fmt.Errorf("failed to create PR (status %d): %s", resp.StatusCode, string(body))
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
		reviewState := c.getReviewState(ctx, repoInfo, account.Token, pr.Number)

		prs = append(prs, PullRequest{
			Number:      pr.Number,
			Title:       pr.Title,
			URL:         pr.HTMLURL,
			Author:      pr.User.Login,
			Branch:      pr.Head.Ref,
			Mergeable:   mergeable,
			StatusState: statusState,
			ReviewState: reviewState,
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

func (c *gitHubClient) getReviewState(ctx context.Context, repoInfo *RepoInfo, token string, prNumber int) ReviewStateType {
	url := fmt.Sprintf("%s/repos/%s/%s/pulls/%d/reviews", githubAPIBase, repoInfo.Owner, repoInfo.Repo, prNumber)

	resp, err := c.doRequest(ctx, "GET", url, nil, token)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return ""
	}

	var reviews []struct {
		User struct {
			Login string `json:"login"`
		} `json:"user"`
		State string `json:"state"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&reviews); err != nil {
		return ""
	}

	// Track the latest review per user
	latestByUser := make(map[string]string)
	for _, r := range reviews {
		if r.State == "APPROVED" || r.State == "CHANGES_REQUESTED" {
			latestByUser[r.User.Login] = r.State
		}
	}

	hasChangesRequested := false
	hasApproved := false
	for _, state := range latestByUser {
		switch state {
		case "CHANGES_REQUESTED":
			hasChangesRequested = true
		case "APPROVED":
			hasApproved = true
		}
	}

	if hasChangesRequested {
		return ReviewStateChangesRequested
	}
	if hasApproved {
		return ReviewStateApproved
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
	ReviewState ReviewStateType `json:"reviewState"`
}

type StatusStateType string

const (
	StatusStateTypeSuccess StatusStateType = "SUCCESS"
	StatusStateTypeFailure StatusStateType = "FAILURE"
	StatusStateTypePending StatusStateType = "PENDING"
)

type ReviewStateType string

const (
	ReviewStateApproved         ReviewStateType = "APPROVED"
	ReviewStateChangesRequested ReviewStateType = "CHANGES_REQUESTED"
)

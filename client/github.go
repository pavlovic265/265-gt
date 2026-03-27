package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/constants"
	helpers "github.com/pavlovic265/265-gt/helpers"
	"github.com/pavlovic265/265-gt/utils/pointer"
)

const (
	githubAPIBase        = "https://api.github.com"
	githubGraphQLAPIBase = "https://api.github.com/graphql"
)

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

func getConfiguredMergeMethod(ctx context.Context) (constants.MergeMethod, error) {
	cfg, ok := config.GetConfig(ctx)
	if !ok {
		return "", ErrConfigNotLoaded
	}
	if cfg.Local == nil || cfg.Local.MergeMethod == "" {
		return "", fmt.Errorf("no merge method configured for this repository; run 'gt config local'")
	}

	return cfg.Local.MergeMethod, nil
}

func (c *gitHubClient) doGraphQLRequest(
	ctx context.Context, token, query string, variables map[string]any, out any,
) error {
	payload := map[string]any{
		"query":     query,
		"variables": variables,
	}

	jsonBody, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(
		ctx, http.MethodPost, githubGraphQLAPIBase, bytes.NewBuffer(jsonBody),
	)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	resp, err := githubHTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("graphql request failed: %s", resp.Status)
	}

	return json.NewDecoder(resp.Body).Decode(out)
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

	payload := map[string]any{
		"query": githubListOpenPullRequestsQuery,
		"variables": map[string]any{
			"owner": repoInfo.Owner,
			"repo":  repoInfo.Repo,
		},
	}

	jsonBody, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(
		ctx, http.MethodPost, githubGraphQLAPIBase, bytes.NewBuffer(jsonBody),
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+account.Token)
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	resp, err := githubHTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to list PRs: %s", resp.Status)
	}

	var result struct {
		Data struct {
			Repository struct {
				PullRequests struct {
					Nodes []struct {
						Number         int    `json:"number"`
						Title          string `json:"title"`
						URL            string `json:"url"`
						Mergeable      string `json:"mergeable"`
						ReviewDecision string `json:"reviewDecision"`
						IsInMergeQueue bool   `json:"isInMergeQueue"`
						Author         *struct {
							Login string `json:"login"`
						} `json:"author"`
						HeadRefName string `json:"headRefName"`
						Commits     struct {
							Nodes []struct {
								Commit struct {
									StatusCheckRollup *struct {
										State string `json:"state"`
									} `json:"statusCheckRollup"`
								} `json:"commit"`
							} `json:"nodes"`
						} `json:"commits"`
					} `json:"nodes"`
				} `json:"pullRequests"`
			} `json:"repository"`
		} `json:"data"`
		Errors []struct {
			Message string `json:"message"`
		} `json:"errors"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if len(result.Errors) > 0 {
		return nil, fmt.Errorf("failed to list PRs: %s", result.Errors[0].Message)
	}

	var prs []PullRequest
	for _, pr := range result.Data.Repository.PullRequests.Nodes {
		if pr.Author == nil || pr.Author.Login != account.User {
			continue
		}

		prs = append(prs, PullRequest{
			Number:      pr.Number,
			Title:       pr.Title,
			URL:         pr.URL,
			Author:      pr.Author.Login,
			Branch:      pr.HeadRefName,
			Mergeable:   pr.Mergeable,
			StatusState: mapGraphQLStatusState(pr.Commits),
			ReviewState: mapGraphQLReviewDecision(pr.ReviewDecision),
			MergeQueued: pr.IsInMergeQueue,
		})
	}

	return prs, nil
}

func mapGraphQLStatusState(commits struct {
	Nodes []struct {
		Commit struct {
			StatusCheckRollup *struct {
				State string `json:"state"`
			} `json:"statusCheckRollup"`
		} `json:"commit"`
	} `json:"nodes"`
},
) StatusStateType {
	if len(commits.Nodes) == 0 {
		return ""
	}

	rollup := commits.Nodes[0].Commit.StatusCheckRollup
	if rollup == nil {
		return ""
	}

	switch rollup.State {
	case "SUCCESS":
		return StatusStateTypeSuccess
	case "FAILURE", "ERROR":
		return StatusStateTypeFailure
	case "EXPECTED", "PENDING":
		return StatusStateTypePending
	default:
		return ""
	}
}

func mapGraphQLReviewDecision(reviewDecision string) ReviewStateType {
	switch reviewDecision {
	case "APPROVED":
		return ReviewStateApproved
	case "CHANGES_REQUESTED":
		return ReviewStateChangesRequested
	default:
		return ""
	}
}

func (c *gitHubClient) HasOpenPullRequestForBranch(
	ctx context.Context, branch string,
) (bool, error) {
	repoInfo, account, err := c.getRepoInfo(ctx)
	if err != nil {
		return false, err
	}

	query := url.Values{}
	query.Set("state", "open")
	query.Set("head", fmt.Sprintf("%s:%s", repoInfo.Owner, branch))

	apiURL := fmt.Sprintf("%s/repos/%s/%s/pulls?%s",
		githubAPIBase, repoInfo.Owner, repoInfo.Repo, query.Encode())

	resp, err := c.doRequest(ctx, "GET", apiURL, nil, account.Token)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return false, fmt.Errorf("failed to list PRs: %s", resp.Status)
	}

	var ghPRs []json.RawMessage

	if err := json.NewDecoder(resp.Body).Decode(&ghPRs); err != nil {
		return false, err
	}

	return len(ghPRs) > 0, nil
}

func (c *gitHubClient) getPullRequestNumberForBranch(
	ctx context.Context, branch string,
) (int, error) {
	repoInfo, account, err := c.getRepoInfo(ctx)
	if err != nil {
		return 0, err
	}

	query := url.Values{}
	query.Set("state", "open")
	query.Set("head", fmt.Sprintf("%s:%s", repoInfo.Owner, branch))

	apiURL := fmt.Sprintf("%s/repos/%s/%s/pulls?%s",
		githubAPIBase, repoInfo.Owner, repoInfo.Repo, query.Encode())

	resp, err := c.doRequest(ctx, "GET", apiURL, nil, account.Token)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return 0, fmt.Errorf("failed to list PRs: %s", resp.Status)
	}

	var ghPRs []struct {
		Number int `json:"number"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&ghPRs); err != nil {
		return 0, err
	}

	if len(ghPRs) == 0 {
		return 0, nil
	}

	return ghPRs[0].Number, nil
}

func (c *gitHubClient) MergePullRequest(ctx context.Context, prNumber int) error {
	repoInfo, account, err := c.getRepoInfo(ctx)
	if err != nil {
		return err
	}

	mergeMethod, err := getConfiguredMergeMethod(ctx)
	if err != nil {
		return err
	}

	if mergeMethod == constants.MergeMethodQueue {
		return c.enqueuePullRequest(ctx, repoInfo, account.Token, prNumber)
	}

	url := fmt.Sprintf("%s/repos/%s/%s/pulls/%d/merge", githubAPIBase, repoInfo.Owner, repoInfo.Repo, prNumber)
	resp, err := c.doRequest(ctx, "PUT", url, map[string]string{
		"merge_method": mergeMethod.String(),
	}, account.Token)
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

func (c *gitHubClient) enqueuePullRequest(
	ctx context.Context, repoInfo *RepoInfo, token string, prNumber int,
) error {
	var queryResult struct {
		Data struct {
			Repository struct {
				PullRequest *struct {
					ID string `json:"id"`
				} `json:"pullRequest"`
			} `json:"repository"`
		} `json:"data"`
		Errors []struct {
			Message string `json:"message"`
		} `json:"errors"`
	}

	err := c.doGraphQLRequest(ctx, token, githubPullRequestNodeIDQuery, map[string]any{
		"owner":  repoInfo.Owner,
		"repo":   repoInfo.Repo,
		"number": prNumber,
	}, &queryResult)
	if err != nil {
		return err
	}
	if len(queryResult.Errors) > 0 {
		return fmt.Errorf("failed to enqueue PR: %s", queryResult.Errors[0].Message)
	}
	if queryResult.Data.Repository.PullRequest == nil || queryResult.Data.Repository.PullRequest.ID == "" {
		return fmt.Errorf("failed to enqueue PR: pull request not found")
	}

	var mutationResult struct {
		Data struct {
			EnqueuePullRequest struct {
				MergeQueueEntry *struct {
					ID string `json:"id"`
				} `json:"mergeQueueEntry"`
			} `json:"enqueuePullRequest"`
		} `json:"data"`
		Errors []struct {
			Message string `json:"message"`
		} `json:"errors"`
	}

	err = c.doGraphQLRequest(ctx, token, githubEnqueuePullRequestMutation, map[string]any{
		"pullRequestId": queryResult.Data.Repository.PullRequest.ID,
	}, &mutationResult)
	if err != nil {
		return err
	}
	if len(mutationResult.Errors) > 0 {
		return fmt.Errorf("failed to enqueue PR: %s", mutationResult.Errors[0].Message)
	}
	if mutationResult.Data.EnqueuePullRequest.MergeQueueEntry == nil {
		return fmt.Errorf("failed to enqueue PR: merge queue entry not created")
	}

	return nil
}

func (c *gitHubClient) UpdatePullRequestBaseBranch(ctx context.Context, branch string) error {
	repoInfo, account, err := c.getRepoInfo(ctx)
	if err != nil {
		return err
	}

	parent, err := c.gitHelper.GetParent(branch)
	if err != nil {
		return err
	}

	prNumber, err := c.getPullRequestNumberForBranch(ctx, branch)
	if err != nil {
		return err
	}
	if prNumber == 0 {
		return nil
	}

	payload := map[string]string{
		"base": parent,
	}

	apiURL := fmt.Sprintf("%s/repos/%s/%s/pulls/%d", githubAPIBase, repoInfo.Owner, repoInfo.Repo, prNumber)
	resp, err := c.doRequest(ctx, "PATCH", apiURL, payload, account.Token)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("failed to update PR base branch: status %d", resp.StatusCode)
	}

	return nil
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
	MergeQueued bool            `json:"mergeQueued"`
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

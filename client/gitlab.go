package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/pavlovic265/265-gt/config"
	helpers "github.com/pavlovic265/265-gt/git_helpers"
	"github.com/pavlovic265/265-gt/utils/pointer"
)

const gitlabAPIBase = "https://gitlab.com/api/v4"

var gitlabHTTPClient = &http.Client{Timeout: 30 * time.Second}

type gitLabClient struct {
	gitHelper helpers.GitHelper
}

func NewGitLabClient(gitHelper helpers.GitHelper) CliClient {
	return &gitLabClient{gitHelper: gitHelper}
}

func (c *gitLabClient) getProjectInfo(ctx context.Context) (string, *config.Account, error) {
	cfg, ok := config.GetConfig(ctx)
	if !ok {
		return "", nil, fmt.Errorf("config not loaded")
	}

	if cfg.Global.ActiveAccount == nil {
		return "", nil, fmt.Errorf("no active account")
	}

	remoteURL, err := c.gitHelper.GetRemoteURL("origin")
	if err != nil {
		return "", nil, fmt.Errorf("failed to get remote URL: %w", err)
	}

	repoInfo, err := ParseRemoteURL(remoteURL)
	if err != nil {
		return "", nil, err
	}

	// GitLab uses URL-encoded project path
	projectPath := url.PathEscape(repoInfo.Owner + "/" + repoInfo.Repo)

	return projectPath, cfg.Global.ActiveAccount, nil
}

func (c *gitLabClient) doRequest(
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

	req.Header.Set("PRIVATE-TOKEN", token)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	return gitlabHTTPClient.Do(req)
}

func (c *gitLabClient) AuthStatus(ctx context.Context) error {
	cfg, ok := config.GetConfig(ctx)
	if !ok {
		return fmt.Errorf("config not loaded")
	}

	account := cfg.Global.ActiveAccount
	if account == nil {
		return fmt.Errorf("no active account")
	}

	resp, err := c.doRequest(ctx, "GET", gitlabAPIBase+"/user", nil, account.Token)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("authentication failed: %s", resp.Status)
	}

	var user struct {
		Username string `json:"username"`
		Name     string `json:"name"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return err
	}

	fmt.Printf("Logged in as %s (%s)\n", user.Username, account.Platform)
	if user.Name != "" {
		fmt.Printf("Name: %s\n", user.Name)
	}

	return nil
}

func (c *gitLabClient) AuthLogin(ctx context.Context, user string) error {
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

func (c *gitLabClient) AuthLogout(ctx context.Context, user string) error {
	cfg, ok := config.GetConfig(ctx)
	if !ok {
		return fmt.Errorf("config not loaded")
	}

	cfg.Global.ActiveAccount = nil
	cfg.MarkDirty()
	return nil
}

func (c *gitLabClient) CreatePullRequest(ctx context.Context, args []string) error {
	projectPath, account, err := c.getProjectInfo(ctx)
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

	payload := map[string]any{
		"source_branch": branch,
		"target_branch": parent,
		"title":         branch,
	}

	apiURL := fmt.Sprintf("%s/projects/%s/merge_requests", gitlabAPIBase, projectPath)
	resp, err := c.doRequest(ctx, "POST", apiURL, payload, account.Token)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 201 {
		var errResp struct {
			Message any `json:"message"`
		}
		_ = json.NewDecoder(resp.Body).Decode(&errResp)
		return fmt.Errorf("failed to create MR: %v", errResp.Message)
	}

	var mr struct {
		WebURL string `json:"web_url"`
		IID    int    `json:"iid"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&mr); err != nil {
		return err
	}

	fmt.Printf("Created MR !%d: %s\n", mr.IID, mr.WebURL)
	return nil
}

func (c *gitLabClient) ListPullRequests(ctx context.Context, args []string) ([]PullRequest, error) {
	projectPath, account, err := c.getProjectInfo(ctx)
	if err != nil {
		return nil, err
	}

	apiURL := fmt.Sprintf("%s/projects/%s/merge_requests?state=opened&author_username=%s",
		gitlabAPIBase, projectPath, account.User)

	resp, err := c.doRequest(ctx, "GET", apiURL, nil, account.Token)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("failed to list MRs: %s", resp.Status)
	}

	var glMRs []struct {
		IID    int    `json:"iid"`
		Title  string `json:"title"`
		WebURL string `json:"web_url"`
		Author struct {
			Username string `json:"username"`
		} `json:"author"`
		SourceBranch string `json:"source_branch"`
		MergeStatus  string `json:"merge_status"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&glMRs); err != nil {
		return nil, err
	}

	var prs []PullRequest
	for _, mr := range glMRs {
		mergeable := "UNKNOWN"
		switch mr.MergeStatus {
		case "can_be_merged":
			mergeable = "MERGEABLE"
		case "cannot_be_merged":
			mergeable = "CONFLICTING"
		}

		prs = append(prs, PullRequest{
			Number:    mr.IID,
			Title:     mr.Title,
			URL:       mr.WebURL,
			Author:    mr.Author.Username,
			Branch:    mr.SourceBranch,
			Mergeable: mergeable,
		})
	}

	return prs, nil
}

func (c *gitLabClient) MergePullRequest(ctx context.Context, prNumber int) error {
	projectPath, account, err := c.getProjectInfo(ctx)
	if err != nil {
		return err
	}

	apiURL := fmt.Sprintf("%s/projects/%s/merge_requests/%d/merge", gitlabAPIBase, projectPath, prNumber)
	resp, err := c.doRequest(ctx, "PUT", apiURL, nil, account.Token)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		var errResp struct {
			Message string `json:"message"`
		}
		_ = json.NewDecoder(resp.Body).Decode(&errResp)
		return fmt.Errorf("failed to merge MR: %s", errResp.Message)
	}

	return nil
}

func (c *gitLabClient) UpdatePullRequestBranch(ctx context.Context, prNumber int) error {
	projectPath, account, err := c.getProjectInfo(ctx)
	if err != nil {
		return err
	}

	apiURL := fmt.Sprintf("%s/projects/%s/merge_requests/%d/rebase", gitlabAPIBase, projectPath, prNumber)
	resp, err := c.doRequest(ctx, "PUT", apiURL, nil, account.Token)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 202 {
		return fmt.Errorf("failed to rebase MR: status %d", resp.StatusCode)
	}

	return nil
}

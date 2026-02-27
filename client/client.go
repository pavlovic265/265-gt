package client

import (
	"context"
	"fmt"

	"github.com/pavlovic265/265-gt/constants"
	helpers "github.com/pavlovic265/265-gt/helpers"
)

type CliClient interface {
	AuthStatus(ctx context.Context) error
	AuthLogin(ctx context.Context, user string) error
	AuthLogout(ctx context.Context, user string) error
	CreatePullRequest(ctx context.Context, args []string) error
	ListPullRequests(ctx context.Context, args []string) ([]PullRequest, error)
	MergePullRequest(ctx context.Context, prNumber int) error
	UpdatePullRequestBranch(ctx context.Context, prNumber int) error
}

func NewRestCliClient(platform constants.Platform, gitHelper helpers.GitHelper) (CliClient, error) {
	switch platform {
	case constants.GitHubPlatform:
		return NewGitHubClient(gitHelper), nil
	case constants.GitLabPlatform:
		return NewGitLabClient(gitHelper), nil
	default:
		return nil, fmt.Errorf("unsupported platform: %q", platform)
	}
}

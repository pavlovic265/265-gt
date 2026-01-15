package client

import (
	"context"

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

var Client map[constants.Platform]CliClient

func InitCliClient(gitHelper helpers.GitHelper) {
	Client = map[constants.Platform]CliClient{
		constants.GitHubPlatform: NewGitHubClient(gitHelper),
		constants.GitLabPlatform: NewGitLabClient(gitHelper),
	}
}

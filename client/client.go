package client

import (
	"context"

	"github.com/pavlovic265/265-gt/constants"
	"github.com/pavlovic265/265-gt/executor"
	helpers "github.com/pavlovic265/265-gt/git_helpers"
)

type CliClient interface {
	AuthStatus(ctx context.Context) error
	AuthLogin(ctx context.Context, user string) error
	AuthLogout(ctx context.Context, user string) error
	CreatePullRequest(ctx context.Context, args []string) error
	ListPullRequests(ctx context.Context, args []string) ([]PullRequest, error)
	MergePullRequest(prNumber int) error
	UpdatePullRequestBranch(prNumber int) error
}

var Client map[constants.Platform]CliClient

func InitCliClient(exe executor.Executor, gitHelper helpers.GitHelper) error {
	Client = map[constants.Platform]CliClient{
		constants.GitHubPlatform: NewGitHubCli(exe, gitHelper),
		constants.GitLabPlatform: NewGitLabCli(exe, gitHelper),
	}
	return nil
}

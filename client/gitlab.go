package client

import (
	"context"

	"github.com/pavlovic265/265-gt/executor"
	helpers "github.com/pavlovic265/265-gt/git_helpers"
)

type gitLabCli struct {
	exe       executor.Executor
	gitHelper helpers.GitHelper
}

func NewGitLabCli(exe executor.Executor, gitHelper helpers.GitHelper) CliClient {
	return &gitLabCli{exe: exe, gitHelper: gitHelper}
}

func (svc gitLabCli) AuthStatus(ctx context.Context) error {
	return nil
}

func (svc gitLabCli) AuthLogin(ctx context.Context, user string) error {
	return nil
}

func (svc gitLabCli) AuthLogout(ctx context.Context, user string) error {
	return nil
}

func (svc gitLabCli) CreatePullRequest(ctx context.Context, args []string) error {
	return nil
}

func (svc gitLabCli) ListPullRequests(ctx context.Context, args []string) ([]PullRequest, error) {
	return nil, nil
}

func (svc gitLabCli) MergePullRequest(prNumber int) error {
	return nil
}

func (svc gitLabCli) UpdatePullRequestBranch(prNumber int) error {
	return nil
}

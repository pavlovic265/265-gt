package client

import "github.com/pavlovic265/265-gt/executor"

type gitLabCli struct {
	exe executor.Executor
}

func NewGitLabCli(exe executor.Executor) CliClient {
	return &gitLabCli{exe: exe}
}

func (svc gitLabCli) AuthStatus() error {
	return nil
}

func (svc gitLabCli) CreatePullRequest(args []string) error {
	return nil
}

func (svc gitLabCli) ListPullRequests(args []string) ([]PullRequest, error) {
	return nil, nil
}

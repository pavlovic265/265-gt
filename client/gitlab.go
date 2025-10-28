package client

import (
	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/executor"
	helpers "github.com/pavlovic265/265-gt/git_helpers"
)

type gitLabCli struct {
	exe           executor.Executor
	configManager config.ConfigManager
	gitHelper     helpers.GitHelper
}

func NewGitLabCli(exe executor.Executor, configManager config.ConfigManager, gitHelper helpers.GitHelper) CliClient {
	return &gitLabCli{exe: exe, configManager: configManager, gitHelper: gitHelper}
}

func (svc gitLabCli) AuthStatus() error {
	return nil
}

func (svc gitLabCli) AuthLogin(user string) error {
	return nil
}

func (svc gitLabCli) AuthLogout(user string) error {
	return nil
}

func (svc gitLabCli) CreatePullRequest(args []string) error {
	return nil
}

func (svc gitLabCli) ListPullRequests(args []string) ([]PullRequest, error) {
	return nil, nil
}

func (svc gitLabCli) MergePullRequest(prNumber int) error {
	return nil
}

func (svc gitLabCli) UpdatePullRequestBranch(prNumber int) error {
	return nil
}

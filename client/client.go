package client

import (
	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/executor"
)

type CliClient interface {
	AuthStatus() error
	AuthLogin(user string) error
	AuthLogout(user string) error
	CreatePullRequest(args []string) error
	ListPullRequests(args []string) ([]PullRequest, error)
}

var Client map[config.Platform]CliClient

func InitCliClient(exe executor.Executor) {
	Client = map[config.Platform]CliClient{
		config.GitHubPlatform: NewGitHubCli(exe),
		config.GitLabPlatform: NewGitLabCli(exe),
	}
}

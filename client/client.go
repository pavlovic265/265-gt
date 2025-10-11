package client

import (
	"github.com/pavlovic265/265-gt/constants"
	"github.com/pavlovic265/265-gt/executor"
)

type CliClient interface {
	AuthStatus() error
	AuthLogin(user string) error
	AuthLogout(user string) error
	CreatePullRequest(args []string) error
	ListPullRequests(args []string) ([]PullRequest, error)
}

var Client map[constants.Platform]CliClient

func InitCliClient(exe executor.Executor) {
	Client = map[constants.Platform]CliClient{
		constants.GitHubPlatform: NewGitHubCli(exe),
		constants.GitLabPlatform: NewGitLabCli(exe),
	}
}

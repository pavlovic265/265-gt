package client

import (
	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/constants"
	"github.com/pavlovic265/265-gt/executor"
	helpers "github.com/pavlovic265/265-gt/git_helpers"
)

type CliClient interface {
	AuthStatus() error
	AuthLogin(user string) error
	AuthLogout(user string) error
	CreatePullRequest(args []string) error
	ListPullRequests(args []string) ([]PullRequest, error)
	MergePullRequest(prNumber int) error
	UpdatePullRequestBranch(prNumber int) error
}

var Client map[constants.Platform]CliClient

func InitCliClient(exe executor.Executor, configManager config.ConfigManager, gitHelper helpers.GitHelper) error {
	Client = map[constants.Platform]CliClient{
		constants.GitHubPlatform: NewGitHubCli(exe, configManager, gitHelper),
		constants.GitLabPlatform: NewGitLabCli(exe, configManager, gitHelper),
	}
	return nil
}

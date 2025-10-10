package client

import (
	"bytes"
	"os/exec"
	"strings"

	"github.com/pavlovic265/265-gt/executor"
)

type CliClient interface {
	AuthStatus() error
	AuthLogin(user string) error
	AuthLogout(user string) error
	CreatePullRequest(args []string) error
	ListPullRequests(args []string) ([]PullRequest, error)
}

var GlobalClient CliClient

func InitCliClient(exe executor.Executor) {
	InitCliClientWithGit(exe, true)
}

func InitCliClientWithGit(exe executor.Executor, requireGit bool) {
	if !requireGit {
		// For commands that don't need git (like version), skip client initialization
		return
	}

	remoteOrigin, err := getGitRemoteOrigin()
	if err != nil {
		// For auth commands, we still want to initialize a client even without git
		// Default to GitHub client for auth operations
		GlobalClient = NewGitHubCli(exe)
		return
	}

	if strings.Contains(remoteOrigin, "github.com") {
		GlobalClient = NewGitHubCli(exe)
	} else if strings.Contains(remoteOrigin, "gitlab.com") {
		GlobalClient = NewGitHubCli(exe)
	}
}

// getGitRemoteOrigin retrieves the remote URL for 'origin'
func getGitRemoteOrigin() (string, error) {
	cmd := exec.Command("git", "remote", "get-url", "origin")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(out.String()), nil
}

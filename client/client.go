package client

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

type CliClient interface {
	CreatePullRequest(args []string) error
}

var GlobalClient CliClient

func InitCliClient() {
	remoteOrigin, err := getGitRemoteOrigin()
	if err != nil {
		fmt.Println("Error retrieving remote URL:", err)
	}

	if strings.Contains(remoteOrigin, "github.com") {
		GlobalClient = NewGitHubCli()
	} else if strings.Contains(remoteOrigin, "gitlab.com") {
		GlobalClient = NewGitHubCli()
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

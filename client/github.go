package client

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/pavlovic265/265-gt/config"
)

type gitHubCli struct{}

func NewGitHubCli() CliClient {
	return &gitHubCli{}
}

func (svc *gitHubCli) CreatePullRequest(args []string) error {
	// gh pr create --assignee MarkoPavlovic265 --fill
	fmt.Println("Creating pull request on GitHub...")
	exeArgs := append([]string{"pr", "create", "--assignee", config.GlobalConfig.GitHub.Assignee, "--fill"}, args...)
	exeCmd := exec.Command("gh", exeArgs...)
	exeCmd.Stdout = os.Stdout
	exeCmd.Stderr = os.Stderr

	if err := exeCmd.Run(); err != nil {
		return fmt.Errorf(
			"error executing gh create --assignee %s --fill %v",
			config.GlobalConfig.GitHub.Assignee,
			err,
		)
	}

	return nil
}

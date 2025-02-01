package client

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/pavlovic265/265-gt/config"
)

type gitHubCli struct{}

func NewGitHubCli() CliClient {
	return &gitHubCli{}
}

func (svc gitHubCli) AuthStatus() error {
	exeArgs := append([]string{"auth", "status"})
	exeCmd := exec.Command("gh", exeArgs...)
	outputByte, err := exeCmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf(
			"faild to check github aith status, with err (%w)",
			err,
		)
	}

	outputStr := string(outputByte)
	sections := strings.Split(strings.Join(strings.Split(outputStr, "\n")[1:], "\n"), "\n\n")

	for _, section := range sections {
		if strings.Contains(section, "- Active account: true") {
			fmt.Fprintf(os.Stderr, section+"\n")
			return nil
		}
	}

	fmt.Fprintf(os.Stderr, "There are no active accounts \n")

	return nil
}

func (svc *gitHubCli) CreatePullRequest(args []string) error {
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

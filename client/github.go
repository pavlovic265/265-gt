package client

import (
	"fmt"
	"os"
	"strings"

	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/executor"
)

type gitHubCli struct {
	exe executor.Executor
}

func NewGitHubCli(exe executor.Executor) CliClient {
	return &gitHubCli{exe: exe}
}

func (svc gitHubCli) AuthStatus() error {
	exeArgs := []string{"auth", "status"}
	output, err := svc.exe.ExecuteWithOutput("gh", exeArgs...)
	if err != nil {
		return err
	}

	outputStr := string(output)
	fmt.Println("er routputStr ", outputStr)
	sections := strings.Split(strings.Join(strings.Split(outputStr, "\n")[1:], "\n"), "\n\n")

	for _, section := range sections {
		if strings.Contains(section, "- Active account: true") {
			fmt.Fprintf(os.Stderr, "%s\n", section)
			return nil
		}
	}

	fmt.Fprintf(os.Stderr, "There are no active accounts\n")

	return nil
}

func (svc *gitHubCli) CreatePullRequest(args []string) error {
	fmt.Println("Creating pull request on GitHub...")

	exeArgs := []string{"pr", "create", "--assignee", config.GlobalConfig.GitHub.Assignee, "--fill"}
	if err := svc.exe.Execute("gh", exeArgs...); err != nil {
		return err
	}

	return nil
}

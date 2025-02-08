package client

import (
	"bytes"
	"encoding/json"
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

func (svc gitHubCli) getActiveAccount() (*config.Account, error) {
	exeArgs := []string{"auth", "status"}
	output, err := svc.exe.ExecuteWithOutput("gh", exeArgs...)
	if err != nil {
		return nil, err
	}

	outputStr := string(output)
	sections := strings.Split(strings.Join(strings.Split(outputStr, "\n")[1:], "\n"), "\n\n")

	for _, section := range sections {
		if strings.Contains(section, "- Active account: true") {
			rows := strings.Split(section, "\n")
			var user, tokenPrefix string
			for _, row := range rows {
				if strings.Contains(row, "keyring") {
					account := strings.Split(row, " ")
					user = account[len(account)-1]
				}
				if strings.Contains(row, "Token:") {
					tokenPrefix = strings.Split(strings.Split(row, " ")[1], "*")[0]
				}
			}

			accoutns := config.GlobalConfig.GitHub.Accounts
			for _, acc := range accoutns {
				if acc.User == user || strings.HasPrefix(acc.Token, tokenPrefix) {
					return &acc, nil
				}
			}
		}
	}

	return nil, fmt.Errorf("could not found account")
}

func (svc gitHubCli) AuthStatus() error {
	exeArgs := []string{"auth", "status"}
	output, err := svc.exe.ExecuteWithOutput("gh", exeArgs...)
	if err != nil {
		return err
	}

	outputStr := string(output)
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

	acc, err := svc.getActiveAccount()
	if err != nil {
		return err
	}

	exeArgs := []string{"pr", "create", "--assignee", acc.User, "--fill"}
	if err := svc.exe.Execute("gh", exeArgs...); err != nil {
		return err
	}

	return nil
}

type PullRequest struct {
	Number int    `json:"number"`
	Title  string `json:"title"`
	URL    string `json:"url"`
	Author string `json:"author"`
}

func (svc *gitHubCli) ListPullRequests(args []string) ([]PullRequest, error) {
	acc, err := svc.getActiveAccount()
	if err != nil {
		return nil, err
	}
	exeArgs := []string{"pr", "list", "--author", acc.User, "--json", "number,title,url,author"}

	var out bytes.Buffer

	if err := svc.exe.ExecuteWithStdout("gh", &out, exeArgs...); err != nil {
		return nil, err
	}
	var rawPRs []struct {
		Number int    `json:"number"`
		Title  string `json:"title"`
		URL    string `json:"url"`
		Author struct {
			Login string `json:"login"`
		} `json:"author"`
	}
	err = json.Unmarshal(out.Bytes(), &rawPRs)
	if err != nil {
		return nil, err
	}

	var prs []PullRequest
	for _, pr := range rawPRs {
		prs = append(prs, PullRequest{
			Number: pr.Number,
			Title:  pr.Title,
			URL:    pr.URL,
			Author: pr.Author.Login,
		})
	}
	return prs, nil
}

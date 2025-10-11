package client

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/constants"
	"github.com/pavlovic265/265-gt/executor"
	"github.com/pavlovic265/265-gt/helpers"
	pointer "github.com/pavlovic265/265-gt/utils/pointer"
)

type gitHubCli struct {
	exe       executor.Executor
	gitHelper helpers.GitHelper
}

func NewGitHubCli(exe executor.Executor) CliClient {
	return &gitHubCli{
		exe:       exe,
		gitHelper: helpers.NewGitHelper(),
	}
}

func (svc gitHubCli) getActiveAccount() (*config.Account, error) {
	exeArgs := []string{"auth", "status"}
	output, err := svc.exe.WithGh().WithArgs(exeArgs).RunWithOutput()
	if err != nil {
		return nil, err
	}

	outputStr := output.String()
	sections := strings.Split(strings.Join(strings.Split(outputStr, "\n")[1:], "\n"), "\n\n")

	for _, section := range sections {
		if strings.Contains(section, "- Active account: true") {
			rows := strings.Split(section, "\n")
			var user, tokenPrefix string
			for _, row := range rows {
				if strings.Contains(row, "keyring") {
					account := strings.Split(row, " ")
					user = account[len(account)-2]
				}
				if strings.Contains(row, "Token:") {
					tokenPrefix = strings.Split(strings.Split(row, " ")[1], "*")[0]
				}
			}

			accounts := config.GlobalConfig.Accounts
			for _, acc := range accounts {
				if acc.User == user && strings.HasPrefix(acc.Token, tokenPrefix) {
					return &acc, nil
				}
			}
		}
	}

	return nil, fmt.Errorf("could not found account")
}

func (svc gitHubCli) AuthStatus() error {
	exeArgs := []string{"auth", "status"}
	output, err := svc.exe.WithGh().WithArgs(exeArgs).RunWithOutput()
	if err != nil {
		return err
	}

	// Format and display the output with beautiful UI
	svc.displayAuthStatus(output.String())

	return nil
}

func (svc gitHubCli) AuthLogin(user string) error {
	accounts := config.GlobalConfig.Accounts
	var account config.Account
	for _, acc := range accounts {
		if acc.User == user {
			account = acc
			break
		}
	}

	exeArgs := []string{"auth", "login", "--with-token"}
	err := svc.exe.WithGh().WithArgs(exeArgs).WithStdin(account.Token).Run()
	if err != nil {
		return err
	}

	err = config.SetActiveAccount(account)
	if err != nil {
		return err
	}

	return nil
}

func (svc gitHubCli) AuthLogout(user string) error {
	if !config.HasActiveAccount() {
		return fmt.Errorf("no active account found")
	}

	exeArgs := []string{"auth", "logout", "-u", user}
	err := svc.exe.WithGh().WithArgs(exeArgs).Run()
	if err != nil {
		return err
	}

	acc, err := svc.getActiveAccount()
	if err != nil {
		return err
	}

	if acc != nil {
		err = config.SetActiveAccount(*acc)
		if err != nil {
			return err
		}
	} else {
		err = config.ClearActiveAccount()
		if err != nil {
			return err
		}
	}

	return nil
}

func (svc *gitHubCli) CreatePullRequest(args []string) error {
	fmt.Println("Creating pull request on GitHub...")
	acc, err := svc.getActiveAccount()
	if err != nil {
		return err
	}

	branch, err := svc.gitHelper.GetCurrentBranchName(svc.exe)
	if err != nil {
		return err
	}
	parent := svc.gitHelper.GetParent(svc.exe, pointer.Deref(branch))

	exeArgs := []string{
		"pr",
		"create",
		"--fill",
		"--assignee", acc.User,
		"--base", parent,
	}

	// Add any additional args (like --draft)
	exeArgs = append(exeArgs, args...)

	err = svc.exe.WithGh().WithArgs(exeArgs).Run()
	if err != nil {
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
	out, err := svc.exe.WithGh().WithArgs(exeArgs).RunWithOutput()
	if err != nil {
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

func (svc gitHubCli) displayAuthStatus(output string) {
	// Simple title with subtle color
	fmt.Println(constants.GetInfoStyle().Render("GitHub Authentication Status"))
	fmt.Println()

	lines := strings.Split(output, "\n")
	var currentPlatform string

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if line == "" {
			continue
		}

		// Platform header (e.g., "github.com")
		if !strings.HasPrefix(line, " ") && !strings.HasPrefix(line, "✓") && !strings.HasPrefix(line, "-") {
			if currentPlatform != "" {
				fmt.Println() // Add spacing between platforms
			}
			currentPlatform = line
			fmt.Println(constants.GetCommandStyle().Render("> " + currentPlatform))
			continue
		}

		// Style the lines with subtle colors
		if strings.HasPrefix(line, "✓") {
			// Success lines (logged in accounts)
			fmt.Println(constants.GetSuccessStyle().Render(line))
		} else if strings.HasPrefix(line, "-") {
			// Detail lines
			if strings.Contains(line, "Active account: true") {
				fmt.Println(constants.GetSuccessStyle().Render("  " + line))
			} else if strings.Contains(line, "Active account: false") {
				fmt.Println(constants.GetDebugStyle().Render("  " + line))
			} else if strings.Contains(line, "Token:") {
				fmt.Println(constants.GetWarningStyle().Render("  " + line))
			} else {
				fmt.Println(constants.GetDebugStyle().Render("  " + line))
			}
		}
	}

	fmt.Println()

	// Show active account from our config
	activeAccount := config.GetActiveAccount()
	if activeAccount != nil {
		fmt.Println(constants.GetSuccessStyle().Render(
			"* Active Account: " + activeAccount.User + " (" + activeAccount.Platform.String() + ")"))
	} else {
		fmt.Println(constants.GetWarningStyle().Render("! No active account set in gt config"))
	}
}

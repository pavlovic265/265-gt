package client

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/constants"
	"github.com/pavlovic265/265-gt/executor"
	helpers "github.com/pavlovic265/265-gt/git_helpers"
	"github.com/pavlovic265/265-gt/utils/pointer"
)

type gitHubCli struct {
	exe       executor.Executor
	gitHelper helpers.GitHelper
}

func NewGitHubCli(
	exe executor.Executor,
	gitHelper helpers.GitHelper,
) CliClient {
	return &gitHubCli{
		exe:       exe,
		gitHelper: gitHelper,
	}
}

func (svc gitHubCli) getActiveAccount(ctx context.Context) (*config.Account, error) {
	cfg, ok := config.GetConfig(ctx)
	if !ok {
		return nil, fmt.Errorf("config not loaded")
	}

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

			for _, acc := range cfg.Global.Accounts {
				if acc.User == user && strings.HasPrefix(acc.Token, tokenPrefix) {
					return &acc, nil
				}
			}
		}
	}

	return nil, fmt.Errorf("could not found account")
}

func (svc gitHubCli) AuthStatus(ctx context.Context) error {
	exeArgs := []string{"auth", "status"}
	output, err := svc.exe.WithGh().WithArgs(exeArgs).RunWithOutput()
	if err != nil {
		return err
	}

	svc.displayAuthStatus(ctx, output.String())

	return nil
}

func (svc gitHubCli) AuthLogin(ctx context.Context, user string) error {
	cfg, ok := config.GetConfig(ctx)
	if !ok {
		return fmt.Errorf("config not loaded")
	}

	var account config.Account
	for _, acc := range cfg.Global.Accounts {
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

	cfg.Global.ActiveAccount = pointer.From(account)
	cfg.MarkDirty()

	return nil
}

func (svc gitHubCli) AuthLogout(ctx context.Context, user string) error {
	cfg, ok := config.GetConfig(ctx)
	if !ok {
		return fmt.Errorf("config not loaded")
	}

	if cfg.Global.ActiveAccount == nil || cfg.Global.ActiveAccount.User == "" {
		return fmt.Errorf("no active account found")
	}

	exeArgs := []string{"auth", "logout", "-u", user}
	err := svc.exe.WithGh().WithArgs(exeArgs).Run()
	if err != nil {
		return err
	}

	acc, err := svc.getActiveAccount(ctx)
	if err != nil {
		return err
	}

	if acc != nil {
		cfg.Global.ActiveAccount = acc
	} else {
		cfg.Global.ActiveAccount = nil
	}
	cfg.MarkDirty()

	return nil
}

func (svc *gitHubCli) CreatePullRequest(ctx context.Context, args []string) error {
	acc, err := svc.getActiveAccount(ctx)
	if err != nil {
		return err
	}

	branch, err := svc.gitHelper.GetCurrentBranch()
	if err != nil {
		return err
	}
	parent, err := svc.gitHelper.GetParent(branch)
	if err != nil {
		return err
	}

	exeArgs := []string{
		"pr",
		"create",
		"--fill",
		"--assignee", acc.User,
		"--base", parent,
	}

	exeArgs = append(exeArgs, args...)

	err = svc.exe.WithGh().WithArgs(exeArgs).Run()
	if err != nil {
		return err
	}

	return nil
}

type PullRequest struct {
	Number      int             `json:"number"`
	Title       string          `json:"title"`
	URL         string          `json:"url"`
	Author      string          `json:"author"`
	Mergeable   string          `json:"mergeable"`
	Branch      string          `json:"headRefName"`
	StatusState StatusStateType `json:"statusState"`
}

type StatusStateType string

var (
	StatusStateTypeSucess  StatusStateType
	StatusStateTypeFailur  StatusStateType
	StatusStateTypePending StatusStateType
)

func (svc *gitHubCli) ListPullRequests(ctx context.Context, args []string) ([]PullRequest, error) {
	acc, err := svc.getActiveAccount(ctx)
	if err != nil {
		return nil, err
	}

	exeArgs := []string{
		"pr", "list", "--author", acc.User, "--json",
		"number,title,url,author,mergeable,headRefName,statusCheckRollup",
	}
	out, err := svc.exe.WithGh().WithArgs(exeArgs).RunWithOutput()
	if err != nil {
		return nil, err
	}

	var rawPRs []struct {
		Number      int    `json:"number"`
		Title       string `json:"title"`
		URL         string `json:"url"`
		Mergeable   string `json:"mergeable"`
		HeadRefName string `json:"headRefName"`
		Author      struct {
			Login string `json:"login"`
		} `json:"author"`
		StatusCheckRollup []struct {
			Typename   string  `json:"__typename"`
			Conclusion *string `json:"conclusion,omitempty"`
			State      *string `json:"state,omitempty"`
		} `json:"statusCheckRollup"`
	}
	err = json.Unmarshal(out.Bytes(), &rawPRs)
	if err != nil {
		return nil, err
	}

	var prs []PullRequest
	for _, pr := range rawPRs {
		var statusState StatusStateType
		if len(pr.StatusCheckRollup) > 0 {
			hasFailure := false
			hasPending := false
			hasSuccess := false

		loop:
			for _, check := range pr.StatusCheckRollup {
				var s string
				switch check.Typename {
				case "StatusContext":
					s = pointer.Deref(check.State)
				case "CheckRun":
					s = pointer.Deref(check.Conclusion)
				}

				switch s {
				case "SKIPPED":
					continue
				case "SUCCESS":
					hasSuccess = true
				case "FAILURE", "TIMED_OUT", "CANCELLED":
					hasFailure = true
					break loop
				default:
					hasPending = true
					break loop
				}
			}

			if hasFailure {
				statusState = StatusStateTypeFailur
			} else if hasPending {
				statusState = StatusStateTypePending
			} else if hasSuccess {
				statusState = StatusStateTypeSucess
			}
		}

		prs = append(prs, PullRequest{
			Number:      pr.Number,
			Title:       pr.Title,
			URL:         pr.URL,
			Author:      pr.Author.Login,
			Mergeable:   pr.Mergeable,
			Branch:      pr.HeadRefName,
			StatusState: statusState,
		})
	}
	return prs, nil
}

func (svc *gitHubCli) MergePullRequest(prNumber int) error {
	exeArgs := []string{"pr", "merge", fmt.Sprintf("%d", prNumber)}
	err := svc.exe.WithGh().WithArgs(exeArgs).Run()
	if err != nil {
		return err
	}
	return nil
}

func (svc *gitHubCli) UpdatePullRequestBranch(prNumber int) error {
	exeArgs := []string{"pr", "comment", fmt.Sprintf("%d", prNumber), "--body", "@dependabot rebase"}
	err := svc.exe.WithGh().WithArgs(exeArgs).Run()
	if err != nil {
		return err
	}
	return nil
}

func (svc gitHubCli) displayAuthStatus(ctx context.Context, output string) {
	fmt.Println("GitHub Authentication Status")
	fmt.Println()

	lines := strings.Split(output, "\n")
	var currentPlatform string

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if line == "" {
			continue
		}

		if !strings.HasPrefix(line, " ") && !strings.HasPrefix(line, "✓") && !strings.HasPrefix(line, "-") {
			if currentPlatform != "" {
				fmt.Println()
			}
			currentPlatform = line
			fmt.Println("> " + currentPlatform)
			continue
		}

		if strings.HasPrefix(line, "✓") {
			fmt.Println(line)
		} else if strings.HasPrefix(line, "-") {
			fmt.Println("  " + line)
		}
	}

	fmt.Println()

	cfg, ok := config.GetConfig(ctx)
	if ok && cfg.Global.ActiveAccount != nil && cfg.Global.ActiveAccount.User != "" {
		activeAccount := cfg.Global.ActiveAccount
		fmt.Println(constants.GetSuccessAnsiStyle().Render(
			"* Active Account: " + activeAccount.User + " (" + activeAccount.Platform.String() + ")"))
	} else {
		fmt.Println("! No active account set in gt config")
	}
}

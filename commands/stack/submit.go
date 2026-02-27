package stack

import (
	"fmt"

	"github.com/pavlovic265/265-gt/client"
	"github.com/pavlovic265/265-gt/components"
	"github.com/pavlovic265/265-gt/config"
	helpers "github.com/pavlovic265/265-gt/helpers"
	"github.com/pavlovic265/265-gt/runner"
	"github.com/pavlovic265/265-gt/utils/log"
	"github.com/spf13/cobra"
)

type submitCommand struct {
	runner    runner.Runner
	gitHelper helpers.GitHelper
	cliClient client.CliClient
}

func NewSubmitCommand(
	runner runner.Runner,
	gitHelper helpers.GitHelper,
	cliClient client.CliClient,
) submitCommand {
	return submitCommand{
		runner:    runner,
		gitHelper: gitHelper,
		cliClient: cliClient,
	}
}

func (svc submitCommand) Command() *cobra.Command {
	var draft bool
	var interactive bool

	cmd := &cobra.Command{
		Use:     "submit-stack",
		Aliases: []string{"ss"},
		Short:   "Push and create PRs for the entire stack",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := svc.gitHelper.EnsureGitRepository(); err != nil {
				return err
			}

			cfg, err := config.RequireGlobal(cmd.Context())
			if err != nil {
				return err
			}

			if cfg.Global.ActiveAccount == nil || cfg.Global.ActiveAccount.User == "" {
				return log.ErrorMsg("no active account found")
			}

			prs, err := svc.cliClient.ListPullRequests(cmd.Context(), []string{})
			if err != nil {
				return log.Error("failed to list pull requests", err)
			}

			existingPRs := make(map[string]bool)
			for _, pr := range prs {
				existingPRs[pr.Branch] = true
			}

			originalBranch, err := svc.gitHelper.GetCurrentBranch()
			if err != nil {
				return err
			}

			queue := []string{originalBranch}
			submitted := 0
			created := 0

			for len(queue) > 0 {
				branch := queue[0]
				queue = queue[1:]

				if err := svc.runner.Git("checkout", branch); err != nil {
					return log.Error(fmt.Sprintf("failed to checkout branch %s", branch), err)
				}

				if interactive {
					choice, err := components.SelectString(
						[]string{"Create PR", "Create Draft PR", "Skip"},
					)
					if err != nil {
						return err
					}

					if choice == "" || choice == "Skip" {
						log.Infof("Skipping %s and its descendants", branch)
						continue
					}

					if choice == "Create Draft PR" {
						draft = true
					} else {
						draft = false
					}
				}

				if err := svc.runner.Git("push", "--force", "origin", branch); err != nil {
					return log.Error(fmt.Sprintf("failed to push branch %s", branch), err)
				}
				submitted++

				if !existingPRs[branch] {
					var prArgs []string
					if draft {
						prArgs = append(prArgs, "--draft")
					}
					if err := svc.cliClient.CreatePullRequest(
						cmd.Context(), prArgs,
					); err != nil {
						return log.Error(
							fmt.Sprintf("failed to create pull request for %s", branch), err,
						)
					}
					created++
					log.Successf("Created PR for %s", branch)
				} else {
					log.Infof("PR already exists for %s", branch)
				}

				children := svc.gitHelper.GetChildren(branch)
				for _, child := range children {
					if child != branch {
						queue = append(queue, child)
					}
				}
			}

			if err := svc.runner.Git("checkout", originalBranch); err != nil {
				return log.Error("failed to checkout original branch", err)
			}

			log.Successf(
				"Submit stack completed: %d pushed, %d PRs created", submitted, created,
			)
			return nil
		},
	}

	cmd.Flags().BoolVarP(&draft, "draft", "d", false, "Create all new PRs as drafts")
	cmd.Flags().BoolVarP(
		&interactive, "interactive", "i", false,
		"Interactively choose per-branch action",
	)

	return cmd
}

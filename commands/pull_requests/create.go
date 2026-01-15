package pullrequests

import (
	"github.com/pavlovic265/265-gt/client"
	"github.com/pavlovic265/265-gt/config"
	helpers "github.com/pavlovic265/265-gt/helpers"
	"github.com/pavlovic265/265-gt/runner"
	"github.com/pavlovic265/265-gt/utils/log"
	"github.com/spf13/cobra"
)

type createCommand struct {
	runner        runner.Runner
	configManager config.ConfigManager
	gitHelper     helpers.GitHelper
}

func NewCreateCommand(
	runner runner.Runner,
	configManager config.ConfigManager,
	gitHelper helpers.GitHelper,
) createCommand {
	return createCommand{
		runner:        runner,
		configManager: configManager,
		gitHelper:     gitHelper,
	}
}

var draft bool

func (svc createCommand) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "create",
		Aliases: []string{"c"},
		Short:   "Create a pull request",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := svc.gitHelper.EnsureGitRepository(); err != nil {
				return err
			}

			cfg, err := config.RequireGlobal(cmd.Context())
			if err != nil {
				return err
			}

			if cfg.Global.ActiveAccount == nil || cfg.Global.ActiveAccount.User == "" {
				return log.ErrorMsg("No active account found")
			}
			account := cfg.Global.ActiveAccount

			if draft {
				args = append([]string{"--draft"}, args...)
			}

			err = client.Client[account.Platform].CreatePullRequest(cmd.Context(), args)
			if err != nil {
				return log.Error("Failed to create pull request", err)
			}
			log.Success("Pull request created successfully")
			return nil
		},
	}

	cmd.Flags().BoolVarP(&draft, "draft", "d", false, "Create a draft pull request")

	return cmd
}

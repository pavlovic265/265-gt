package pr

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
	cliClient     client.CliClient
}

func NewCreateCommand(
	runner runner.Runner,
	configManager config.ConfigManager,
	gitHelper helpers.GitHelper,
	cliClient client.CliClient,
) createCommand {
	return createCommand{
		runner:        runner,
		configManager: configManager,
		gitHelper:     gitHelper,
		cliClient:     cliClient,
	}
}

func (svc createCommand) Command() *cobra.Command {
	var draft bool

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
				return log.ErrorMsg("no active account found")
			}

			if draft {
				args = append([]string{"--draft"}, args...)
			}

			err = svc.cliClient.CreatePullRequest(cmd.Context(), args)
			if err != nil {
				return log.Error("failed to create pull request", err)
			}
			log.Success("Pull request created successfully")
			return nil
		},
	}

	cmd.Flags().BoolVarP(&draft, "draft", "d", false, "Create a draft pull request")

	return cmd
}

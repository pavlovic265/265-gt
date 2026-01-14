package pullrequests

import (
	"github.com/pavlovic265/265-gt/client"
	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/executor"
	"github.com/pavlovic265/265-gt/utils/log"
	"github.com/spf13/cobra"
)

type createCommand struct {
	exe           executor.Executor
	configManager config.ConfigManager
}

func NewCreateCommand(
	exe executor.Executor,
	configManager config.ConfigManager,
) createCommand {
	return createCommand{
		exe:           exe,
		configManager: configManager,
	}
}

var draft bool

func (svc createCommand) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "create",
		Aliases: []string{"c"},
		Short:   "Create a pull request",
		RunE: func(cmd *cobra.Command, args []string) error {
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

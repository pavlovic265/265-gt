package auth

import (
	"github.com/pavlovic265/265-gt/client"
	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/utils/log"
	"github.com/spf13/cobra"
)

type statusCommand struct {
	configManager config.ConfigManager
}

func NewStatusCommand(
	configManager config.ConfigManager,
) statusCommand {
	return statusCommand{
		configManager: configManager,
	}
}

func (svc statusCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:                "status",
		Aliases:            []string{"st"},
		Short:              "see status of current auth user",
		DisableFlagParsing: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := config.RequireGlobal(cmd.Context())
			if err != nil {
				return err
			}

			if cfg.Global.ActiveAccount == nil || cfg.Global.ActiveAccount.User == "" {
				return log.ErrorMsg("No active account found")
			}
			account := cfg.Global.ActiveAccount

			err = client.Client[account.Platform].AuthStatus(cmd.Context())
			if err != nil {
				return log.Error("Authentication failed", err)
			}

			return nil
		},
	}
}

package auth

import (
	"github.com/pavlovic265/265-gt/client"
	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/utils/log"
	"github.com/spf13/cobra"
)

type statusCommand struct {
	configManager config.ConfigManager
	cliClient     client.CliClient
}

func NewStatusCommand(
	configManager config.ConfigManager,
	cliClient client.CliClient,
) statusCommand {
	return statusCommand{
		configManager: configManager,
		cliClient:     cliClient,
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
				return log.ErrorMsg("no active account found")
			}

			err = svc.cliClient.AuthStatus(cmd.Context())
			if err != nil {
				return log.Error("authentication failed", err)
			}

			return nil
		},
	}
}

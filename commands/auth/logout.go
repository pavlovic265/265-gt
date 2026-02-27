package auth

import (
	"github.com/pavlovic265/265-gt/client"
	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/utils/log"
	"github.com/spf13/cobra"
)

type logoutCommand struct {
	configManager config.ConfigManager
	cliClient     client.CliClient
}

func NewLogoutCommand(
	configManager config.ConfigManager,
	cliClient client.CliClient,
) logoutCommand {
	return logoutCommand{
		configManager: configManager,
		cliClient:     cliClient,
	}
}

func (svc logoutCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:                "logout",
		Aliases:            []string{"lo"},
		Short:              "logout user with token",
		DisableFlagParsing: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := config.RequireGlobal(cmd.Context())
			if err != nil {
				return err
			}

			if cfg.Global.ActiveAccount == nil || cfg.Global.ActiveAccount.User == "" {
				return log.ErrorMsg("no active account found")
			}
			account := cfg.Global.ActiveAccount

			if err := svc.cliClient.AuthLogout(cmd.Context(), account.User); err != nil {
				return log.Error("logout failed", err)
			}

			log.Successf("Successfully logged out from %s", account.User)
			return nil
		},
	}
}

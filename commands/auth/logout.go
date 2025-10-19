package auth

import (
	"github.com/pavlovic265/265-gt/client"
	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/executor"
	"github.com/pavlovic265/265-gt/utils/log"
	"github.com/spf13/cobra"
)

type logoutCommand struct {
	exe           executor.Executor
	configManager config.ConfigManager
}

func NewLogoutCommand(
	exe executor.Executor,
	configManager config.ConfigManager,
) logoutCommand {
	return logoutCommand{
		exe:           exe,
		configManager: configManager,
	}
}

func (svc logoutCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:                "logout",
		Aliases:            []string{"lo"},
		Short:              "logout user with token",
		DisableFlagParsing: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			account := svc.configManager.GetActiveAccount()
			if !svc.configManager.HasActiveAccount() {
				return log.ErrorMsg("No active account found")
			}

			if err := client.Client[account.Platform].AuthLogout(account.User); err != nil {
				return log.Error("Logout failed", err)
			}

			log.Success("Successfully logged out from " + account.User)
			return nil
		},
	}
}

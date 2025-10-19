package auth

import (
	"github.com/pavlovic265/265-gt/client"
	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/executor"
	"github.com/pavlovic265/265-gt/utils/log"
	"github.com/spf13/cobra"
)

type statusCommand struct {
	exe           executor.Executor
	configManager config.ConfigManager
}

func NewStatusCommand(
	exe executor.Executor,
	configManager config.ConfigManager,
) statusCommand {
	return statusCommand{
		exe:           exe,
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
			if !svc.configManager.HasActiveAccount() {
				return log.ErrorMsg("No active account found")
			}
			account := svc.configManager.GetActiveAccount()

			err := client.Client[account.Platform].AuthStatus()
			if err != nil {
				return log.Error("Authentication failed", err)
			}

			return nil
		},
	}
}

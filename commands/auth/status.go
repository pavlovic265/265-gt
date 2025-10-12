package auth

import (
	"fmt"

	"github.com/pavlovic265/265-gt/client"
	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/executor"
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
			fmt.Println("Checking status...")
			if !svc.configManager.HasActiveAccount() {
				fmt.Println("No active account found")
				return nil
			}
			account := svc.configManager.GetActiveAccount()

			err := client.Client[account.Platform].AuthStatus()
			if err != nil {
				fmt.Println("Authentication failed")
				return err
			}

			fmt.Println("Authentication successful")
			return nil
		},
	}
}

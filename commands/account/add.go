package account

import (
	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/executor"
	"github.com/pavlovic265/265-gt/utils/log"
	"github.com/spf13/cobra"
)

type addCommand struct {
	exe           executor.Executor
	configManager config.ConfigManager
}

func NewAddCommand(
	exe executor.Executor,
	configManager config.ConfigManager,
) addCommand {
	return addCommand{
		exe:           exe,
		configManager: configManager,
	}
}

func (ac addCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:   "add",
		Short: "Add a new account",
		Long:  "Add a new GitHub or GitLab account interactively",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Load existing config
			globalConfig, err := ac.configManager.LoadGlobalConfig()
			if err != nil {
				return log.Error("Global config not found. Run 'gt config global' to create it first", err)
			}

			// Run the account form
			accounts, err := HandleAddAccunts()
			if err != nil {
				return log.Error("Failed to add accounts", err)
			}

			if len(accounts) == 0 {
				log.Info("No accounts added")
				return nil
			}

			// Append new accounts
			globalConfig.Accounts = append(globalConfig.Accounts, accounts...)

			// Save config
			if err := ac.configManager.SaveGlobalConfig(*globalConfig); err != nil {
				return log.Error("Failed to save config", err)
			}

			log.Successf("Added %d account(s)", len(accounts))
			return nil
		},
	}
}

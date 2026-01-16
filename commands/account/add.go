package account

import (
	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/runner"
	"github.com/pavlovic265/265-gt/utils/log"
	"github.com/spf13/cobra"
)

type addCommand struct {
	runner        runner.Runner
	configManager config.ConfigManager
}

func NewAddCommand(
	runner runner.Runner,
	configManager config.ConfigManager,
) addCommand {
	return addCommand{
		runner:        runner,
		configManager: configManager,
	}
}

func (ac addCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:   "add",
		Short: "Add a new account",
		Long:  "Add a new GitHub or GitLab account interactively",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := config.RequireGlobal(cmd.Context())
			if err != nil {
				return err
			}

			accounts, err := HandleAddAccounts()
			if err != nil {
				return log.Error("failed to add accounts", err)
			}

			if len(accounts) == 0 {
				log.Info("No accounts added")
				return nil
			}

			// Append new accounts to context - will be saved by PersistentPostRunE
			cfg.Global.Accounts = append(cfg.Global.Accounts, accounts...)
			cfg.MarkDirty()

			log.Successf("Added %d account(s)", len(accounts))
			return nil
		},
	}
}

package account

import (
	"fmt"
	"slices"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/pavlovic265/265-gt/components"
	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/utils/log"
	"github.com/spf13/cobra"
)

type removeCommand struct {
	configManager config.ConfigManager
}

func NewRemoveCommand(
	configManager config.ConfigManager,
) removeCommand {
	return removeCommand{
		configManager: configManager,
	}
}

func (rc removeCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:     "remove",
		Aliases: []string{"rm"},
		Short:   "Remove an account",
		Long:    "Select and remove an existing GitHub or GitLab account",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Load existing config
			globalConfig, err := rc.configManager.LoadGlobalConfig()
			if err != nil {
				return log.Error("Global config not found. Run 'gt config global' to create it first", err)
			}

			if len(globalConfig.Accounts) == 0 {
				log.Info("No accounts configured")
				fmt.Println("\nRun 'gt account add' to add an account")
				return nil
			}

			// Build choices for search and select
			var choices []string
			for _, account := range globalConfig.Accounts {
				platform := account.Platform.String()
				choices = append(choices, fmt.Sprintf("%s (%s) - %s", account.User, platform, account.Email))
			}

			// Show account selector
			selectModel := components.ListModel{
				AllChoices: choices,
				Choices:    choices,
				Cursor:     0,
			}

			selectProgram := tea.NewProgram(selectModel)
			m, err := selectProgram.Run()
			if err != nil {
				return log.Error("Failed to select account", err)
			}

			selectedIndex := -1
			if m, ok := m.(components.ListModel); ok {
				if m.Selected == "" {
					log.Info("No account selected")
					return nil
				}

				// Find the index of the selected account
				for i, choice := range choices {
					if choice == m.Selected {
						selectedIndex = i
						break
					}
				}
			}

			if selectedIndex == -1 {
				return log.Error("Failed to select account", fmt.Errorf("invalid selection"))
			}

			// Get the account to remove
			accountToRemove := globalConfig.Accounts[selectedIndex]

			// Remove the account from the list
			globalConfig.Accounts = slices.Delete(globalConfig.Accounts, selectedIndex, selectedIndex+1)

			// If the removed account was the active account, clear it
			activeAccount := rc.configManager.GetActiveAccount()
			if activeAccount.User == accountToRemove.User && activeAccount.Platform == accountToRemove.Platform {
				globalConfig.ActiveAccount = &config.Account{}
				log.Info("Removed active account. Run 'gt auth' to set a new active account")
			}

			// Save config
			if err := rc.configManager.SaveGlobalConfig(*globalConfig); err != nil {
				return log.Error("Failed to save config", err)
			}

			log.Successf("Removed account: %s (%s)", accountToRemove.User, accountToRemove.Platform.String())
			return nil
		},
	}
}

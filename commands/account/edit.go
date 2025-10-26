package account

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/pavlovic265/265-gt/components"
	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/executor"
	"github.com/pavlovic265/265-gt/utils/log"
	"github.com/spf13/cobra"
)

type editCommand struct {
	exe           executor.Executor
	configManager config.ConfigManager
}

func NewEditCommand(
	exe executor.Executor,
	configManager config.ConfigManager,
) editCommand {
	return editCommand{
		exe:           exe,
		configManager: configManager,
	}
}

func (ec editCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:   "edit",
		Short: "Edit an existing account",
		Long:  "Select and edit an existing GitHub or GitLab account",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Load existing config
			globalConfig, err := ec.configManager.LoadGlobalConfig()
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

			var selectedIndex = -1
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

			// Edit the selected account
			selectedAccount := globalConfig.Accounts[selectedIndex]
			editedAccount, err := HandleEditAccount(&selectedAccount)
			if err != nil {
				return log.Error("Failed to edit account", err)
			}

			// Update the account in config
			globalConfig.Accounts[selectedIndex] = *editedAccount

			// Update active account if it was the one being edited
			activeAccount := ec.configManager.GetActiveAccount()
			if activeAccount.User == selectedAccount.User && activeAccount.Platform == selectedAccount.Platform {
				globalConfig.ActiveAccount = editedAccount
			}

			// Save config
			if err := ec.configManager.SaveGlobalConfig(*globalConfig); err != nil {
				return log.Error("Failed to save config", err)
			}

			log.Success("Account updated successfully")
			return nil
		},
	}
}

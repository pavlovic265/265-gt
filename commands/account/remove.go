package account

import (
	"fmt"
	"slices"
	"strings"

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
			cfg, err := config.RequireGlobal(cmd.Context())
			if err != nil {
				return err
			}

			if len(cfg.Global.Accounts) == 0 {
				log.Info("No accounts configured")
				fmt.Println("\nRun 'gt account add' to add an account")
				return nil
			}

			// Build choices for search and select
			var choices []string
			for _, account := range cfg.Global.Accounts {
				platform := account.Platform.String()
				choices = append(choices, fmt.Sprintf("%s (%s) - %s", account.User, platform, account.Email))
			}

			// Show account selector
			selectModel := components.ListModel[string]{
				AllChoices: choices,
				Choices:    choices,
				Cursor:     0,
				Formatter:  func(s string) string { return s },
				Matcher:    func(s, query string) bool { return strings.Contains(s, query) },
			}

			selectProgram := tea.NewProgram(selectModel)
			m, err := selectProgram.Run()
			if err != nil {
				return log.Error("Failed to select account", err)
			}

			selectedIndex := -1
			if m, ok := m.(components.ListModel[string]); ok {
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
			accountToRemove := cfg.Global.Accounts[selectedIndex]

			// Remove the account from the list - will be saved by PersistentPostRunE
			cfg.Global.Accounts = slices.Delete(cfg.Global.Accounts, selectedIndex, selectedIndex+1)

			// If the removed account was the active account, clear it
			if cfg.Global.ActiveAccount != nil &&
				cfg.Global.ActiveAccount.User == accountToRemove.User &&
				cfg.Global.ActiveAccount.Platform == accountToRemove.Platform {
				cfg.Global.ActiveAccount = nil
				log.Info("Removed active account. Run 'gt auth' to set a new active account")
			}
			cfg.MarkDirty()

			log.Successf("Removed account: %s (%s)", accountToRemove.User, accountToRemove.Platform.String())
			return nil
		},
	}
}

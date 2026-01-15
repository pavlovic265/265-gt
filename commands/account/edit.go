package account

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/pavlovic265/265-gt/components"
	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/runner"
	"github.com/pavlovic265/265-gt/utils/log"
	"github.com/spf13/cobra"
)

type editCommand struct {
	runner        runner.Runner
	configManager config.ConfigManager
}

func NewEditCommand(
	runner runner.Runner,
	configManager config.ConfigManager,
) editCommand {
	return editCommand{
		runner:        runner,
		configManager: configManager,
	}
}

func (ec editCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:   "edit",
		Short: "Edit an existing account",
		Long:  "Select and edit an existing GitHub or GitLab account",
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

			var choices []string
			for _, account := range cfg.Global.Accounts {
				platform := account.Platform.String()
				choices = append(choices, fmt.Sprintf("%s (%s) - %s", account.User, platform, account.Email))
			}

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

			var selectedIndex = -1
			if m, ok := m.(components.ListModel[string]); ok {
				if m.Selected == "" {
					log.Info("No account selected")
					return nil
				}

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

			selectedAccount := cfg.Global.Accounts[selectedIndex]
			editedAccount, err := HandleEditAccount(&selectedAccount)
			if err != nil {
				return log.Error("Failed to edit account", err)
			}

			// Update the account in context - will be saved by PersistentPostRunE
			cfg.Global.Accounts[selectedIndex] = *editedAccount

			// Update active account if it was the one being edited
			if cfg.Global.ActiveAccount != nil &&
				cfg.Global.ActiveAccount.User == selectedAccount.User &&
				cfg.Global.ActiveAccount.Platform == selectedAccount.Platform {
				cfg.Global.ActiveAccount = editedAccount
			}
			cfg.MarkDirty()

			log.Success("Account updated successfully")
			return nil
		},
	}
}

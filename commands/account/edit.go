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
	cmd := &cobra.Command{
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

			// Check flags
			tokenFlag, _ := cmd.Flags().GetBool("token")
			gpgFlag, _ := cmd.Flags().GetBool("gpg")

			// Select account
			selectedIndex, err := selectAccount(cfg.Global.Accounts)
			if err != nil {
				return err
			}
			if selectedIndex == -1 {
				log.Info("No account selected")
				return nil
			}

			selectedAccount := &cfg.Global.Accounts[selectedIndex]

			// Handle quick update flags
			if tokenFlag {
				if err := HandleTokenSetup(selectedAccount); err != nil {
					return log.Error("failed to update token", err)
				}
				cfg.MarkDirty()
				log.Success("Account updated successfully")
				return nil
			}

			if gpgFlag {
				if err := HandleGPGSetup(selectedAccount); err != nil {
					return log.Error("failed to update GPG key", err)
				}
				cfg.MarkDirty()
				log.Success("Account updated successfully")
				return nil
			}

			// Full edit flow
			editedAccount, err := HandleEditAccount(selectedAccount)
			if err != nil {
				return log.Error("failed to edit account", err)
			}

			// Offer SSH setup if not configured
			if editedAccount.SSHKeyPath == "" || editedAccount.SSHHost == "" {
				if err := HandleSSHSetup(editedAccount, ec.runner); err != nil {
					log.Warningf("SSH setup failed: %v", err)
				}
			}

			// Offer token setup if not configured
			if editedAccount.Token == "" {
				if err := HandleTokenSetup(editedAccount); err != nil {
					log.Warningf("Token setup failed: %v", err)
				}
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

	cmd.Flags().BoolP("token", "t", false, "Update token only")
	cmd.Flags().Bool("gpg", false, "Update GPG signing key only")

	return cmd
}

func selectAccount(accounts []config.Account) (int, error) {
	var choices []string
	for _, account := range accounts {
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
		return -1, log.Error("failed to select account", err)
	}

	if m, ok := m.(components.ListModel[string]); ok {
		if m.Selected == "" {
			return -1, nil
		}

		for i, choice := range choices {
			if choice == m.Selected {
				return i, nil
			}
		}
	}

	return -1, log.ErrorMsg("failed to select account: invalid selection")
}

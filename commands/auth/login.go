package auth

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/pavlovic265/265-gt/client"
	"github.com/pavlovic265/265-gt/components"
	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/executor"
	"github.com/pavlovic265/265-gt/utils/log"
	"github.com/spf13/cobra"
)

type loginCommand struct {
	exe           executor.Executor
	configManager config.ConfigManager
}

func NewLoginCommand(
	exe executor.Executor,
	configManager config.ConfigManager,
) loginCommand {
	return loginCommand{
		exe:           exe,
		configManager: configManager,
	}
}

func (svc loginCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:                "login",
		Aliases:            []string{"lo"},
		Short:              "login user with token",
		DisableFlagParsing: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			accounts := svc.configManager.GetAccounts()

			var users []string
			for _, acc := range accounts {
				users = append(users, acc.User)
			}

			initialModel := components.ListModel{
				AllChoices: users,
				Choices:    users,
				Cursor:     0,
				Query:      "",
			}

			program := tea.NewProgram(initialModel)

			if finalModel, err := program.Run(); err == nil {
				if m, ok := finalModel.(components.ListModel); ok && m.Selected != "" {
					var account config.Account
					for _, acc := range accounts {
						if acc.User == m.Selected {
							account = acc
							break
						}
					}
					if err := client.Client[account.Platform].AuthLogin(account.User); err != nil {
						return log.Error("Authentication failed", err)
					}

					log.Success("Successfully authenticated with " + m.Selected)
				} else {
					return log.ErrorMsg("No user selected for authentication")
				}
			} else {
				return log.Error("Failed to display user selection menu", err)
			}
			return nil
		},
	}
}

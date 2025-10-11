package auth

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/pavlovic265/265-gt/client"
	"github.com/pavlovic265/265-gt/components"
	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/constants"
	"github.com/pavlovic265/265-gt/executor"
	"github.com/spf13/cobra"
)

type loginCommand struct {
	exe executor.Executor
}

func NewLoginCommand(
	exe executor.Executor,
) loginCommand {
	return loginCommand{
		exe: exe,
	}
}

func (svc loginCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:                "login",
		Aliases:            []string{"lo"},
		Short:              "login user with token",
		DisableFlagParsing: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			accounts := config.GlobalConfig.Accounts

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
					fmt.Println("Authentication started for", m.Selected)
					var account config.Account
					for _, acc := range accounts {
						if acc.User == m.Selected {
							account = acc
							break
						}
					}
					if err := client.Client[account.Platform].AuthLogin(account.User); err != nil {
						return err
					}

					fmt.Println(constants.SuccessIndicator("Successfully authenticated with " + m.Selected))
				}
			} else {
				return err
			}
			return nil
		},
	}
}

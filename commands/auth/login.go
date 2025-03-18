package auth

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/pavlovic265/265-gt/client"
	"github.com/pavlovic265/265-gt/components"
	"github.com/pavlovic265/265-gt/config"
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
			accoutns := config.Config.GlobalConfig.GitHub.Accounts

			var users []string
			for _, acc := range accoutns {
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
					if err := client.GlobalClient.AuthLogin(m.Selected); err != nil {
						return err
					}

					fmt.Println("Successfully authenticated with", m.Selected)
				}
			} else {
				return err
			}
			return nil
		},
	}
}

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

type logoutCommand struct {
	exe executor.Executor
}

func NewLogoutCommand(
	exe executor.Executor,
) logoutCommand {
	return logoutCommand{
		exe: exe,
	}
}

func (svc logoutCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:                "logout",
		Aliases:            []string{"lo"},
		Short:              "logout user with token",
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
					if err := client.GlobalClient.AuthLogout(m.Selected); err != nil {
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

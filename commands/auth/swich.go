package auth

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/pavlovic265/265-gt/components"
	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/executor"
	"github.com/spf13/cobra"
)

type swichCommand struct {
	exe executor.Executor
}

func NewSwichCommand(
	exe executor.Executor,
) swichCommand {
	return swichCommand{
		exe: exe,
	}
}

func (svc swichCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:     "swich",
		Aliases: []string{"sw"},
		Short:   "swich accounts",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				if err := svc.swichUser(args[0]); err != nil {
					return err
				}
			} else {
				if err := svc.selectAndSwichUser(); err != nil {
					return err
				}
			}
			return nil
		},
	}
}

func (svc swichCommand) swichUser(user string) error {
	var token string
	acocunts := config.GlobalConfig.GitHub.Accounts
	for _, acc := range acocunts {
		if user == acc.User {
			token = acc.Token
			break
		}
	}
	if token == "" {
		return fmt.Errorf("user (%s) does not exits in config", user)
	}

	exeArgs := []string{"auth", "login", "--with-token"}
	err := svc.exe.ExecuteWithStdin("gt", token, exeArgs...)
	if err != nil {
		return err
	}
	return nil
}

func (svc swichCommand) selectAndSwichUser() error {
	acocunts := config.GlobalConfig.GitHub.Accounts
	var users []string
	for _, acc := range acocunts {
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
			fmt.Println("Swiching accounts...")
			err := svc.swichUser(m.Selected)
			if err != nil {
				return err
			}
			fmt.Printf("Swiched to %s\n", m.Selected)
		}
	} else {
		return err
	}
	return nil
}

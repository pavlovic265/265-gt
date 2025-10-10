package auth

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/pavlovic265/265-gt/components"
	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/executor"
	"github.com/pavlovic265/265-gt/utils/pointer"
	"github.com/spf13/cobra"
)

type switchCommand struct {
	exe executor.Executor
}

func NewSwitchCommand(
	exe executor.Executor,
) switchCommand {
	return switchCommand{
		exe: exe,
	}
}

func (svc switchCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:     "switch",
		Aliases: []string{"sw"},
		Short:   "switch accounts",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				if err := svc.switchUser(args[0]); err != nil {
					return err
				}
			} else {
				if err := svc.selectAndswitchUser(); err != nil {
					return err
				}
			}
			return nil
		},
	}
}

func (svc switchCommand) switchUser(user string) error {
	accounts := config.GlobalConfig.Accounts
	var account *config.Account
	for _, acc := range accounts {
		if user == acc.User {
			account = pointer.From(acc)
			break
		}
	}
	if account == nil {
		return fmt.Errorf("user (%s) does not exits in config", user)
	}

	exeArgs := []string{"auth", "switch", "--user", account.User}
	err := svc.exe.WithGh().WithArgs(exeArgs).Run()
	if err != nil {
		return err
	}

	err = config.SetActiveAccount(pointer.Deref(account))
	if err != nil {
		return err
	}

	return nil
}

func (svc switchCommand) selectAndswitchUser() error {
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
			err := svc.switchUser(m.Selected)
			if err != nil {
				return err
			}
		}
	} else {
		return err
	}
	return nil
}

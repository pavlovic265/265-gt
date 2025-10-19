package auth

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/pavlovic265/265-gt/components"
	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/executor"
	"github.com/pavlovic265/265-gt/utils/log"
	"github.com/pavlovic265/265-gt/utils/pointer"
	"github.com/spf13/cobra"
)

type switchCommand struct {
	exe           executor.Executor
	configManager config.ConfigManager
}

func NewSwitchCommand(
	exe executor.Executor,
	configManager config.ConfigManager,
) switchCommand {
	return switchCommand{
		exe:           exe,
		configManager: configManager,
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
	accounts := svc.configManager.GetAccounts()
	var account *config.Account
	for _, acc := range accounts {
		if user == acc.User {
			account = pointer.From(acc)
			break
		}
	}
	if account == nil {
		log.Warning("User '" + user + "' does not exist in config")
		return nil
	}

	exeArgs := []string{"auth", "switch", "--user", account.User}
	err := svc.exe.WithGh().WithArgs(exeArgs).Run()
	if err != nil {
		return log.Error("Failed to switch account", err)
	}

	err = svc.configManager.SetActiveAccount(pointer.Deref(account))
	if err != nil {
		return log.Error("Failed to update active account", err)
	}

	log.Success("Switched to account: " + account.User)
	return nil
}

func (svc switchCommand) selectAndswitchUser() error {
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
			err := svc.switchUser(m.Selected)
			if err != nil {
				return err
			}
		} else {
			return log.ErrorMsg("No user selected for switching")
		}
	} else {
		return log.Error("Failed to display user selection menu", err)
	}
	return nil
}

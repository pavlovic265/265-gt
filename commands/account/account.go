package account

import (
	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/executor"
	"github.com/spf13/cobra"
)

type accountCommand struct {
	exe           executor.Executor
	configManager config.ConfigManager
}

func NewAccountCommand(
	exe executor.Executor,
	configManager config.ConfigManager,
) accountCommand {
	return accountCommand{
		exe:           exe,
		configManager: configManager,
	}
}

func (ac accountCommand) Command() *cobra.Command {
	accountCmd := &cobra.Command{
		Use:     "account",
		Aliases: []string{"acc"},
		Short:   "Manage accounts",
		Long:    "Manage your GitHub and GitLab accounts for gt",
	}

	accountCmd.AddCommand(NewAddCommand(ac.exe, ac.configManager).Command())
	accountCmd.AddCommand(NewListCommand(ac.configManager).Command())
	accountCmd.AddCommand(NewEditCommand(ac.exe, ac.configManager).Command())
	accountCmd.AddCommand(NewRemoveCommand(ac.configManager).Command())

	return accountCmd
}

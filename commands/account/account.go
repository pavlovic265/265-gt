package account

import (
	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/runner"
	"github.com/spf13/cobra"
)

type accountCommand struct {
	runner        runner.Runner
	configManager config.ConfigManager
}

func NewAccountCommand(
	runner runner.Runner,
	configManager config.ConfigManager,
) accountCommand {
	return accountCommand{
		runner:        runner,
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

	accountCmd.AddCommand(NewAddCommand(ac.runner, ac.configManager).Command())
	accountCmd.AddCommand(NewListCommand(ac.configManager).Command())
	accountCmd.AddCommand(NewEditCommand(ac.runner, ac.configManager).Command())
	accountCmd.AddCommand(NewRemoveCommand(ac.configManager).Command())
	accountCmd.AddCommand(NewAttachCommand(ac.configManager).Command())

	return accountCmd
}

package createconfig

import (
	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/runner"
	"github.com/spf13/cobra"
)

type configCommand struct {
	runner        runner.Runner
	configManager config.ConfigManager
}

func NewConfigCommand(
	runner runner.Runner,
	configManager config.ConfigManager,
) configCommand {
	return configCommand{
		runner:        runner,
		configManager: configManager,
	}
}

func (svc configCommand) Command() *cobra.Command {
	configCmd := &cobra.Command{
		Use:     "config",
		Aliases: []string{"conf"},
		Short:   "create config",
	}
	configCmd.AddCommand(NewGlobalCommand(svc.runner, svc.configManager).Command())
	configCmd.AddCommand(NewLocalCommand(svc.runner, svc.configManager).Command())

	return configCmd
}

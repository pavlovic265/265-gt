package createconfig

import (
	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/executor"
	"github.com/spf13/cobra"
)

type configCommand struct {
	runner        executor.Runner
	configManager config.ConfigManager
}

func NewConfigCommand(
	runner executor.Runner,
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

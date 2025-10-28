package createconfig

import (
	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/executor"
	"github.com/spf13/cobra"
)

type configCommand struct {
	exe           executor.Executor
	configManager config.ConfigManager
}

func NewConfigCommand(
	exe executor.Executor,
	configManager config.ConfigManager,
) configCommand {
	return configCommand{
		exe:           exe,
		configManager: configManager,
	}
}

func (svc configCommand) Command() *cobra.Command {
	configCmd := &cobra.Command{
		Use:     "config",
		Aliases: []string{"conf"},
		Short:   "create config",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			svc.configManager.InitGlobalConfig()
		},
	}
	configCmd.AddCommand(NewGlobalCommand(svc.exe, svc.configManager).Command())
	configCmd.AddCommand(NewLocalCommand(svc.exe, svc.configManager).Command())

	return configCmd
}

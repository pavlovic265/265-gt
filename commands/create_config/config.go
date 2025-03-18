package createconfig

import (
	"github.com/pavlovic265/265-gt/executor"
	"github.com/spf13/cobra"
)

type configCommand struct {
	exe executor.Executor
}

func NewConfigCommand(
	exe executor.Executor,
) configCommand {
	return configCommand{
		exe: exe,
	}
}

func (svc configCommand) Command() *cobra.Command {
	configCmd := &cobra.Command{
		Use:     "config",
		Aliases: []string{"conf"},
		Short:   "create config",
	}
	configCmd.AddCommand(NewGlobalCommand(svc.exe).Command())

	return configCmd
}

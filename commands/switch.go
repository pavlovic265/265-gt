package commands

import (
	"github.com/pavlovic265/265-gt/executor"
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
		Short:   "switch back to previous branch",
		RunE: func(cmd *cobra.Command, args []string) error {
			exeArgs := append([]string{"checkout", "-"}, args...)
			_, err := svc.exe.Execute("git", exeArgs...)
			return err
		},
	}
}

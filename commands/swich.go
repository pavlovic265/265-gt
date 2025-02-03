package commands

import (
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

func (svc *swichCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:     "swich",
		Aliases: []string{"sw"},
		Short:   "swich back to previous branch",
		RunE: func(cmd *cobra.Command, args []string) error {
			exeArgs := append([]string{"checkout", "-"}, args...)

			return svc.exe.Execute("git", exeArgs...)
		},
	}
}

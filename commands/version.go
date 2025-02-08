package commands

import (
	"github.com/pavlovic265/265-gt/executor"
	"github.com/spf13/cobra"
)

type versionCommand struct {
	exe executor.Executor
}

func NewVersionCommand(
	exe executor.Executor,
) versionCommand {
	return versionCommand{
		exe: exe,
	}
}

func (svc versionCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "version of current build",
		RunE: func(cmd *cobra.Command, args []string) error {
			exeArgs := append([]string{"checkout", "-"}, args...)

			return svc.exe.Execute("git", exeArgs...)
		},
	}
}

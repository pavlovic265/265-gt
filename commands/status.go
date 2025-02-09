package commands

import (
	"github.com/pavlovic265/265-gt/executor"
	"github.com/spf13/cobra"
)

type statusCommand struct {
	exe executor.Executor
}

func NewStatusCommand(
	exe executor.Executor,
) statusCommand {
	return statusCommand{
		exe: exe,
	}
}

func (svc statusCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:                "status",
		Aliases:            []string{"st"},
		Short:              "git status",
		DisableFlagParsing: true,
		SilenceUsage:       true,
		RunE: func(cmd *cobra.Command, args []string) error {
			exeArgs := append([]string{"status"}, args...)

			_, err := svc.exe.Execute("git", exeArgs...)
			if err != nil {
				return err
			}

			return nil
		},
	}
}

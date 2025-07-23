package commands

import (
	"github.com/pavlovic265/265-gt/executor"
	"github.com/spf13/cobra"
)

type unstageCommand struct {
	exe executor.Executor
}

func NewUnstageCommand(
	exe executor.Executor,
) unstageCommand {
	return unstageCommand{
		exe: exe,
	}
}

func (svc unstageCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:     "unstage",
		Aliases: []string{"us"},
		Short:   "unstage ",
		RunE: func(cmd *cobra.Command, args []string) error {
			exeArgs := append([]string{"restore", "--staged"}, args...)
			err := svc.exe.WithGit().WithArgs(exeArgs).Run()
			return err
		},
	}
}

package commit

import (
	"time"

	"github.com/pavlovic265/265-gt/executor"
	"github.com/spf13/cobra"
)

type createCommand struct {
	exe executor.Executor
}

func NewCreateCommand(
	exe executor.Executor,
) createCommand {
	return createCommand{
		exe: exe,
	}
}

func (svc createCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:     "create",
		Aliases: []string{"c"},
		Short:   "create commit",
		RunE: func(cmd *cobra.Command, args []string) error {
			var message string
			if len(args) == 0 {
				message = "New commit: " + time.Now().Format("02-Jan-2006 15:04:05")
			}

			exeArgs := []string{"commit", "-am", message}
			err := svc.exe.WithGit().WithArgs(exeArgs).Run()
			if err != nil {
				return err
			}
			return nil
		},
	}
}

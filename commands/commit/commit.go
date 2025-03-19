package commit

import (
	"time"

	"github.com/pavlovic265/265-gt/executor"
	"github.com/spf13/cobra"
)

type commitCommand struct {
	exe executor.Executor
}

func NewCommitCommand(
	exe executor.Executor,
) commitCommand {
	return commitCommand{
		exe: exe,
	}
}

func (svc commitCommand) Command() *cobra.Command {
	commitCmd := &cobra.Command{
		Use:     "commit",
		Aliases: []string{"cm"},
		Short:   "create commit",
		RunE: func(cmd *cobra.Command, args []string) error {
			message := time.Now().Format("02-Jan-2006 15:04:05")
			if len(args) != 0 {
				message = string(args[0])
			}

			exeArgs := []string{"commit", "-am", message}
			err := svc.exe.WithGit().WithArgs(exeArgs).Run()
			if err != nil {
				return err
			}
			return nil
		},
	}

	commitCmd.AddCommand(NewEmptyCommand(svc.exe).Command())

	return commitCmd
}

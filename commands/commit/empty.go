package commit

import (
	"fmt"
	"time"

	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/executor"
	"github.com/spf13/cobra"
)

type emptyCommand struct {
	exe executor.Executor
}

func NewEmptyCommand(
	exe executor.Executor,
) emptyCommand {
	return emptyCommand{
		exe: exe,
	}
}

func (svc emptyCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:     "empty",
		Aliases: []string{"e"},
		Short:   "empty new commit",
		RunE: func(cmd *cobra.Command, args []string) error {
			message := "empty commit - " + time.Now().Format("02-Jan-2006 15:04:05")

			exeArgs := []string{"commit", "--allow-empty", "-m", message}
			err := svc.exe.WithGit().WithArgs(exeArgs).Run()
			if err != nil {
				return err
			}
			fmt.Println(config.SuccessIndicator("Empty commit created with message: '" + message + "'"))
			return nil
		},
	}
}

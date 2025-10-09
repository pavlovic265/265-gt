package commit

import (
	"fmt"

	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/executor"
	timeutils "github.com/pavlovic265/265-gt/utils/timeutils"
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
	var empty bool

	commitCmd := &cobra.Command{
		Use:     "commit",
		Aliases: []string{"cm"},
		Short:   "create commit",
		RunE: func(cmd *cobra.Command, args []string) error {
			message := timeutils.Now().Format(timeutils.LayoutUserFriendly)
			if len(args) != 0 {
				message = string(args[0])
			}

			exeArgs := []string{"commit", "-m", message}
			if empty {
				exeArgs = []string{"commit", "--allow-empty", "-m", message}
				message = "empty commit - " + message
			}

			err := svc.exe.WithGit().WithArgs(exeArgs).Run()
			if err != nil {
				return err
			}

			if empty {
				fmt.Println(config.SuccessIndicator("Empty commit created with message: '" + message + "'"))
			} else {
				fmt.Println(config.SuccessIndicator("Commit created with message: '" + message + "'"))
			}
			return nil
		},
	}

	commitCmd.Flags().BoolVarP(&empty, "empty", "e", false, "create an empty commit")

	return commitCmd
}

package commit

import (
	"fmt"

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
			if empty {
				return svc.handleEmptyCommit()
			}

			if len(args) == 0 {
				return fmt.Errorf("no message provided")
			}
			message := string(args[0])
			return svc.handleCommit(message)

		},
	}

	commitCmd.Flags().BoolVarP(&empty, "empty", "e", false, "create an empty commit")

	return commitCmd
}

func (svc commitCommand) handleEmptyCommit() error {
	message := timeutils.Now().Format(timeutils.LayoutUserFriendly)
	exeArgs := []string{"commit", "--allow-empty", "-m", message}
	err := svc.exe.WithGit().WithArgs(exeArgs).Run()
	if err != nil {
		return err
	}
	fmt.Println("Empty commit created with message: '" + message + "'")
	return nil
}

func (svc commitCommand) handleCommit(message string) error {
	exeArgs := []string{"commit", "-m", message}
	err := svc.exe.WithGit().WithArgs(exeArgs).Run()
	if err != nil {
		return err
	}
	fmt.Println("Commit created with message: '" + message + "'")
	return nil
}

package commands

import (
	"fmt"

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
			if err != nil {
				return err
			}
			if len(args) == 0 {
				fmt.Println("All staged changes unstaged")
			} else {
				fmt.Println("Files unstaged successfully")
			}
			return nil
		},
	}
}

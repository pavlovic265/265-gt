package commit

import (
	"fmt"

	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/executor"
	"github.com/pavlovic265/265-gt/utils"
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
			message := utils.Now().Format(utils.LayoutUserFriendly)
			if len(args) != 0 {
				message = string(args[0])
			}

			exeArgs := []string{"commit", "-m", message}
			err := svc.exe.WithGit().WithArgs(exeArgs).Run()
			if err != nil {
				return err
			}
			fmt.Println(config.SuccessIndicator("Commit created with message: '" + message + "'"))
			return nil
		},
	}

	commitCmd.AddCommand(NewEmptyCommand(svc.exe).Command())

	return commitCmd
}

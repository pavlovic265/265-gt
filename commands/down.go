package commands

import (
	"fmt"

	"github.com/pavlovic265/265-gt/executor"
	"github.com/pavlovic265/265-gt/helpers"
	pointer "github.com/pavlovic265/265-gt/utils/pointer"
	"github.com/spf13/cobra"
)

type downCommand struct {
	exe       executor.Executor
	gitHelper helpers.GitHelper
}

func NewDownCommand(
	exe executor.Executor,
	gitHelper helpers.GitHelper,
) downCommand {
	return downCommand{
		exe:       exe,
		gitHelper: gitHelper,
	}
}

func (svc downCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:   "down",
		Short: "move to brunch down in stack",
		RunE: func(cmd *cobra.Command, args []string) error {
			branch, err := svc.gitHelper.GetCurrentBranchName()
			if err != nil {
				return err
			}
			parent := svc.gitHelper.GetParent(pointer.Deref(branch))

			exeArgs := []string{"checkout", parent}
			err = svc.exe.WithGit().WithArgs(exeArgs).Run()
			if err != nil {
				return err
			}

			fmt.Println("Moved down to branch '" + parent + "'")
			return nil
		},
	}
}

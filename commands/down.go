package commands

import (
	"fmt"

	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/executor"
	"github.com/pavlovic265/265-gt/helpers"
	"github.com/pavlovic265/265-gt/utils"
	"github.com/spf13/cobra"
)

type downCommand struct {
	exe executor.Executor
}

func NewDownCommand(
	exe executor.Executor,
) downCommand {
	return downCommand{
		exe: exe,
	}
}

func (svc downCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:   "down",
		Short: "move to brunch down in stack",
		RunE: func(cmd *cobra.Command, args []string) error {
			branch, err := helpers.GetCurrentBranchName(svc.exe)
			if err != nil {
				return err
			}
			parent := helpers.GetParent(svc.exe, utils.Deref(branch))

			exeArgs := []string{"checkout", parent}
			err = svc.exe.WithGit().WithArgs(exeArgs).Run()
			if err != nil {
				return err
			}

			fmt.Println(config.SuccessIndicator("Moved down to branch '" + parent + "'"))
			return nil
		},
	}
}

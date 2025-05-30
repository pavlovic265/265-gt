package commands

import (
	"github.com/pavlovic265/265-gt/executor"
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
			branch, err := utils.GetCurrentBranchName(svc.exe)
			if err != nil {
				return err
			}
			parent := utils.GetParent(svc.exe, *branch)

			exeArgs := []string{"checkout", parent}
			err = svc.exe.WithGit().WithArgs(exeArgs).Run()
			if err != nil {
				return err
			}

			return nil
		},
	}
}

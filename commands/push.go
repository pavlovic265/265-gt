package commands

import (
	"github.com/pavlovic265/265-gt/executor"
	"github.com/pavlovic265/265-gt/utils"
	"github.com/spf13/cobra"
)

type pushCommand struct {
	exe executor.Executor
}

func NewPushCommand(
	exe executor.Executor,
) pushCommand {
	return pushCommand{
		exe: exe,
	}
}

func (svc pushCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:     "push",
		Aliases: []string{"pu"},
		Short:   "push branch always froce",
		RunE: func(cmd *cobra.Command, args []string) error {
			currentBranch, err := utils.GetCurrentBranchName(svc.exe)
			if err != nil {
				return err
			}

			exeArgs := []string{"push", "--force", "origin", *currentBranch}
			err = svc.exe.WithGit().WithArgs(exeArgs).Run()
			if err != nil {
				return err
			}

			return nil
		},
	}
}

package commands

import (
	"fmt"

	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/executor"
	"github.com/pavlovic265/265-gt/helpers"
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
			currentBranch, err := helpers.GetCurrentBranchName(svc.exe)
			if err != nil {
				return err
			}

			exeArgs := []string{"push", "--force", "origin", utils.Deref(currentBranch)}
			err = svc.exe.WithGit().WithArgs(exeArgs).Run()
			if err != nil {
				return err
			}

			fmt.Println(config.SuccessIndicator("Branch '" + utils.Deref(currentBranch) + "' pushed successfully"))
			return nil
		},
	}
}

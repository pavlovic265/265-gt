package commands

import (
	"fmt"

	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/executor"
	"github.com/pavlovic265/265-gt/helpers"
	pointer "github.com/pavlovic265/265-gt/utils/pointer"
	"github.com/spf13/cobra"
)

type pushCommand struct {
	exe       executor.Executor
	gitHelper helpers.GitHelper
}

func NewPushCommand(
	exe executor.Executor,
	gitHelper helpers.GitHelper,
) pushCommand {
	return pushCommand{
		exe:       exe,
		gitHelper: gitHelper,
	}
}

func (svc pushCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:     "push",
		Aliases: []string{"pu"},
		Short:   "push branch always froce",
		RunE: func(cmd *cobra.Command, args []string) error {
			currentBranch, err := svc.gitHelper.GetCurrentBranchName(svc.exe)
			if err != nil {
				return err
			}

			exeArgs := []string{"push", "--force", "origin", pointer.Deref(currentBranch)}
			err = svc.exe.WithGit().WithArgs(exeArgs).Run()
			if err != nil {
				return err
			}

			fmt.Println(config.SuccessIndicator("Branch '" + pointer.Deref(currentBranch) + "' pushed successfully"))
			return nil
		},
	}
}

package commands

import (
	"fmt"

	"github.com/pavlovic265/265-gt/executor"
	"github.com/pavlovic265/265-gt/helpers"
	pointer "github.com/pavlovic265/265-gt/utils/pointer"
	"github.com/spf13/cobra"
)

type pullCommand struct {
	exe       executor.Executor
	gitHelper helpers.GitHelper
}

func NewPullCommand(
	exe executor.Executor,
	gitHelper helpers.GitHelper,
) pullCommand {
	return pullCommand{
		exe:       exe,
		gitHelper: gitHelper,
	}
}

func (svc pullCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:     "pull",
		Aliases: []string{"pl"},
		Short:   "pull branch",
		RunE: func(cmd *cobra.Command, args []string) error {
			currentBranch, err := svc.gitHelper.GetCurrentBranchName()
			if err != nil {
				return err
			}

			exeArgs := []string{"pull", "origin", pointer.Deref(currentBranch)}
			err = svc.exe.WithGit().WithArgs(exeArgs).Run()
			if err != nil {
				return err
			}

			fmt.Println("Branch '" + pointer.Deref(currentBranch) + "' pulled successfully")
			return nil
		},
	}
}

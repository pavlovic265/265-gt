package commands

import (
	"fmt"

	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/executor"
	"github.com/pavlovic265/265-gt/helpers"
	pointer "github.com/pavlovic265/265-gt/utils/pointer"
	"github.com/spf13/cobra"
)

type pullCommand struct {
	exe executor.Executor
}

func NewPullCommand(
	exe executor.Executor,
) pullCommand {
	return pullCommand{
		exe: exe,
	}
}

func (svc pullCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:     "pull",
		Aliases: []string{"pl"},
		Short:   "pull branch",
		RunE: func(cmd *cobra.Command, args []string) error {
			currentBranch, err := helpers.GetCurrentBranchName(svc.exe)
			if err != nil {
				return err
			}

				exeArgs := []string{"pull", "origin", pointer.Deref(currentBranch)}
			err = svc.exe.WithGit().WithArgs(exeArgs).Run()
			if err != nil {
				return err
			}

			fmt.Println(config.SuccessIndicator("Branch '" + pointer.Deref(currentBranch) + "' pulled successfully"))
			return nil
		},
	}
}

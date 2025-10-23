package stack

import (
	"fmt"

	"github.com/pavlovic265/265-gt/executor"
	"github.com/pavlovic265/265-gt/helpers"
	"github.com/pavlovic265/265-gt/utils/pointer"
	"github.com/spf13/cobra"
)

type restackCommand struct {
	exe       executor.Executor
	gitHelper helpers.GitHelper
}

func NewRestackCommand(
	exe executor.Executor,
	gitHelper helpers.GitHelper,
) restackCommand {
	return restackCommand{
		exe:       exe,
		gitHelper: gitHelper,
	}
}

func (svc restackCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:     "restack",
		Aliases: []string{"r"},
		Short:   "Restack branches",
		RunE: func(cmd *cobra.Command, args []string) error {
			branch, err := svc.gitHelper.GetCurrentBranchName()
			if err != nil {
				return err
			}
			parent, err := svc.gitHelper.GetParent(pointer.Deref(branch))
			if err != nil {
				return err
			}

			fmt.Println(branch, parent)
			return nil
		},
	}
}

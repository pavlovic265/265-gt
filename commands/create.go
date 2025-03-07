package commands

import (
	"fmt"

	"github.com/pavlovic265/265-gt/components"
	"github.com/pavlovic265/265-gt/executor"
	"github.com/pavlovic265/265-gt/utils"
	"github.com/spf13/cobra"
)

type createCommand struct {
	exe executor.Executor
}

func NewCreateCommand(
	exe executor.Executor,
) createCommand {
	return createCommand{
		exe: exe,
	}
}

func (svc createCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:     "create",
		Aliases: []string{"c"},
		Short:   "create branch",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return fmt.Errorf("missing branch name")
			}
			branch := args[0]

			parent, err := utils.GetCurrentBranchName(svc.exe)
			if err != nil {
				return err
			}

			err = components.SetParent(svc.exe, *parent, branch)
			if err != nil {
				return err
			}

			exeArgs := []string{"checkout", "-b", branch}
			err = svc.exe.WithGit().WithArgs(exeArgs).Run()
			if err != nil {
				err := components.DeleteParent(svc.exe, branch)
				if err != nil {
					return err
				}

				return err
			}
			return nil
		},
	}
}

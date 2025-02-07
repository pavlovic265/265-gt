package commands

import (
	"fmt"

	"github.com/pavlovic265/265-gt/executor"
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
			exeArgs := []string{"checkout", "-b", args[0]}

			return svc.exe.Execute("git", exeArgs...)
		},
	}
}

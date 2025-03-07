package branch

import (
	"github.com/pavlovic265/265-gt/executor"
	"github.com/spf13/cobra"
)

type branchCommand struct {
	exe executor.Executor
}

func NewBranchCommand(
	exe executor.Executor,
) branchCommand {
	return branchCommand{
		exe: exe,
	}
}

func (svc branchCommand) Command() *cobra.Command {
	branchCmd := &cobra.Command{
		Use:     "branch",
		Aliases: []string{"br"},
		Short:   "git branch",
	}
	branchCmd.AddCommand(NewDeleteCommand(svc.exe).Command())

	return branchCmd
}

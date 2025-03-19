package commit

import (
	"github.com/pavlovic265/265-gt/executor"
	"github.com/spf13/cobra"
)

type commitCommand struct {
	exe executor.Executor
}

func NewCommitCommand(
	exe executor.Executor,
) commitCommand {
	return commitCommand{
		exe: exe,
	}
}

func (svc commitCommand) Command() *cobra.Command {
	commitCmd := &cobra.Command{
		Use:     "commit",
		Aliases: []string{"cm"},
		Short:   "create commit",
	}
	commitCmd.AddCommand(NewCreateCommand(svc.exe).Command())
	commitCmd.AddCommand(NewEmptyCommand(svc.exe).Command())

	return commitCmd
}

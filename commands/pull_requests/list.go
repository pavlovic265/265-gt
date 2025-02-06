package pullrequests

import (
	"github.com/pavlovic265/265-gt/executor"
	"github.com/spf13/cobra"
)

type listCommand struct {
	exe executor.Executor
}

func NewListCommand(
	exe executor.Executor,
) listCommand {
	return listCommand{
		exe: exe,
	}
}

func (svc listCommand) Command() *cobra.Command {
	pullRequestCmd := &cobra.Command{
		Use:     "list",
		Short:   "show list of pull request",
		Aliases: []string{"li"},
	}

	pullRequestCmd.AddCommand(NewCreateCommand(svc.exe).Command())

	return pullRequestCmd
}

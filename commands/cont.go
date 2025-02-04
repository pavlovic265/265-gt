package commands

import (
	"github.com/pavlovic265/265-gt/executor"
	"github.com/spf13/cobra"
)

type contCommand struct {
	exe executor.Executor
}

func NewContCommand(
	exe executor.Executor,
) contCommand {
	return contCommand{
		exe: exe,
	}
}

func (svc contCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:   "cont",
		Short: "short for rebase --contine",
		RunE: func(cmd *cobra.Command, args []string) error {
			exeArgs := []string{"rebase", "--continue"}

			return svc.exe.Execute("git", exeArgs...)
		},
	}
}

package commands

import (
	"fmt"

	"github.com/pavlovic265/265-gt/config"
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
			err := svc.exe.WithGit().WithArgs(exeArgs).Run()
			if err != nil {
				return err
			}
			fmt.Println(config.SuccessIndicator("Rebase continued successfully"))
			return nil
		},
	}
}

package commands

import (
	"fmt"

	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/executor"
	"github.com/spf13/cobra"
)

type updateCommand struct {
	exe executor.Executor
}

func NewUpdateCommand(
	exe executor.Executor,
) updateCommand {
	return updateCommand{
		exe: exe,
	}
}

func (svc updateCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:   "update",
		Short: "update of current build",
		RunE: func(cmd *cobra.Command, args []string) error {
			exeArgs := []string{"-c", `curl -fsSL https://raw.githubusercontent.com/pavlovic265/265-gt/main/scripts/install.sh | bash`}
			err := svc.exe.WithName("bash").WithArgs(exeArgs).Run()
			if err != nil {
				return err
			}
			fmt.Println(config.SuccessIndicator("Tool updated successfully"))
			return nil
		},
	}
}

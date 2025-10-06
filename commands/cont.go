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
	cmd := &cobra.Command{
		Use:   "cont",
		Short: "short for rebase --continue",
		RunE: func(cmd *cobra.Command, args []string) error {
			exeArgs := []string{"rebase", "--continue"}

			// Add --no-edit flag if requested
			if noEdit, _ := cmd.Flags().GetBool("no-edit"); noEdit {
				exeArgs = append(exeArgs, "--no-edit")
			}

			err := svc.exe.WithGit().WithArgs(exeArgs).Run()
			if err != nil {
				return err
			}
			fmt.Println(config.SuccessIndicator("Rebase continued successfully"))
			return nil
		},
	}

	// Add a flag to skip editing
	cmd.Flags().BoolP("no-edit", "n", false, "Continue rebase without editing commit message")

	return cmd
}

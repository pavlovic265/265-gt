package commands

import (
	"fmt"

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
		Short: "short for rebase --continue",
		RunE: func(cmd *cobra.Command, args []string) error {
			exeArgs := []string{"-c", "core.editor=vim -f", "rebase", "--continue"}
			err := svc.exe.WithGit().WithArgs(exeArgs).WithGitEditor("vim -f").Run()
			if err != nil {
				return err
			}

			// Run stty sane to fix any terminal issues that might have occurred
			// This is especially useful when Git opens an editor (like vim) during rebase
			// that can mess up terminal display settings
			// Side effects: Resets any custom terminal settings to standard defaults
			// _ = exec.Command("stty", "sane").Run() // Ignore stty errors as they're not critical

			fmt.Println("Rebase continued successfully")
			return nil
		},
	}
}

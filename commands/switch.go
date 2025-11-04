package commands

import (
	"fmt"
	"os"

	"github.com/pavlovic265/265-gt/executor"
	helpers "github.com/pavlovic265/265-gt/git_helpers"
	"github.com/pavlovic265/265-gt/utils/log"
	"github.com/spf13/cobra"
)

type switchCommand struct {
	exe       executor.Executor
	gitHelper helpers.GitHelper
}

func NewSwitchCommand(
	exe executor.Executor,
	gitHelper helpers.GitHelper,
) switchCommand {
	return switchCommand{
		exe:       exe,
		gitHelper: gitHelper,
	}
}

func (svc switchCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:     "switch",
		Aliases: []string{"sw"},
		Short:   "switch back to previous branch",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if err := svc.gitHelper.EnsureGitRepository(); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			exeArgs := []string{"checkout", "-"}
			err := svc.exe.WithGit().WithArgs(exeArgs).Run()
			if err != nil {
				return log.Error("Failed to switch to previous branch", err)
			}

			return nil
		},
	}
}

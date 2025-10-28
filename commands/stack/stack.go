package stack

import (
	"fmt"
	"os"

	"github.com/pavlovic265/265-gt/executor"
	helpers "github.com/pavlovic265/265-gt/git_helpers"
	"github.com/spf13/cobra"
)

type stackCommand struct {
	exe       executor.Executor
	gitHelper helpers.GitHelper
}

func NewStackCommand(
	exe executor.Executor,
	gitHelper helpers.GitHelper,
) stackCommand {
	return stackCommand{
		exe:       exe,
		gitHelper: gitHelper,
	}
}

func (svc stackCommand) Command() *cobra.Command {
	stackCmd := &cobra.Command{
		Use:     "stack",
		Short:   "commands for pull request",
		Aliases: []string{"s"},
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if err := svc.gitHelper.EnsureGitRepository(); err != nil {
				fmt.Printf("%v\n", err)
				os.Exit(1)
			}
		},
	}

	stackCmd.AddCommand(NewRestackCommand(svc.exe, svc.gitHelper).Command())

	return stackCmd
}

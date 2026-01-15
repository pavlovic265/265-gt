package stack

import (
	"fmt"
	"os"

	helpers "github.com/pavlovic265/265-gt/helpers"
	"github.com/pavlovic265/265-gt/runner"
	"github.com/spf13/cobra"
)

type stackCommand struct {
	runner    runner.Runner
	gitHelper helpers.GitHelper
}

func NewStackCommand(
	runner runner.Runner,
	gitHelper helpers.GitHelper,
) stackCommand {
	return stackCommand{
		runner:    runner,
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

	stackCmd.AddCommand(NewRestackCommand(svc.runner, svc.gitHelper).Command())

	return stackCmd
}

package commands

import (
	"fmt"
	"os"

	helpers "github.com/pavlovic265/265-gt/helpers"
	"github.com/pavlovic265/265-gt/runner"
	"github.com/pavlovic265/265-gt/utils/log"
	"github.com/spf13/cobra"
)

type switchCommand struct {
	runner    runner.Runner
	gitHelper helpers.GitHelper
}

func NewSwitchCommand(
	runner runner.Runner,
	gitHelper helpers.GitHelper,
) switchCommand {
	return switchCommand{
		runner:    runner,
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
			if err := svc.runner.Git("checkout", "-"); err != nil {
				return log.Error("Failed to switch to previous branch", err)
			}
			return nil
		},
	}
}

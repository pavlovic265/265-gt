package commands

import (
	"fmt"
	"os"

	helpers "github.com/pavlovic265/265-gt/git_helpers"
	"github.com/pavlovic265/265-gt/runner"
	"github.com/pavlovic265/265-gt/utils/log"
	"github.com/spf13/cobra"
)

type addCommand struct {
	runner    runner.Runner
	gitHelper helpers.GitHelper
}

func NewAddCommand(
	runner runner.Runner,
	gitHelper helpers.GitHelper,
) addCommand {
	return addCommand{
		runner:    runner,
		gitHelper: gitHelper,
	}
}

func (svc addCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:                "add",
		Short:              "git add",
		Aliases:            []string{"a"},
		DisableFlagParsing: true,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if err := svc.gitHelper.EnsureGitRepository(); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := svc.runner.Git(append([]string{"add"}, args...)...); err != nil {
				return log.Error("Failed to stage files", err)
			}

			if len(args) == 0 {
				log.Success("All changes staged")
			} else {
				log.Success("Files staged successfully")
			}
			return nil
		},
	}
}

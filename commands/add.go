package commands

import (
	"fmt"
	"os"

	"github.com/pavlovic265/265-gt/executor"
	helpers "github.com/pavlovic265/265-gt/git_helpers"
	"github.com/pavlovic265/265-gt/utils/log"
	"github.com/spf13/cobra"
)

type addCommand struct {
	exe       executor.Executor
	gitHelper helpers.GitHelper
}

func NewAddCommand(
	exe executor.Executor,
	gitHelper helpers.GitHelper,
) addCommand {
	return addCommand{
		exe:       exe,
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
			exeArgs := append([]string{"add"}, args...)
			err := svc.exe.WithGit().WithArgs(exeArgs).Run()
			if err != nil {
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

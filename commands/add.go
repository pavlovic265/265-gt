package commands

import (
	"github.com/pavlovic265/265-gt/executor"
	"github.com/pavlovic265/265-gt/utils/log"
	"github.com/spf13/cobra"
)

type addCommand struct {
	exe executor.Executor
}

func NewAddCommand(
	exe executor.Executor,
) addCommand {
	return addCommand{
		exe: exe,
	}
}

func (svc addCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:                "add",
		Short:              "git add",
		Aliases:            []string{"a"},
		DisableFlagParsing: true,
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

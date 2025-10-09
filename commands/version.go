package commands

import (
	"fmt"

	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/executor"
	"github.com/spf13/cobra"
)

type versionCommand struct {
	exe executor.Executor
}

func NewVersionCommand(
	exe executor.Executor,
) versionCommand {
	return versionCommand{
		exe: exe,
	}
}

func (svc versionCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "version of current build",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Read version from config
			version := config.Config.GlobalConfig.Version.LastVersion
			if version == "" {
				version = "unknown"
			}

			fmt.Println(version)
			return nil
		},
	}
}

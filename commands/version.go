package commands

import (
	"fmt"

	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/executor"
	"github.com/pavlovic265/265-gt/helpers"
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
	cmd := &cobra.Command{
		Use:   "version",
		Short: "version of current build",
		RunE: func(cmd *cobra.Command, args []string) error {
			latest, _ := cmd.Flags().GetBool("latest")

			if latest {
				return svc.getLatestVersion()
			}

			return svc.getCurrentVersion()
		},
	}

	cmd.Flags().BoolP("latest", "l", false, "get latest version from repository")

	return cmd
}

func (svc versionCommand) getCurrentVersion() error {
	// Read version from config
	version := config.GlobalConfig.Version.LastVersion
	if version == "" {
		version = "unknown"
	}

	fmt.Printf("%s %s\n",
		config.GetInfoStyle().Render("Current version:"),
		config.GetCommandStyle().Render(version))
	return nil
}

func (svc versionCommand) getLatestVersion() error {
	version, err := helpers.GetLatestVersion()
	if err != nil {
		return err
	}

	fmt.Printf("%s %s\n",
		config.GetInfoStyle().Render("Latest version:"),
		config.GetCommandStyle().Render(version))
	return nil
}

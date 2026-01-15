package commands

import (
	"fmt"

	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/runner"
	"github.com/pavlovic265/265-gt/version"
	"github.com/spf13/cobra"
)

type versionCommand struct {
	runner        runner.Runner
	configManager config.ConfigManager
}

func NewVersionCommand(
	runner runner.Runner,
	configManager config.ConfigManager,
) versionCommand {
	return versionCommand{
		runner:        runner,
		configManager: configManager,
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

			return svc.getCurrentVersion(cmd)
		},
	}

	cmd.Flags().BoolP("latest", "l", false, "get latest version from repository")

	return cmd
}

func (svc versionCommand) getCurrentVersion(cmd *cobra.Command) error {
	// Read version from context config
	version := "unknown"
	if cfg, ok := config.GetConfig(cmd.Context()); ok && cfg.Global.Version != nil {
		if cfg.Global.Version.CurrentVersion != "" {
			version = cfg.Global.Version.CurrentVersion
		}
	}

	fmt.Printf("%s %s\n",
		"Current version:",
		version)
	return nil
}

func (svc versionCommand) getLatestVersion() error {
	v, err := version.GetLatestVersion()
	if err != nil {
		return err
	}

	fmt.Printf("%s %s\n",
		"Latest version:",
		v)
	return nil
}

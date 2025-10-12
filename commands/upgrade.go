package commands

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/executor"
	"github.com/pavlovic265/265-gt/utils/pointer"
	"github.com/spf13/cobra"
)

type UpgradeCommand struct {
	exe executor.Executor
}

func NewUpgradeCommand(
	exe executor.Executor,
) UpgradeCommand {
	return UpgradeCommand{
		exe: exe,
	}
}

func (svc UpgradeCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:   "upgrade",
		Short: "upgrade of current build",
		RunE: func(cmd *cobra.Command, args []string) error {
			version, isLatest, err := svc.isLatestVersion()
			if err != nil {
				return err
			}
			if isLatest {
				fmt.Printf("You are already on the latest version: %s\n", pointer.Deref(version))
				return nil // silently fail if version is not found
			}

			installURL := "https://raw.githubusercontent.com/pavlovic265/265-gt/main/scripts/install.sh"
			exeArgs := []string{"-c", fmt.Sprintf("curl -fsSL %s | bash", installURL)}
			err = svc.exe.WithName("bash").WithArgs(exeArgs).Run()
			if err != nil {
				return err
			}

			if err := config.UpdateVersion(pointer.Deref(version)); err != nil {
				fmt.Printf("Warning: Failed to update version in config: %v\n", err)
				return err
			}

			fmt.Printf("✓ %s\n",
				"Tool upgraded successfully")
			return nil
		},
	}
}

func (svc UpgradeCommand) isLatestVersion() (*string, bool, error) {
	// Get latest version from GitHub API
	url := "https://api.github.com/repos/pavlovic265/265-gt/releases/latest"
	resp, err := http.Get(url)
	if err != nil {
		return nil, false, err
	}
	defer func() { _ = resp.Body.Close() }()

	var result struct {
		TagName string `json:"tag_name"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, false, err
	}

	currentVersion := config.GlobalConfig.Version.LastVersion
	if currentVersion == result.TagName {
		return pointer.From(result.TagName), true, nil
	}

	return pointer.From(result.TagName), false, nil
}

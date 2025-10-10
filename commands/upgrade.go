package commands

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/executor"
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
			installURL := "https://raw.githubusercontent.com/pavlovic265/265-gt/main/scripts/install.sh"
			exeArgs := []string{"-c", fmt.Sprintf("curl -fsSL %s | bash", installURL)}
			err := svc.exe.WithName("bash").WithArgs(exeArgs).Run()
			if err != nil {
				return err
			}

			// Update version in config after successful upgrade
			if err := svc.updateVersionInConfig(); err != nil {
				fmt.Printf("Warning: Failed to update version in config: %v\n", err)
			}

			fmt.Printf("%s %s\n",
				config.SuccessIconOnly(),
				config.GetSuccessStyle().Render("Tool upgraded successfully"))
			return nil
		},
	}
}

func (svc UpgradeCommand) updateVersionInConfig() error {
	// Get latest version from GitHub API
	url := "https://api.github.com/repos/pavlovic265/265-gt/releases/latest"
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer func() { _ = resp.Body.Close() }()

	var result struct {
		TagName string `json:"tag_name"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return err
	}

	// Save the updated config
	if err := config.UpdateVersion(result.TagName); err != nil {
		return err
	}

	return nil
}

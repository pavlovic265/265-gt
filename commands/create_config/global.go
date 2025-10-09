package createconfig

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/executor"
	pointer "github.com/pavlovic265/265-gt/utils/pointer"
	timeutils "github.com/pavlovic265/265-gt/utils/timeutils"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

type globalCommand struct {
	exe executor.Executor
}

func (gc globalCommand) getLatestVersion() (string, error) {
	output, err := gc.exe.WithName("gt").WithArgs([]string{"version", "-l"}).RunWithOutput()
	if err != nil {
		return "", err
	}

	// Parse the output to extract just the version
	lines := strings.Split(strings.TrimSpace(output.String()), "\n")
	for _, line := range lines {
		if strings.Contains(line, "Latest version:") {
			parts := strings.Split(line, ":")
			if len(parts) >= 2 {
				return strings.TrimSpace(parts[1]), nil
			}
		}
	}

	return "", fmt.Errorf("could not parse version from output")
}

func NewGlobalCommand(
	exe executor.Executor,
) globalCommand {
	return globalCommand{
		exe: exe,
	}
}

func (svc globalCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:                "global",
		Aliases:            []string{"gl"},
		Short:              "generate global config",
		DisableFlagParsing: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			homeDir, err := os.UserHomeDir()
			if err != nil {
				return err
			}
			filePath := filepath.Join(homeDir, config.FileName)

			_, err = os.Stat(filePath)
			if errors.Is(err, os.ErrNotExist) {
				file, err := os.Create(filePath)
				if err != nil {
					return err
				}
				defer func() { _ = file.Close() }()

				platform, err := HandleSelectPlatform()
				if err != nil {
					return err
				}

				accounts, err := HandleAddAccunts()
				if err != nil {
					return err
				}

				theme, err := HandleSelectTheme()
				if err != nil {
					return err
				}

				// Fetch latest version from repository
				latestVersion, err := svc.getLatestVersion()
				if err != nil {
					// If we can't fetch the latest version, use empty string
					latestVersion = ""
				}

				globalConfig := config.GlobalConfigStruct{
					Version: config.Version{
						LastChecked: timeutils.Now().Format(timeutils.LayoutISOWithTime),
						LastVersion: latestVersion,
					},
					Theme: config.Theme{
						Type: pointer.Deref(theme),
					},
				}
				if pointer.Deref(platform) == config.GitHubPlatform.String() {
					globalConfig.GitHub = config.GitHub{
						Accounts: accounts,
					}
				}

				encoder := yaml.NewEncoder(file)
				encoder.SetIndent(2)
				err = encoder.Encode(&globalConfig)
				if err != nil {
					return err
				}
				defer func() { _ = encoder.Close() }()

				return nil
			} else if err != nil {
				return fmt.Errorf("error checking file: %w", err)
			} else {
				fmt.Printf("File '%s' exists.\n", config.FileName)
			}

			return nil
		},
	}
}

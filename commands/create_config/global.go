package createconfig

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/executor"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

type globalCommand struct {
	exe executor.Executor
}

func NewGlobalCommand(
	exe executor.Executor,
) globalCommand {
	return globalCommand{
		exe: exe,
	}
}

var fileName = ".gtconfig.yaml"

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
			filePath := filepath.Join(homeDir, fileName)

			_, err = os.Stat(filePath)
			if errors.Is(err, os.ErrNotExist) {
				file, err := os.Create(filePath)
				if err != nil {
					return err
				}
				defer file.Close()

				platform, err := HandleSelectPlatform()
				if err != nil {
					return err
				}

				accounts, err := HandleAddAccunts()
				if err != nil {
					return err
				}

				newConfig := config.Config{}
				if *platform == "GitHub" {
					newConfig.GitHub = config.GitHub{
						Accounts: accounts,
					}
				}
				encoder := yaml.NewEncoder(file)
				encoder.SetIndent(2)
				err = encoder.Encode(&newConfig)
				if err != nil {
					return err
				}
				defer encoder.Close()

				return nil
			} else if err != nil {
				return fmt.Errorf("error checking file: %w", err)
			} else {
				fmt.Printf("File '%s' exists.\n", fileName)
			}

			return nil
		},
	}
}

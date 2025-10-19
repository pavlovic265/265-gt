package createconfig

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/constants"
	"github.com/pavlovic265/265-gt/executor"
	"github.com/pavlovic265/265-gt/utils/log"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

type localCommand struct {
	exe executor.Executor
}

func NewLocalCommand(
	exe executor.Executor,
) localCommand {
	return localCommand{
		exe: exe,
	}
}

func (svc localCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:                "local",
		Aliases:            []string{"lo"},
		Short:              "generate local config",
		DisableFlagParsing: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			filePath := filepath.Join(".", constants.FileName)

			_, err := os.Stat(filePath)
			if errors.Is(err, os.ErrNotExist) {
				file, err := os.Create(filePath)
				if err != nil {
					return log.Error("Failed to create local config file", err)
				}
				defer func() { _ = file.Close() }()

				branches, err := HandleAddProtectedBranch()
				if err != nil {
					return log.Error("Failed to configure protected branches", err)
				}

				// add branch to skip for deletions
				localConfg := config.LocalConfigStruct{
					Protected: branches,
				}

				encoder := yaml.NewEncoder(file)
				encoder.SetIndent(2)
				err = encoder.Encode(&localConfg)
				if err != nil {
					return log.Error("Failed to encode local configuration", err)
				}
				defer func() { _ = encoder.Close() }()

				log.Success("Local configuration created successfully")
				return nil
			} else if err != nil {
				return log.Error("Error checking local config file", err)
			} else {
				log.Warning("Local config file already exists")
			}

			return nil
		},
	}
}

package createconfig

import (
	"fmt"

	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/executor"
	"github.com/pavlovic265/265-gt/utils/log"
	"github.com/spf13/cobra"
)

type localCommand struct {
	exe           executor.Executor
	configManager config.ConfigManager
}

func NewLocalCommand(
	exe executor.Executor,
	configManager config.ConfigManager,
) localCommand {
	return localCommand{
		exe:           exe,
		configManager: configManager,
	}
}

func (svc localCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:                "local",
		Aliases:            []string{"lo"},
		Short:              "generate local config",
		DisableFlagParsing: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			// filePath := filepath.Join(".", constants.FileName)

			branches, err := HandleAddProtectedBranch()
			if err != nil {
				return err
			}

			fmt.Println("branches", branches)
			err = svc.configManager.SaveProtectedBranches(branches)
			if err != nil {
				return log.Error("Failed to save local configuration", err)
			}

			log.Success("Local configuration updated successfully")
			return nil
		},
	}
}

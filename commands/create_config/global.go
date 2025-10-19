package createconfig

import (
	"errors"
	"os"

	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/executor"
	"github.com/pavlovic265/265-gt/helpers"
	"github.com/pavlovic265/265-gt/utils/log"
	timeutils "github.com/pavlovic265/265-gt/utils/timeutils"
	"github.com/spf13/cobra"
)

type globalCommand struct {
	exe           executor.Executor
	configManager config.ConfigManager
}

func NewGlobalCommand(
	exe executor.Executor,
	configManager config.ConfigManager,
) globalCommand {
	return globalCommand{
		exe:           exe,
		configManager: configManager,
	}
}

func (svc globalCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:                "global",
		Aliases:            []string{"gl"},
		Short:              "generate global config",
		DisableFlagParsing: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			_, err := svc.configManager.GetGlobalConfigPath()
			if errors.Is(err, os.ErrNotExist) {
				err = svc.createGlobalConfig()
				if err != nil {
					return log.Error("Failed to create global configuration", err)
				}
				return nil
			}

			if err != nil {
				return log.Error("Error checking global config file", err)
			}
			log.Warning("Global config file already exists")

			return nil
		},
	}
}

func (svc globalCommand) createGlobalConfig() error {
	accounts, err := HandleAddAccunts()
	if err != nil {
		return log.Error("Failed to configure accounts", err)
	}

	theme, err := HandleSelectTheme()
	if err != nil {
		return log.Error("Failed to configure theme", err)
	}

	latestVersion, err := helpers.GetLatestVersion()
	if err != nil {
		return log.Error("Failed to get latest version", err)
	}

	globalConfig := config.GlobalConfigStruct{
		Accounts: accounts,
		Version: &config.Version{
			LastChecked:    timeutils.Now().Format(timeutils.LayoutISOWithTime),
			CurrentVersion: latestVersion,
		},
		ActiveAccount: &config.Account{},
		Theme:         theme,
	}

	err = svc.configManager.SaveGlobalConfig(globalConfig)
	if err != nil {
		return log.Error("Failed to save global configuration", err)
	}

	log.Success("Global configuration created successfully")
	return nil
}

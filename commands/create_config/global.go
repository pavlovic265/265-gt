package createconfig

import (
	"errors"
	"fmt"
	"os"

	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/constants"
	"github.com/pavlovic265/265-gt/executor"
	"github.com/pavlovic265/265-gt/helpers"
	pointer "github.com/pavlovic265/265-gt/utils/pointer"
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
			configPath, err := svc.configManager.GetGlobalConfigPath()
			if errors.Is(err, os.ErrNotExist) {
				err = svc.createGlobalConfig(configPath)
				if err != nil {
					return err
				}
				return nil
			}

			if err != nil {
				return fmt.Errorf("error checking file: %w", err)
			}
			fmt.Printf("File '%s' exists.\n", constants.FileName)

			return nil
		},
	}
}

func (svc globalCommand) createGlobalConfig(configPath string) error {

	accounts, err := HandleAddAccunts()
	if err != nil {
		return err
	}

	theme, err := HandleSelectTheme()
	if err != nil {
		return err
	}

	latestVersion, err := helpers.GetLatestVersion()
	if err != nil {
		return err
	}

	globalConfig := config.GlobalConfigStruct{
		Accounts: accounts,
		Version: config.Version{
			LastChecked:    timeutils.Now().Format(timeutils.LayoutISOWithTime),
			CurrentVersion: latestVersion,
		},
		ActiveAccount: config.Account{},
		Theme:         pointer.Deref(theme),
	}

	err = svc.configManager.SaveGlobalConfig(globalConfig)
	if err != nil {
		return err
	}

	return nil
}

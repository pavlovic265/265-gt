package createconfig

import (
	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/executor"
	helpers "github.com/pavlovic265/265-gt/git_helpers"
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
			_, err := svc.configManager.LoadGlobalConfig()
			if err == nil {
				log.Warning("Global config file already exists")
				return nil
			}

			err = svc.createGlobalConfig()
			if err != nil {
				return log.Error("Failed to create global configuration", err)
			}
			return nil
		},
	}
}

func (svc globalCommand) createGlobalConfig() error {
	theme, err := HandleSelectTheme()
	if err != nil {
		return log.Error("Failed to configure theme", err)
	}

	latestVersion, err := helpers.GetLatestVersion()
	if err != nil {
		// If we can't get version, use empty string
		latestVersion = ""
	}

	globalConfig := config.GlobalConfigStruct{
		Accounts: []config.Account{},
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
	log.Info("Run 'gt account add' to add your first account")
	return nil
}

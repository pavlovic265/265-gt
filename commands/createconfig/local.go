package createconfig

import (
	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/runner"
	"github.com/pavlovic265/265-gt/utils/log"
	"github.com/spf13/cobra"
)

type localCommand struct {
	runner        runner.Runner
	configManager config.ConfigManager
}

func NewLocalCommand(
	runner runner.Runner,
	configManager config.ConfigManager,
) localCommand {
	return localCommand{
		runner:        runner,
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
			cfg, ok := config.GetConfig(cmd.Context())
			if !ok {
				return log.ErrorMsg("config not loaded")
			}

			branches, err := HandleAddProtectedBranch()
			if err != nil {
				return err
			}

			if cfg.Local == nil {
				cfg.Local = &config.LocalConfigStruct{}
			}
			cfg.Local.Protected = append(cfg.Local.Protected, branches...)
			cfg.MarkLocalDirty()

			log.Info("Note: 'main' and 'master' are protected by default")
			log.Success("Local configuration updated successfully")
			return nil
		},
	}
}

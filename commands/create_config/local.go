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
			cfg, ok := config.GetConfig(cmd.Context())
			if !ok {
				return log.ErrorMsg("Config not loaded")
			}

			branches, err := HandleAddProtectedBranch()
			if err != nil {
				return err
			}

			fmt.Println("branches", branches)

			if cfg.Local == nil {
				cfg.Local = &config.LocalConfigStruct{}
			}
			cfg.Local.Protected = append(cfg.Local.Protected, branches...)
			cfg.MarkLocalDirty()

			log.Success("Local configuration updated successfully")
			return nil
		},
	}
}

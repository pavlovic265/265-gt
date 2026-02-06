package account

import (
	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/helpers"
	"github.com/pavlovic265/265-gt/utils/log"
	"github.com/spf13/cobra"
)

type attachCommand struct {
	configManager config.ConfigManager
}

func NewAttachCommand(
	configManager config.ConfigManager,
) attachCommand {
	return attachCommand{
		configManager: configManager,
	}
}

func (atc attachCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:     "attach [directory]",
		Aliases: []string{"at"},
		Short:   "Attach active account to a directory",
		Long:    "Configure Git to use the active account's credentials for a specific directory",
		Args:    cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := config.RequireGlobal(cmd.Context())
			if err != nil {
				return err
			}

			if cfg.Global.ActiveAccount == nil || cfg.Global.ActiveAccount.User == "" {
				return log.ErrorMsg("no active account; run 'gt auth' to set an active account")
			}

			// Determine target directory
			targetDir := "."
			if len(args) > 0 {
				targetDir = args[0]
			}

			// Use shared helper to attach account
			if err := helpers.AttachAccountToDir(cfg.Global.ActiveAccount, targetDir); err != nil {
				return err
			}

			return nil
		},
	}
}

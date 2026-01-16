package auth

import (
	"github.com/pavlovic265/265-gt/client"
	"github.com/pavlovic265/265-gt/components"
	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/utils/log"
	"github.com/spf13/cobra"
)

type loginCommand struct {
	configManager config.ConfigManager
}

func NewLoginCommand(
	configManager config.ConfigManager,
) loginCommand {
	return loginCommand{
		configManager: configManager,
	}
}

func (svc loginCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:                "login",
		Aliases:            []string{"lg"},
		Short:              "login user with token",
		DisableFlagParsing: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := config.RequireGlobal(cmd.Context())
			if err != nil {
				return err
			}

			accounts := cfg.Global.Accounts

			var users []string
			for _, acc := range accounts {
				users = append(users, acc.User)
			}

			selected, err := components.SelectString(users)
			if err != nil {
				return log.Error("failed to display user selection menu", err)
			}
			if selected == "" {
				return log.ErrorMsg("no user selected for authentication")
			}

			var account config.Account
			for _, acc := range accounts {
				if acc.User == selected {
					account = acc
					break
				}
			}

			if err := client.Client[account.Platform].AuthLogin(cmd.Context(), account.User); err != nil {
				return log.Error("authentication failed", err)
			}

			log.Successf("Successfully authenticated with %s", selected)
			return nil
		},
	}
}

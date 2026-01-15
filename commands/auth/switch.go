package auth

import (
	"github.com/pavlovic265/265-gt/components"
	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/utils/log"
	"github.com/pavlovic265/265-gt/utils/pointer"
	"github.com/spf13/cobra"
)

type switchCommand struct {
	configManager config.ConfigManager
}

func NewSwitchCommand(
	configManager config.ConfigManager,
) switchCommand {
	return switchCommand{
		configManager: configManager,
	}
}

func (svc switchCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:     "switch",
		Aliases: []string{"sw"},
		Short:   "switch accounts",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := config.RequireGlobal(cmd.Context())
			if err != nil {
				return err
			}

			if len(args) > 0 {
				if err := svc.switchUser(cfg, args[0]); err != nil {
					return err
				}
			} else {
				if err := svc.selectAndswitchUser(cfg); err != nil {
					return err
				}
			}
			return nil
		},
	}
}

func (svc switchCommand) switchUser(cfg *config.ConfigContext, user string) error {
	accounts := cfg.Global.Accounts
	var account *config.Account
	for _, acc := range accounts {
		if user == acc.User {
			account = pointer.From(acc)
			break
		}
	}
	if account == nil {
		log.Warningf("User '%s' does not exist in config", user)
		return nil
	}

	cfg.Global.ActiveAccount = account
	cfg.MarkDirty()

	log.Successf("Switched to account: %s", account.User)
	return nil
}

func (svc switchCommand) selectAndswitchUser(cfg *config.ConfigContext) error {
	accounts := cfg.Global.Accounts
	var users []string
	for _, acc := range accounts {
		users = append(users, acc.User)
	}

	selected, err := components.SelectString(users)
	if err != nil {
		return log.Error("Failed to display user selection menu", err)
	}
	if selected == "" {
		return log.ErrorMsg("No user selected for switching")
	}

	return svc.switchUser(cfg, selected)
}

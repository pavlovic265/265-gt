package auth

import (
	"github.com/pavlovic265/265-gt/config"
	"github.com/spf13/cobra"
)

type authCommand struct {
	configManager config.ConfigManager
}

func NewAuthCommand(
	configManager config.ConfigManager,
) authCommand {
	return authCommand{
		configManager: configManager,
	}
}

func (svc authCommand) Command() *cobra.Command {
	authCmd := &cobra.Command{
		Use:   "auth",
		Short: "auth user",
	}
	authCmd.AddCommand(NewStatusCommand(svc.configManager).Command())
	authCmd.AddCommand(NewLoginCommand(svc.configManager).Command())
	authCmd.AddCommand(NewLogoutCommand(svc.configManager).Command())
	authCmd.AddCommand(NewSwitchCommand(svc.configManager).Command())

	return authCmd
}

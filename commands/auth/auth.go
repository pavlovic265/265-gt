package auth

import (
	"github.com/pavlovic265/265-gt/client"
	"github.com/pavlovic265/265-gt/config"
	"github.com/spf13/cobra"
)

type authCommand struct {
	configManager config.ConfigManager
	cliClient     client.CliClient
}

func NewAuthCommand(
	configManager config.ConfigManager,
	cliClient client.CliClient,
) authCommand {
	return authCommand{
		configManager: configManager,
		cliClient:     cliClient,
	}
}

func (svc authCommand) Command() *cobra.Command {
	authCmd := &cobra.Command{
		Use:   "auth",
		Short: "auth user",
	}
	authCmd.AddCommand(NewStatusCommand(svc.configManager, svc.cliClient).Command())
	authCmd.AddCommand(NewLoginCommand(svc.configManager, svc.cliClient).Command())
	authCmd.AddCommand(NewLogoutCommand(svc.configManager, svc.cliClient).Command())
	authCmd.AddCommand(NewSwitchCommand(svc.configManager).Command())

	return authCmd
}

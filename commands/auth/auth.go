package auth

import (
	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/executor"
	"github.com/spf13/cobra"
)

type authCommand struct {
	exe           executor.Executor
	configManager config.ConfigManager
}

func NewAuthCommand(
	exe executor.Executor,
	configManager config.ConfigManager,
) authCommand {
	return authCommand{
		exe:           exe,
		configManager: configManager,
	}
}

func (svc authCommand) Command() *cobra.Command {
	authCmd := &cobra.Command{
		Use:   "auth",
		Short: "auth user",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			svc.configManager.InitGlobalConfig()
		},
	}
	authCmd.AddCommand(NewStatusCommand(svc.exe, svc.configManager).Command())
	authCmd.AddCommand(NewLoginCommand(svc.exe, svc.configManager).Command())
	authCmd.AddCommand(NewLogoutCommand(svc.exe, svc.configManager).Command())
	authCmd.AddCommand(NewSwitchCommand(svc.exe, svc.configManager).Command())

	return authCmd
}

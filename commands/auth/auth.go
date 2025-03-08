package auth

import (
	"github.com/pavlovic265/265-gt/executor"
	"github.com/spf13/cobra"
)

type authCommand struct {
	exe executor.Executor
}

func NewAuthCommand(
	exe executor.Executor,
) authCommand {
	return authCommand{
		exe: exe,
	}
}

func (svc authCommand) Command() *cobra.Command {
	authCmd := &cobra.Command{
		Use:   "auth",
		Short: "auth user",
	}
	authCmd.AddCommand(NewStatusCommand(svc.exe).Command())
	authCmd.AddCommand(NewLoginCommand(svc.exe).Command())
	authCmd.AddCommand(NewSwitchCommand(svc.exe).Command())

	return authCmd
}

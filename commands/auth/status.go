package auth

import (
	"fmt"

	"github.com/pavlovic265/265-gt/client"
	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/executor"
	"github.com/spf13/cobra"
)

type statusCommand struct {
	exe executor.Executor
}

func NewStatusCommand(
	exe executor.Executor,
) statusCommand {
	return statusCommand{
		exe: exe,
	}
}

func (svc statusCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:                "status",
		Aliases:            []string{"st"},
		Short:              "see status of current auth user",
		DisableFlagParsing: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("Checking status...")

			err := client.GlobalClient.AuthStatus()
			if err != nil {
				fmt.Println(config.ErrorIndicator("Authentication failed"))
				return err
			}
			
			fmt.Println(config.SuccessIndicator("Authentication successful"))
			return nil
		},
	}
}

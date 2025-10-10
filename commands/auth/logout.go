package auth

import (
	"fmt"

	"github.com/pavlovic265/265-gt/client"
	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/executor"
	"github.com/spf13/cobra"
)

type logoutCommand struct {
	exe executor.Executor
}

func NewLogoutCommand(
	exe executor.Executor,
) logoutCommand {
	return logoutCommand{
		exe: exe,
	}
}

func (svc logoutCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:                "logout",
		Aliases:            []string{"lo"},
		Short:              "logout user with token",
		DisableFlagParsing: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			account := config.GetActiveAccount()
			if account == nil {
				return fmt.Errorf("no active account found")
			}

			fmt.Println("Unauthentication started for", account.User)

			if err := client.Client[account.Platform].AuthLogout(account.User); err != nil {
				return err
			}

			fmt.Println("Successfully unauthenticated with", account.User)

			return nil
		},
	}
}

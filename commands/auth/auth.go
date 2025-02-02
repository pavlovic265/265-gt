package auth

import (
	"github.com/spf13/cobra"
)

var authCmd = &cobra.Command{
	Use:                "auth",
	Short:              "auth user",
	DisableFlagParsing: true,
}

func NewAuth() *cobra.Command {
	authCmd.AddCommand(NewAuthStatusCommand())
	authCmd.AddCommand(NewAuthSwichAccountCommand())

	return authCmd
}

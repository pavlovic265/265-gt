package auth

import (
	"fmt"
	"os"

	"github.com/pavlovic265/265-gt/client"
	"github.com/spf13/cobra"
)

func NewStatus() *cobra.Command {
	return &cobra.Command{
		Use:                "status",
		Aliases:            []string{"st"},
		Short:              "see status of current auth user",
		DisableFlagParsing: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			// git checkout -b
			err := client.GlobalClient.AuthStatus()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error with auth: %v\n", err)
				os.Exit(1)
			}

			return nil
		},
	}
}

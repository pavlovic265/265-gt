package main

import (
	"fmt"
	"os"

	"github.com/pavlovic265/265-gt/client"
	"github.com/pavlovic265/265-gt/commands"
	"github.com/pavlovic265/265-gt/commands/auth"
	pullrequests "github.com/pavlovic265/265-gt/commands/pull_requests"
	"github.com/pavlovic265/265-gt/config"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gt",
	Short: "",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		config.InitConfig()
		client.InitCliClient()
		return nil
	},
}

func main() {
	rootCmd.AddCommand(commands.NewStatusCommand())
	rootCmd.AddCommand(commands.NewCreateCommand())
	rootCmd.AddCommand(commands.NewCheckoutCommand())
	rootCmd.AddCommand(commands.NewSwichCommand())
	rootCmd.AddCommand(commands.NewMoveCommand())
	rootCmd.AddCommand(commands.NewAddCommand())
	rootCmd.AddCommand(commands.NewContCommand())
	rootCmd.AddCommand(commands.NewPushCommand())
	rootCmd.AddCommand(commands.NewPullCommand())

	rootCmd.AddCommand(pullrequests.NewPullRequestCommand())

	rootCmd.AddCommand(auth.NewAuth())

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

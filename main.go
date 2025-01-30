package main

import (
	"fmt"
	"os"

	"github.com/pavlovic265/265-gt/client"
	"github.com/pavlovic265/265-gt/commands"
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
	rootCmd.AddCommand(commands.Status())
	rootCmd.AddCommand(commands.Create())
	rootCmd.AddCommand(commands.Checkout())
	rootCmd.AddCommand(commands.Swich())
	rootCmd.AddCommand(commands.Move())
	rootCmd.AddCommand(commands.Add())
	rootCmd.AddCommand(commands.Cont())
	rootCmd.AddCommand(commands.Push())
	rootCmd.AddCommand(commands.Pull())
	rootCmd.AddCommand(pullrequests.NewPullRequestCommand())

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

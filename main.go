package main

import (
	"fmt"
	"os"

	"github.com/pavlovic265/265-gt/client"
	"github.com/pavlovic265/265-gt/commands"
	"github.com/pavlovic265/265-gt/commands/auth"
	pullrequests "github.com/pavlovic265/265-gt/commands/pull_requests"
	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/executor"

	"github.com/spf13/cobra"
)

var exe = executor.NewExe()

var rootCmd = &cobra.Command{
	Use:   "gt",
	Short: "",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		config.InitConfig()
		client.InitCliClient(exe)
		return nil
	},
}

func main() {
	addCommand := commands.NewAddCommand(exe)
	rootCmd.AddCommand(addCommand.Command())

	statusCommand := commands.NewStatusCommand(exe)
	rootCmd.AddCommand(statusCommand.Command())

	swichCommand := commands.NewSwichCommand(exe)
	rootCmd.AddCommand(swichCommand.Command())

	pushCommand := commands.NewPushCommand(exe)
	rootCmd.AddCommand(pushCommand.Command())

	pullCommand := commands.NewPullCommand(exe)
	rootCmd.AddCommand(pullCommand.Command())

	createCommand := commands.NewCreateCommand(exe)
	rootCmd.AddCommand(createCommand.Command())

	contCommand := commands.NewContCommand(exe)
	rootCmd.AddCommand(contCommand.Command())

	checkoutCommand := commands.NewCheckoutCommand(exe)
	rootCmd.AddCommand(checkoutCommand.Command())

	moveCommand := commands.NewMoveCommand(exe)
	rootCmd.AddCommand(moveCommand.Command())

	pullRequestCommand := pullrequests.NewPullRequestCommand(exe)
	rootCmd.AddCommand(pullRequestCommand.Command())

	rootCmd.AddCommand(auth.NewAuth())

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

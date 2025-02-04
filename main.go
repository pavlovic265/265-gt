package main

import (
	"github.com/pavlovic265/265-gt/client"
	"github.com/pavlovic265/265-gt/commands"
	pullrequests "github.com/pavlovic265/265-gt/commands/pull_requests"
	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/executor"

	"github.com/spf13/cobra"
)

var exe = executor.NewExe()

var rootCmd = &cobra.Command{
	Use: "gt",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		config.InitConfig()
		client.InitCliClient(exe)
	},
}

func main() {
	rootCmd.AddCommand(commands.NewAddCommand(exe).Command())
	rootCmd.AddCommand(commands.NewStatusCommand(exe).Command())
	rootCmd.AddCommand(commands.NewSwichCommand(exe).Command())
	rootCmd.AddCommand(commands.NewPushCommand(exe).Command())
	rootCmd.AddCommand(commands.NewPullCommand(exe).Command())
	rootCmd.AddCommand(commands.NewCreateCommand(exe).Command())
	rootCmd.AddCommand(commands.NewContCommand(exe).Command())
	rootCmd.AddCommand(commands.NewCheckoutCommand(exe).Command())
	rootCmd.AddCommand(commands.NewMoveCommand(exe).Command())
	rootCmd.AddCommand(pullrequests.NewPullRequestCommand(exe).Command())

	rootCmd.Execute()
}

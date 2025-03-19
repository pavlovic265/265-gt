package main

import (
	"github.com/pavlovic265/265-gt/client"
	"github.com/pavlovic265/265-gt/commands"
	"github.com/pavlovic265/265-gt/commands/auth"
	"github.com/pavlovic265/265-gt/commands/commit"
	createconfig "github.com/pavlovic265/265-gt/commands/create_config"
	pullrequests "github.com/pavlovic265/265-gt/commands/pull_requests"
	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/executor"

	"github.com/spf13/cobra"
)

var exe = executor.NewExe()

var rootCmd = &cobra.Command{
	Use: "gt",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		config.InitConfig(exe)
		client.InitCliClient(exe)
	},
}

func main() {
	rootCmd.AddCommand(commands.NewAddCommand(exe).Command())
	rootCmd.AddCommand(commands.NewStatusCommand(exe).Command())
	rootCmd.AddCommand(commands.NewSwitchCommand(exe).Command())
	rootCmd.AddCommand(commands.NewPushCommand(exe).Command())
	rootCmd.AddCommand(commands.NewPullCommand(exe).Command())
	rootCmd.AddCommand(commands.NewCreateCommand(exe).Command())
	rootCmd.AddCommand(commands.NewContCommand(exe).Command())
	rootCmd.AddCommand(commands.NewCheckoutCommand(exe).Command())
	rootCmd.AddCommand(commands.NewMoveCommand(exe).Command())
	rootCmd.AddCommand(commands.NewDeleteCommand(exe).Command())
	rootCmd.AddCommand(commands.NewCleanCommand(exe).Command())
	rootCmd.AddCommand(commands.NewVersionCommand(exe).Command())
	rootCmd.AddCommand(commands.NewUpdateCommand(exe).Command())
	rootCmd.AddCommand(commands.NewDownCommand(exe).Command())
	rootCmd.AddCommand(commands.NewUpCommand(exe).Command())
	rootCmd.AddCommand(commit.NewCommitCommand(exe).Command())
	rootCmd.AddCommand(pullrequests.NewPullRequestCommand(exe).Command())
	rootCmd.AddCommand(auth.NewAuthCommand(exe).Command())
	rootCmd.AddCommand(createconfig.NewConfigCommand(exe).Command())

	rootCmd.Execute()
}

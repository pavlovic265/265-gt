package remote

import (
	"github.com/pavlovic265/265-gt/client"
	"github.com/pavlovic265/265-gt/helpers"
	"github.com/pavlovic265/265-gt/runner"
	"github.com/spf13/cobra"
)

func RegisterCommands(root *cobra.Command, r runner.Runner, gh helpers.GitHelper, cc client.CliClient) {
	root.AddCommand(NewPullCommand(r, gh).Command())
	root.AddCommand(NewPushCommand(r, gh, cc).Command())
	root.AddCommand(NewCloneCommand(r, gh).Command())
}

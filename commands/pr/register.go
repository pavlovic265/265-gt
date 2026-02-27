package pr

import (
	"github.com/pavlovic265/265-gt/client"
	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/helpers"
	"github.com/pavlovic265/265-gt/runner"
	"github.com/spf13/cobra"
)

func RegisterCommands(
	root *cobra.Command,
	r runner.Runner,
	cm config.ConfigManager,
	gh helpers.GitHelper,
	cc client.CliClient,
) {
	root.AddCommand(NewPullRequestCommand(r, cm, gh, cc).Command())
}

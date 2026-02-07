package pr

import (
	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/helpers"
	"github.com/pavlovic265/265-gt/runner"
	"github.com/spf13/cobra"
)

func RegisterCommands(root *cobra.Command, r runner.Runner, cm config.ConfigManager, gh helpers.GitHelper) {
	root.AddCommand(NewPullRequestCommand(r, cm, gh).Command())
}

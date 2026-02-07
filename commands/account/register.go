package account

import (
	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/runner"
	"github.com/spf13/cobra"
)

func RegisterCommands(root *cobra.Command, r runner.Runner, cm config.ConfigManager) {
	root.AddCommand(NewAccountCommand(r, cm).Command())
}

package basic

import (
	"github.com/pavlovic265/265-gt/helpers"
	"github.com/pavlovic265/265-gt/runner"
	"github.com/spf13/cobra"
)

func RegisterCommands(root *cobra.Command, r runner.Runner, gh helpers.GitHelper) {
	root.AddCommand(NewAddCommand(r, gh).Command())
	root.AddCommand(NewStatusCommand(r, gh).Command())
	root.AddCommand(NewUnstageCommand(r, gh).Command())
}

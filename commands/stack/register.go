package stack

import (
	"github.com/pavlovic265/265-gt/helpers"
	"github.com/pavlovic265/265-gt/runner"
	"github.com/spf13/cobra"
)

func RegisterCommands(root *cobra.Command, r runner.Runner, gh helpers.GitHelper) {
	root.AddCommand(NewStackCommand(r, gh).Command())
}

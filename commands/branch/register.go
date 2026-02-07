package branch

import (
	"github.com/pavlovic265/265-gt/helpers"
	"github.com/pavlovic265/265-gt/runner"
	"github.com/spf13/cobra"
)

func RegisterCommands(root *cobra.Command, r runner.Runner, gh helpers.GitHelper) {
	root.AddCommand(NewCreateCommand(r, gh).Command())
	root.AddCommand(NewDeleteCommand(r, gh).Command())
	root.AddCommand(NewMoveCommand(r, gh).Command())
	root.AddCommand(NewTrackCommand(r, gh).Command())
	root.AddCommand(NewUpCommand(r, gh).Command())
	root.AddCommand(NewDownCommand(r, gh).Command())
	root.AddCommand(NewCheckoutCommand(r, gh).Command())
	root.AddCommand(NewSwitchCommand(r, gh).Command())
	root.AddCommand(NewContCommand(r, gh).Command())
	root.AddCommand(NewCleanCommand(r, gh).Command())
}

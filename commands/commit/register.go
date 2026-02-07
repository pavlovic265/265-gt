package commit

import (
	"github.com/pavlovic265/265-gt/helpers"
	"github.com/pavlovic265/265-gt/runner"
	"github.com/spf13/cobra"
)

func RegisterCommands(root *cobra.Command, r runner.Runner, gh helpers.GitHelper) {
	root.AddCommand(NewCommitCommand(r, gh).Command())
}

package auth

import (
	"github.com/pavlovic265/265-gt/config"
	"github.com/spf13/cobra"
)

func RegisterCommands(root *cobra.Command, cm config.ConfigManager) {
	root.AddCommand(NewAuthCommand(cm).Command())
}

package auth

import (
	"github.com/pavlovic265/265-gt/client"
	"github.com/pavlovic265/265-gt/config"
	"github.com/spf13/cobra"
)

func RegisterCommands(root *cobra.Command, cm config.ConfigManager, cc client.CliClient) {
	root.AddCommand(NewAuthCommand(cm, cc).Command())
}

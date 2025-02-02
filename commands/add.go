package commands

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func NewAddCommand() *cobra.Command {
	return &cobra.Command{
		Use:                "add",
		Short:              "git add",
		DisableFlagParsing: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			exeArgs := append([]string{"add"}, args...)
			exeCmd := exec.Command("git", exeArgs...)
			exeCmd.Stdout = os.Stdout
			exeCmd.Stderr = os.Stderr

			if err := exeCmd.Run(); err != nil {
				fmt.Fprintf(os.Stderr, "Error executing git add: %v\n", err)
				os.Exit(1)
			}

			return nil
		},
	}
}

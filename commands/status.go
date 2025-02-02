package commands

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func NewStatusCommand() *cobra.Command {
	return &cobra.Command{
		Use:                "status",
		Aliases:            []string{"st"},
		Short:              "git status",
		DisableFlagParsing: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			exeArgs := append([]string{"status"}, args...)
			exeCmd := exec.Command("git", exeArgs...)
			exeCmd.Stdout = os.Stdout
			exeCmd.Stderr = os.Stderr

			if err := exeCmd.Run(); err != nil {
				fmt.Fprintf(os.Stderr, "Error executing git status: %v\n", err)
				os.Exit(1)
			}

			return nil
		},
	}
}

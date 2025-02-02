package commands

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func NewCreateCommand() *cobra.Command {
	return &cobra.Command{
		Use:                "create",
		Aliases:            []string{"c"},
		Short:              "create branch",
		DisableFlagParsing: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			exeArgs := append([]string{"checkout", "-b"}, args...)
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

package commands

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:                "status",
	Aliases:            []string{"st"},
	Short:              "A simple CLI tool to show Git status",
	DisableFlagParsing: true,
}

func Status() *cobra.Command {
	statusCmd.RunE = func(cmd *cobra.Command, args []string) error {
		exeArgs := append([]string{"status"}, args...)
		exeCmd := exec.Command("git", exeArgs...)
		exeCmd.Stdout = os.Stdout
		exeCmd.Stderr = os.Stderr

		if err := exeCmd.Run(); err != nil {
			fmt.Fprintf(os.Stderr, "Error executing git status: %v\n", err)
			os.Exit(1)
		}

		return nil
	}

	return statusCmd
}

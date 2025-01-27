package commands

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var swichCmd = &cobra.Command{
	Use:                "swich",
	Aliases:            []string{"sw"},
	Short:              "Swich back to previous branch",
	DisableFlagParsing: true,
}

func Swich() *cobra.Command {
	swichCmd.RunE = func(cmd *cobra.Command, args []string) error {
		exeArgs := append([]string{"checkout", "-"}, args...)
		exeCmd := exec.Command("git", exeArgs...)
		exeCmd.Stdout = os.Stdout
		exeCmd.Stderr = os.Stderr

		if err := exeCmd.Run(); err != nil {
			fmt.Fprintf(os.Stderr, "Error executing git checkout -: %v\n", err)
			os.Exit(1)
		}

		return nil
	}

	return swichCmd
}

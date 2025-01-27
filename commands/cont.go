package commands

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var contCmd = &cobra.Command{
	Use:                "cont",
	Short:              "Used only finished resloving conflicnts",
	DisableFlagParsing: true,
}

func Cont() *cobra.Command {
	contCmd.RunE = func(cmd *cobra.Command, args []string) error {
		exeArgs := append([]string{"rebase", "--continue"}, args...)
		exeCmd := exec.Command("git", exeArgs...)
		exeCmd.Stdout = os.Stdout
		exeCmd.Stderr = os.Stderr

		if err := exeCmd.Run(); err != nil {
			fmt.Fprintf(os.Stderr, "Error executing git cont: %v\n", err)
			os.Exit(1)
		}

		return nil
	}

	return contCmd
}

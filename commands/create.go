package branch

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:                "create",
	Aliases:            []string{"c"},
	Short:              "create branch",
	DisableFlagParsing: true,
}

func Create() *cobra.Command {
	// git checkout -b
	createCmd.RunE = func(cmd *cobra.Command, args []string) error {
		exeArgs := append([]string{"checkout", "-b"}, args...)
		exeCmd := exec.Command("git", exeArgs...)
		exeCmd.Stdout = os.Stdout
		exeCmd.Stderr = os.Stderr

		if err := exeCmd.Run(); err != nil {
			fmt.Fprintf(os.Stderr, "Error executing git status: %v\n", err)
			os.Exit(1)
		}

		return nil
	}

	return createCmd
}

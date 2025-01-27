package commands

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var pushCmd = &cobra.Command{
	Use:                "push",
	Aliases:            []string{"pu"},
	Short:              "push branch",
	DisableFlagParsing: true,
}

func Push() *cobra.Command {
	pushCmd.RunE = func(cmd *cobra.Command, args []string) error {
		currentBranch, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
		if err != nil {
			return fmt.Errorf("failed to get current branch: %w", err)
		}
		exeArgs := append([]string{"push", "--force", "origin", string(currentBranch)}, args...)
		exeCmd := exec.Command("git", exeArgs...)
		exeCmd.Stdout = os.Stdout
		exeCmd.Stderr = os.Stderr

		if err := exeCmd.Run(); err != nil {
			fmt.Fprintf(os.Stderr, "Error executing git push: %v\n", err)
			os.Exit(1)
		}

		return nil
	}

	return pushCmd
}

package commands

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func NewPullCommand() *cobra.Command {
	return &cobra.Command{
		Use:                "pull",
		Aliases:            []string{"pl"},
		Short:              "pull branch",
		DisableFlagParsing: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			currentBranch, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
			if err != nil {
				return fmt.Errorf("failed to get current branch: %w", err)
			}

			currentBranchName := string(currentBranch[:len(currentBranch)-1])

			exeArgs := append([]string{"pull", "origin", currentBranchName}, args...)
			exeCmd := exec.Command("git", exeArgs...)
			exeCmd.Stdout = os.Stdout
			exeCmd.Stderr = os.Stderr

			if err := exeCmd.Run(); err != nil {
				fmt.Fprintf(os.Stderr, "Error executing git pull: %v\n", err)
				os.Exit(1)
			}

			return nil
		},
	}
}

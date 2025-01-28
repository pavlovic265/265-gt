package commands

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var moveCmd = &cobra.Command{
	Use:                "move",
	Aliases:            []string{"mo"},
	Short:              "rebase branch onto other branch",
	DisableFlagParsing: true,
}

func Move() *cobra.Command {
	moveCmd.RunE = func(cmd *cobra.Command, args []string) error {
		branch := args[0]
		err := rebaseOnto(branch)
		if err != nil {
			fmt.Printf("Error rebasing branch: %s\n", err)
			os.Exit(1)
		}

		return nil
	}

	return moveCmd
}

func rebaseOnto(branch string) error {
	// Get the current branch name
	currentBranch, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
	if err != nil {
		return fmt.Errorf("failed to get current branch: %w", err)
	}

	// Trim newline from command output
	currentBranchName := string(currentBranch[:len(currentBranch)-1])

	// Checkout the branch that needs to be rebased
	checkoutCmd := exec.Command("git", "checkout", branch)
	checkoutCmd.Stdout = os.Stdout
	checkoutCmd.Stderr = os.Stderr
	if err := checkoutCmd.Run(); err != nil {
		return fmt.Errorf("failed to checkout branch %s: %w", branch, err)
	}

	// Rebase onto the current branch
	rebaseCmd := exec.Command("git", "rebase", "--onto", branch, currentBranchName+"~1", currentBranchName)
	rebaseCmd.Stdout = os.Stdout
	rebaseCmd.Stderr = os.Stderr
	if err := rebaseCmd.Run(); err != nil {
		return fmt.Errorf("failed to rebase branch %s onto %s: %w", currentBranchName, branch, err)
	}

	return nil
}

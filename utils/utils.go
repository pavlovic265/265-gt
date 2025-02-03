package utils

import (
	"fmt"

	"github.com/pavlovic265/265-gt/executor"
)

func GetCurrentBranchName(exe executor.Executor) (*string, error) {
	exeArgs := []string{"rev-parse", "--abbrev-ref", "HEAD"}
	output, err := exe.ExecuteWithOutput("git", exeArgs...)
	if err != nil {
		return nil, fmt.Errorf("failed to get current branch: %w", err)
	}
	currentBranch := string(output[:len(output)-1])

	return &currentBranch, nil
}

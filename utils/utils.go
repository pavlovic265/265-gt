package utils

import (
	"strings"

	"github.com/pavlovic265/265-gt/executor"
)

func GetCurrentBranchName(exe executor.Executor) (*string, error) {
	exeArgs := []string{"rev-parse", "--abbrev-ref", "HEAD"}
	output, err := exe.Execute("git", exeArgs...)
	if err != nil {
		return nil, err
	}
	currentBranch := string(output[:len(output)-1])

	return &currentBranch, nil
}

func GetBranches(exe executor.Executor) ([]string, error) {
	exeArgs := []string{"branch", "--list"}
	output, err := exe.Execute("git", exeArgs...)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(output), "\n")
	var branches []string
	for _, line := range lines {
		branch := strings.TrimSpace(line)
		if branch != "" {
			branches = append(branches, strings.TrimPrefix(branch, "* "))
		}
	}
	return branches, nil
}

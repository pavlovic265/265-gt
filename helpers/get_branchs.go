package helpers

import (
	"strings"

	"github.com/pavlovic265/265-gt/executor"
)

// GetCurrentBranchName gets the current branch name
func (gh *GitHelperImpl) GetCurrentBranchName(exe executor.Executor) (*string, error) {
	exeArgs := []string{"rev-parse", "--abbrev-ref", "HEAD"}
	output, err := exe.WithGit().WithArgs(exeArgs).RunWithOutput()
	if err != nil {
		return nil, err
	}
	strOutput := output.String()
	currentBranch := strOutput[:len(strOutput)-1]

	return &currentBranch, nil
}

// GetBranches gets all local branches
func (gh *GitHelperImpl) GetBranches(exe executor.Executor) ([]string, error) {
	exeArgs := []string{"branch", "--list"}
	output, err := exe.WithGit().WithArgs(exeArgs).RunWithOutput()
	if err != nil {
		return nil, err
	}

	lines := strings.Split(output.String(), "\n")
	var branches []string
	for _, line := range lines {
		branch := strings.TrimSpace(line)
		if branch != "" {
			branches = append(branches, strings.TrimPrefix(branch, "* "))
		}
	}
	return branches, nil
}


package helpers

import (
	"strings"
)

// GetCurrentBranchName gets the current branch name
func (gh *GitHelperImpl) GetCurrentBranchName() (*string, error) {
	exeArgs := []string{"rev-parse", "--abbrev-ref", "HEAD"}
	output, err := gh.exe.WithGit().WithArgs(exeArgs).RunWithOutput()
	if err != nil {
		return nil, err
	}
	strOutput := output.String()
	currentBranch := strOutput[:len(strOutput)-1]

	return &currentBranch, nil
}

// GetBranches gets all local branches
func (gh *GitHelperImpl) GetBranches() ([]string, error) {
	exeArgs := []string{"branch", "--list"}
	output, err := gh.exe.WithGit().WithArgs(exeArgs).RunWithOutput()
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

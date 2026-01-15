package helpers

import (
	"fmt"
	"os"
)

func (gh *GitHelperImpl) IsGitRepository() error {
	_, err := gh.runner.GitOutput("rev-parse", "--git-dir")
	if err != nil {
		return fmt.Errorf("not a git repository (or any of the parent directories): .git")
	}
	return nil
}

func (gh *GitHelperImpl) GetGitRoot() (string, error) {
	output, err := gh.runner.GitOutput("rev-parse", "--show-toplevel")
	if err != nil {
		return "", fmt.Errorf("not a git repository (or any of the parent directories): .git")
	}
	return output, nil
}

func (gh *GitHelperImpl) EnsureGitRepository() error {
	err := gh.IsGitRepository()
	if err != nil {
		currentDir, _ := os.Getwd()
		return fmt.Errorf("❌ No git repository found\n\n"+
			"Current directory: %s\n"+
			"Expected: A directory containing a .git folder\n\n"+
			"To fix this:\n"+
			"  1. Navigate to a git repository: cd /path/to/your/repo\n"+
			"  2. Or initialize a new git repository: git init\n"+
			"  3. Or clone an existing repository: git clone <repository-url>\n\n"+
			"Error: %v", currentDir, err)
	}
	return nil
}

package helpers

import (
	"fmt"
	"os"
	"strings"

	"github.com/pavlovic265/265-gt/executor"
)

// IsGitRepository checks if the current directory is a git repository
func (gh *GitHelperImpl) IsGitRepository(exe executor.Executor) error {
	exeArgs := []string{"rev-parse", "--git-dir"}
	_, err := exe.WithGit().WithArgs(exeArgs).RunWithOutput()
	if err != nil {
		return fmt.Errorf("not a git repository (or any of the parent directories): .git")
	}
	return nil
}

// GetGitRoot gets the root directory of the git repository
func (gh *GitHelperImpl) GetGitRoot(exe executor.Executor) (string, error) {
	exeArgs := []string{"rev-parse", "--show-toplevel"}
	output, err := exe.WithGit().WithArgs(exeArgs).RunWithOutput()
	if err != nil {
		return "", fmt.Errorf("not a git repository (or any of the parent directories): .git")
	}
	return strings.TrimSpace(output.String()), nil
}

// EnsureGitRepository ensures the current directory is a git repository
func (gh *GitHelperImpl) EnsureGitRepository(exe executor.Executor) error {
	err := gh.IsGitRepository(exe)
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


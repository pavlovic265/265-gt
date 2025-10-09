package helpers

import (
	"fmt"
	"os"
	"strings"

	"github.com/pavlovic265/265-gt/executor"
)

func IsGitRepository(exe executor.Executor) error {
	exeArgs := []string{"rev-parse", "--git-dir"}
	_, err := exe.WithGit().WithArgs(exeArgs).RunWithOutput()
	if err != nil {
		return fmt.Errorf("not a git repository (or any of the parent directories): .git")
	}
	return nil
}

func GetGitRoot(exe executor.Executor) (string, error) {
	exeArgs := []string{"rev-parse", "--show-toplevel"}
	output, err := exe.WithGit().WithArgs(exeArgs).RunWithOutput()
	if err != nil {
		return "", fmt.Errorf("not a git repository (or any of the parent directories): .git")
	}
	return strings.TrimSpace(output.String()), nil
}

func EnsureGitRepository(exe executor.Executor) error {
	err := IsGitRepository(exe)
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

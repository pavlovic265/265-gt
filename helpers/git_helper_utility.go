package helpers

import (
	"os"
	"strings"
)

// RelinkParentChildren updates the parent pointers of the deleted branch's children
// to point to the deleted branch's parent
func (gh *GitHelperImpl) RelinkParentChildren(parent string, branchChildren []string) error {
	if parent == "" {
		// branch is not tracked, nothing to relink
		return nil
	}

	// Update parent pointer for each child of the deleted branch
	for _, child := range branchChildren {
		if err := gh.SetParent(parent, child); err != nil {
			return err
		}
	}

	return nil
}

func (gh *GitHelperImpl) IsRebaseInProgress() bool {
	// Check for rebase-merge directory (interactive rebase)
	exeArgs := []string{"rev-parse", "--git-path", "rebase-merge"}
	output, err := gh.exe.WithGit().WithArgs(exeArgs).RunWithOutput()
	if err == nil {
		path := strings.TrimSpace(output.String())
		if _, err := os.Stat(path); err == nil {
			return true
		}
	}

	// Check for rebase-apply directory (non-interactive rebase or git am)
	exeArgs = []string{"rev-parse", "--git-path", "rebase-apply"}
	output, err = gh.exe.WithGit().WithArgs(exeArgs).RunWithOutput()
	if err == nil {
		path := strings.TrimSpace(output.String())
		if _, err := os.Stat(path); err == nil {
			return true
		}
	}

	return false
}

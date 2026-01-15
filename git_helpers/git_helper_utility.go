package helpers

import "os"

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
	path, err := gh.runner.GitOutput("rev-parse", "--git-path", "rebase-merge")
	if err == nil {
		if _, err := os.Stat(path); err == nil {
			return true
		}
	}

	// Check for rebase-apply directory (non-interactive rebase or git am)
	path, err = gh.runner.GitOutput("rev-parse", "--git-path", "rebase-apply")
	if err == nil {
		if _, err := os.Stat(path); err == nil {
			return true
		}
	}

	return false
}

func (gh *GitHelperImpl) GetRemoteURL(remoteName string) (string, error) {
	return gh.runner.GitOutput("remote", "get-url", remoteName)
}

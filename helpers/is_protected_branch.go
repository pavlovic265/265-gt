package helpers

import "github.com/pavlovic265/265-gt/config"

// IsProtectedBranch checks if a branch is protected
func (gh *GitHelperImpl) IsProtectedBranch(branch string) bool {
	for _, pBranch := range config.Config.LocalConfig.Protected {
		if branch == pBranch {
			return true
		}
	}
	return false
}

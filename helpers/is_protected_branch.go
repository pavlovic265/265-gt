package helpers

// IsProtectedBranch checks if a branch is protected
func (gh *GitHelperImpl) IsProtectedBranch(branch string) bool {
	for _, pBranch := range gh.configManager.GetProtectedBranches() {
		if branch == pBranch {
			return true
		}
	}
	return false
}

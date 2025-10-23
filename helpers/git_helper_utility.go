package helpers

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
	// If rev-parse succeeds (no error) ⇒ rebase in progress
	exeArgs := []string{"rev-parse", "-q", "--verify", "REBASE_HEAD"}
	err := gh.exe.WithGit().WithArgs(exeArgs).Run()
	return err == nil
}

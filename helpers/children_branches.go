package helpers

// GetChildren gets the children branches for a given branch by scanning all branches
func (gh *GitHelperImpl) GetChildren(branch string) []string {
	branches, err := gh.GetBranches()
	if err != nil {
		return nil
	}

	var children []string
	for _, b := range branches {
		if gh.GetParent(b) == branch {
			children = append(children, b)
		}
	}
	return children
}

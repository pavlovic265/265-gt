package helpers

func (gh *GitHelperImpl) GetChildren(branch string) []string {
	branches, err := gh.GetBranches()
	if err != nil {
		return nil
	}

	var children []string
	for _, b := range branches {
		parent, err := gh.GetParent(b)
		if err != nil {
			// Skip branches where we can't get the parent
			continue
		}
		if parent == branch {
			children = append(children, b)
		}
	}
	return children
}

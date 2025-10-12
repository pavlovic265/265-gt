package helpers

import (
	"strings"
)

func (gh *GitHelperImpl) RelinkParentChildren(
	parent string,
	parentChildren string,
	branch string,
	branchChildren string,
) error {
	if parent == "" {
		// branch is not tracked
		return nil
	}

	// 1. get branch and children
	splitBranchChildren := gh.UnmarshalChildren(branchChildren)

	// 2. get parent and children
	splitParentChildren := gh.UnmarshalChildren(parentChildren)

	var children []string
	// 3. filter branch from parent children
	for _, child := range splitParentChildren {
		if child != branch {
			children = append(children, child)
		}
	}

	// 4. assing branch children to parent children and assing new parent to branch children
	for _, child := range splitBranchChildren {
		// 4.1 assign new parent to children
		if err := gh.SetParent(parent, child); err != nil {
			return err
		}
		// 4.2 assign child to parent children
		children = append(children, child)
	}

	childrenStr := gh.marshalChildren(children)

	if err := gh.SetChildren(parent, childrenStr); err != nil {
		return err
	}

	return nil
}

func (gh *GitHelperImpl) UnmarshalChildren(children string) []string {
	if len(children) == 0 {
		return nil
	}

	return strings.Split(children, " ")
}

func (gh *GitHelperImpl) marshalChildren(children []string) string {
	return strings.TrimSpace(strings.Join(children, " "))
}

package utils

import (
	"strings"

	"github.com/pavlovic265/265-gt/executor"
)

func RelinkParentChildren(
	exe executor.Executor,
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
	splitBranchChildren := UnmarshalChildren(branchChildren)

	// 2. get parent and children
	splitParentChildren := UnmarshalChildren(parentChildren)

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
		if err := SetParent(exe, parent, child); err != nil {
			return err
		}
		// 4.2 assign child to parent children
		children = append(children, child)
	}

	childrenStr := marshalChildren(children)

	if err := SetChildren(exe, parent, childrenStr); err != nil {
		return err
	}

	return nil
}

func UnmarshalChildren(children string) []string {
	if len(children) == 0 {
		return nil
	}

	return strings.Split(children, " ")
}

func marshalChildren(children []string) string {
	return strings.TrimSpace(strings.Join(children, " "))
}

package utils

import (
	"strings"

	"github.com/pavlovic265/265-gt/executor"
)

/**
 * test3 - parent (test2) children()
 * test2 - parent (test1) children(test3) -> getting parent(test1), children(test3) - delring test2 - delting parent(test1)(children(test2) - adding parent(test1)(children(test3))
 * test1 - parent (main)  children(test2)
 * main
 **/
func RelinkParentChildren(
	exe executor.Executor,
	parent string,
	branch string,
) {
	branchChildren := GetChildren(exe, branch)
	splitBranchChildren := unmarshalChildren(branchChildren)
	if splitBranchChildren == nil {
		return
	}

	parentChildren := GetChildren(exe, parent)
	splitParentChildren := unmarshalChildren(parentChildren)
	if splitParentChildren == nil {
		return
	}

	var children []string
	for _, child := range splitParentChildren {
		if child != branch {
			children = append(children, child)
		}
	}

	children = append(children, splitBranchChildren...)
	childrenStr := marshalChildren(children)

	SetChildren(exe, parent, childrenStr)
}

func unmarshalChildren(children string) []string {
	splitChildren := strings.Split(children, " ")
	if len(splitChildren) == 0 {
		return nil
	}

	return splitChildren
}

func marshalChildren(children []string) string {
	return strings.TrimSpace(strings.Join(children, " "))
}

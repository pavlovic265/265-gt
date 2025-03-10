package utils

import (
	"fmt"
	"strings"

	"github.com/pavlovic265/265-gt/executor"
)

/**
 * test3 - parent (test2) children()
 * test2 - parent (test1) children(test3)
 *       -> getting parent(test1), children(test3)
 *       - delring test2 - delting parent(test1)(children(test2)
 *       - adding parent(test1)(children(test3))
 * test1 - parent (main)  children(test2)
 * main
 **/
func RelinkParentChildren(
	exe executor.Executor,
	parent string,
	branch string,
) error {
	fmt.Println("branch ", branch)
	fmt.Println("parent ", parent)
	fmt.Println(":>>>>>>>>>>>>>>>>>")
	branchChildren := GetChildren(exe, branch)
	splitBranchChildren := unmarshalChildren(branchChildren)
	if splitBranchChildren == nil {
		return nil
	}

	parentChildren := GetChildren(exe, parent)
	splitParentChildren := unmarshalChildren(parentChildren)
	if splitParentChildren == nil {
		return nil
	}

	var children []string
	for _, child := range splitParentChildren {
		if child != branch {
			children = append(children, child)
		}
	}

	for _, child := range splitBranchChildren {
		if err := SetParent(exe, parent, child); err != nil {
			return err
		}
		children = append(children, child)
	}

	childrenStr := marshalChildren(children)

	if childrenStr != "" {
		if err := SetChildren(exe, parent, childrenStr); err != nil {
			return err
		}
	} else {
		if err := DeleteChildren(exe, parent); err != nil {
			return err
		}
	}

	return nil
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

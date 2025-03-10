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
	fmt.Println("branch:", branch)
	fmt.Println("parent:", parent)

	parentChildren := GetChildren(exe, parent)
	fmt.Printf("parent children (%s) len (%d)\n", parentChildren, len(parentChildren))
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

	branchChildren := GetChildren(exe, branch)
	fmt.Printf("branch children (%s) len (%d)\n", branchChildren, len(branchChildren))
	fmt.Println("", len(branchChildren) != 0)
	splitBranchChildren := unmarshalChildren(branchChildren)

	if len(splitBranchChildren) != 0 {
		for _, child := range splitBranchChildren {
			fmt.Printf("assign (%s) to parent (%s)\n", child, parent)
			if err := SetParent(exe, parent, child); err != nil {
				return err
			}
			fmt.Printf("child (%s) len (%d)\n", child, len(child))
			children = append(children, child)
		}
	}

	childrenStr := marshalChildren(children)
	fmt.Println("childrenStr: ", childrenStr, len(childrenStr), childrenStr != "", childrenStr != " ")

	if len(childrenStr) != 0 {
		if err := SetChildren(exe, parent, childrenStr); err != nil {
			return err
		}
	}

	fmt.Println(":>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
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

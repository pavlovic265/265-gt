package helpers

import (
	"strings"

	"github.com/pavlovic265/265-gt/executor"
)

// SetParent sets the parent branch for a given child branch
func (gh *GitHelperImpl) SetParent(exe executor.Executor, parent string, child string) error {
	exeArgs := []string{"config", "branch." + child + ".parent", parent}
	err := exe.WithGit().WithArgs(exeArgs).Run()
	if err != nil {
		return err
	}
	return nil
}

// GetParent gets the parent branch for a given branch
func (gh *GitHelperImpl) GetParent(exe executor.Executor, branch string) string {
	exeArgs := []string{"config", "--get", "branch." + branch + ".parent"}
	output, _ := exe.WithGit().WithArgs(exeArgs).RunWithOutput()
	parent := strings.TrimSuffix(output.String(), "\n")
	return parent
}

// DeleteParent removes the parent configuration for a given branch
func (gh *GitHelperImpl) DeleteParent(exe executor.Executor, branch string) error {
	exeArgs := []string{"config", "--unset", "branch." + branch + ".parent"}
	err := exe.WithGit().WithArgs(exeArgs).Run()
	if err != nil {
		return err
	}
	return nil
}

// DeleteFromParentChildren removes a branch from its parent's children list
func (gh *GitHelperImpl) DeleteFromParentChildren(exe executor.Executor, parent, branch string) error {
	children := gh.GetChildren(exe, parent)
	splitChildren := strings.Split(children, " ")
	if len(splitChildren) == 0 {
		return nil
	}

	var newChildren []string

	for _, child := range splitChildren {
		if child != branch {
			newChildren = append(newChildren, child)
		}
	}
	if len(newChildren) > 0 {
		joinChildren := strings.TrimSpace(strings.Join(newChildren, " "))
		if err := gh.SetChildren(exe, parent, joinChildren); err != nil {
			return err
		}
	} else {
		_ = gh.DeleteChildren(exe, parent)
	}
	return nil
}

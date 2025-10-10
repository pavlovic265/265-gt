package helpers

import (
	"strings"

	"github.com/pavlovic265/265-gt/executor"
)

// SetChildren sets the children branches for a given parent branch
func (gh *GitHelperImpl) SetChildren(exe executor.Executor, parent string, children string) error {
	exeArgs := []string{"config", "branch." + parent + ".children", children}
	err := exe.WithGit().WithArgs(exeArgs).Run()
	if err != nil {
		return err
	}
	return nil
}

// GetChildren gets the children branches for a given branch
func (gh *GitHelperImpl) GetChildren(exe executor.Executor, branch string) string {
	exeArgs := []string{"config", "--get", "branch." + branch + ".children"}
	output, _ := exe.WithGit().WithArgs(exeArgs).RunWithOutput()
	children := strings.TrimSuffix(output.String(), "\n")
	return children
}

// DeleteChildren removes the children configuration for a given branch
func (gh *GitHelperImpl) DeleteChildren(exe executor.Executor, branch string) error {
	exeArgs := []string{"config", "--unset", "branch." + branch + ".children"}
	err := exe.WithGit().WithArgs(exeArgs).Run()
	if err != nil {
		return err
	}
	return nil
}

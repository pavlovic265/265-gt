package helpers

import (
	"strings"
)

// SetChildren sets the children branches for a given parent branch
func (gh *GitHelperImpl) SetChildren(parent string, children string) error {
	exeArgs := []string{"config", "branch." + parent + ".children", children}
	err := gh.exe.WithGit().WithArgs(exeArgs).Run()
	if err != nil {
		return err
	}
	return nil
}

// GetChildren gets the children branches for a given branch
func (gh *GitHelperImpl) GetChildren(branch string) string {
	exeArgs := []string{"config", "--get", "branch." + branch + ".children"}
	output, _ := gh.exe.WithGit().WithArgs(exeArgs).RunWithOutput()
	children := strings.TrimSuffix(output.String(), "\n")
	return children
}

// DeleteChildren removes the children configuration for a given branch
func (gh *GitHelperImpl) DeleteChildren(branch string) error {
	exeArgs := []string{"config", "--unset", "branch." + branch + ".children"}
	err := gh.exe.WithGit().WithArgs(exeArgs).Run()
	if err != nil {
		return err
	}
	return nil
}

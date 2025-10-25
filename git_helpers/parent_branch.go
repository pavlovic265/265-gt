package helpers

import (
	"strings"
)

// SetParent sets the parent branch for a given child branch
func (gh *GitHelperImpl) SetParent(parent string, child string) error {
	exeArgs := []string{"config", "--local", "gt.branch." + child + ".parent", parent}
	err := gh.exe.WithGit().WithArgs(exeArgs).Run()
	if err != nil {
		return err
	}
	return nil
}

// GetParent gets the parent branch for a given branch
func (gh *GitHelperImpl) GetParent(branch string) (string, error) {
	exeArgs := []string{"config", "--local", "--get", "gt.branch." + branch + ".parent"}
	output, err := gh.exe.WithGit().WithArgs(exeArgs).RunWithOutput()
	parent := strings.TrimSpace(output.String())
	return parent, err
}

// DeleteParent removes the parent configuration for a given branch
func (gh *GitHelperImpl) DeleteParent(branch string) error {
	exeArgs := []string{"config", "--local", "--unset", "gt.branch." + branch + ".parent"}
	err := gh.exe.WithGit().WithArgs(exeArgs).Run()
	if err != nil {
		return err
	}
	return nil
}

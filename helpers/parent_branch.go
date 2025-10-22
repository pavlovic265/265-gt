package helpers

import (
	"strings"
)

// SetParent sets the parent branch for a given child branch
func (gh *GitHelperImpl) SetParent(parent string, child string) error {
	exeArgs := []string{"config", "branch." + child + ".parent", parent}
	err := gh.exe.WithGit().WithArgs(exeArgs).Run()
	if err != nil {
		return err
	}
	return nil
}

// GetParent gets the parent branch for a given branch
func (gh *GitHelperImpl) GetParent(branch string) string {
	exeArgs := []string{"config", "--get", "branch." + branch + ".parent"}
	output, _ := gh.exe.WithGit().WithArgs(exeArgs).RunWithOutput()
	parent := strings.TrimSuffix(output.String(), "\n")
	return parent
}

// DeleteParent removes the parent configuration for a given branch
func (gh *GitHelperImpl) DeleteParent(branch string) error {
	exeArgs := []string{"config", "--unset", "branch." + branch + ".parent"}
	err := gh.exe.WithGit().WithArgs(exeArgs).Run()
	if err != nil {
		return err
	}
	return nil
}

package helpers

import (
	"strings"
)

func (gh *GitHelperImpl) SetParent(parent string, child string) error {
	exeArgs := []string{"config", "--local", "gt.branch." + child + ".parent", parent}
	err := gh.exe.WithGit().WithArgs(exeArgs).Run()
	if err != nil {
		return err
	}
	return nil
}

func (gh *GitHelperImpl) GetParent(branch string) (string, error) {
	exeArgs := []string{"config", "--local", "--get", "gt.branch." + branch + ".parent"}
	output, err := gh.exe.WithGit().WithArgs(exeArgs).RunWithOutput()
	parent := strings.TrimSpace(output.String())
	return parent, err
}

func (gh *GitHelperImpl) DeleteParent(branch string) error {
	exeArgs := []string{"config", "--local", "--unset", "gt.branch." + branch + ".parent"}
	err := gh.exe.WithGit().WithArgs(exeArgs).Run()
	if err != nil {
		return err
	}
	return nil
}

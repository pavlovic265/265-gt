package helpers

import (
	"strings"

	"github.com/pavlovic265/265-gt/constants"
)

func (gh *GitHelperImpl) SetPending(branchType constants.Branch, branch string) error {
	exeArgs := []string{"config", "--local", "gt.pending." + branchType.String(), branch}
	err := gh.exe.WithGit().WithArgs(exeArgs).Run()
	if err != nil {
		return err
	}
	return nil
}

func (gh *GitHelperImpl) GetPending(branchType constants.Branch) (string, error) {
	exeArgs := []string{"config", "--local", "--get", "gt.pending." + branchType.String()}
	output, err := gh.exe.WithGit().WithArgs(exeArgs).RunWithOutput()
	branch := strings.TrimSpace(output.String())
	return branch, err
}

func (gh *GitHelperImpl) DeletePending(branchType constants.Branch) error {
	exeArgs := []string{"config", "--local", "--unset", "gt.pending." + branchType.String()}
	err := gh.exe.WithGit().WithArgs(exeArgs).Run()
	if err != nil {
		return err
	}
	return nil
}

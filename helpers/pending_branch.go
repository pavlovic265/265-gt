package helpers

import (
	"github.com/pavlovic265/265-gt/constants"
)

func (gh *GitHelperImpl) SetPending(branchType constants.Branch, branch string) error {
	return gh.runner.Git("config", "--local", "gt.pending."+branchType.String(), branch)
}

func (gh *GitHelperImpl) GetPending(branchType constants.Branch) (string, error) {
	return gh.runner.GitOutput("config", "--local", "--get", "gt.pending."+branchType.String())
}

func (gh *GitHelperImpl) DeletePending(branchType constants.Branch) error {
	return gh.runner.Git("config", "--local", "--unset", "gt.pending."+branchType.String())
}

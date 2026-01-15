package helpers

import (
	"github.com/pavlovic265/265-gt/constants"
)

func (gh *GitHelperImpl) SetPending(branchType constants.Branch, branch string) error {
	return gh.runner.Git("config", "--local", constants.GitConfigPendingPrefix+branchType.String(), branch)
}

func (gh *GitHelperImpl) GetPending(branchType constants.Branch) (string, error) {
	return gh.runner.GitOutput("config", "--local", "--get", constants.GitConfigPendingPrefix+branchType.String())
}

func (gh *GitHelperImpl) DeletePending(branchType constants.Branch) error {
	return gh.runner.Git("config", "--local", "--unset", constants.GitConfigPendingPrefix+branchType.String())
}

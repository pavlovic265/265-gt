package helpers

import "github.com/pavlovic265/265-gt/constants"

func (gh *GitHelperImpl) SetParent(parent string, child string) error {
	key := constants.GitConfigBranchPrefix + child + constants.GitConfigParentSuffix
	return gh.runner.Git("config", "--local", key, parent)
}

func (gh *GitHelperImpl) GetParent(branch string) (string, error) {
	key := constants.GitConfigBranchPrefix + branch + constants.GitConfigParentSuffix
	return gh.runner.GitOutput("config", "--local", "--get", key)
}

func (gh *GitHelperImpl) DeleteParent(branch string) error {
	key := constants.GitConfigBranchPrefix + branch + constants.GitConfigParentSuffix
	return gh.runner.Git("config", "--local", "--unset", key)
}

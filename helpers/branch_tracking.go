package helpers

import (
	"fmt"

	"github.com/pavlovic265/265-gt/constants"
)

// Parent branch tracking

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

// Pending branch tracking

func (gh *GitHelperImpl) SetPending(branchType constants.Branch, branch string) error {
	return gh.runner.Git("config", "--local", constants.GitConfigPendingPrefix+branchType.String(), branch)
}

func (gh *GitHelperImpl) GetPending(branchType constants.Branch) (string, error) {
	return gh.runner.GitOutput("config", "--local", "--get", constants.GitConfigPendingPrefix+branchType.String())
}

func (gh *GitHelperImpl) DeletePending(branchType constants.Branch) error {
	return gh.runner.Git("config", "--local", "--unset", constants.GitConfigPendingPrefix+branchType.String())
}

// Children branches

func (gh *GitHelperImpl) GetChildren(branch string) []string {
	branches, err := gh.GetBranches()
	if err != nil {
		return nil
	}

	var children []string
	for _, b := range branches {
		parent, err := gh.GetParent(b)
		if err != nil {
			continue
		}
		if parent == branch {
			children = append(children, b)
		}
	}
	return children
}

// Branch validation

func (gh *GitHelperImpl) ValidateBranchName(name string) error {
	if name == "" {
		return fmt.Errorf("branch name cannot be empty")
	}

	_, err := gh.runner.GitOutput("check-ref-format", "--branch", name)
	if err != nil {
		return fmt.Errorf("invalid branch name '%s'", name)
	}

	return nil
}

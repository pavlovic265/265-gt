package helpers

import (
	"strings"

	"github.com/pavlovic265/265-gt/constants"
	"github.com/pavlovic265/265-gt/utils/log"
)

func (gh *GitHelperImpl) GetCurrentBranch() (string, error) {
	return gh.runner.GitOutput("rev-parse", "--abbrev-ref", "HEAD")
}

func (gh *GitHelperImpl) GetBranches() ([]string, error) {
	output, err := gh.runner.GitOutput("branch", "--list")
	if err != nil {
		return nil, err
	}

	lines := strings.Split(output, "\n")
	var branches []string
	for _, line := range lines {
		branch := strings.TrimSpace(line)
		if branch != "" {
			branches = append(branches, strings.TrimPrefix(branch, "* "))
		}
	}
	return branches, nil
}

func (gh *GitHelperImpl) RebaseBranch(branch string, parent string) error {
	if err := gh.runner.Git("checkout", branch); err != nil {
		return log.Error("Failed to checkout current branch", err)
	}

	_ = gh.SetPending(constants.ParentBranch, parent)
	_ = gh.SetPending(constants.ChildBranch, branch)

	if err := gh.runner.Git("rebase", parent); err != nil {
		log.Warning("Rebase paused due to conflicts. Resolve them, then run `gt cont` or abort.")
		return log.Error("Rebase paused", err)
	}

	_ = gh.DeletePending(constants.ParentBranch)
	_ = gh.DeletePending(constants.ChildBranch)

	if err := gh.SetParent(parent, branch); err != nil {
		return log.Error("Failed to set parent branch relationship", err)
	}

	log.Successf("Branch '%s' rebased onto '%s' successfully", branch, parent)

	return nil
}

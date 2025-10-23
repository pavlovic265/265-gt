package helpers

import (
	"strings"

	"github.com/pavlovic265/265-gt/constants"
	"github.com/pavlovic265/265-gt/utils/log"
)

// GetCurrentBranchName gets the current branch name
func (gh *GitHelperImpl) GetCurrentBranchName() (*string, error) {
	exeArgs := []string{"rev-parse", "--abbrev-ref", "HEAD"}
	output, err := gh.exe.WithGit().WithArgs(exeArgs).RunWithOutput()
	if err != nil {
		return nil, err
	}
	strOutput := output.String()
	currentBranch := strOutput[:len(strOutput)-1]

	return &currentBranch, nil
}

// GetBranches gets all local branches
func (gh *GitHelperImpl) GetBranches() ([]string, error) {
	exeArgs := []string{"branch", "--list"}
	output, err := gh.exe.WithGit().WithArgs(exeArgs).RunWithOutput()
	if err != nil {
		return nil, err
	}

	lines := strings.Split(output.String(), "\n")
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
	exeArgs := []string{"checkout", branch}
	err := gh.exe.WithGit().WithArgs(exeArgs).Run()
	if err != nil {
		return log.Error("Failed to checkout current branch", err)
	}

	_ = gh.SetPending(constants.ParentBranch, parent)
	_ = gh.SetPending(constants.ChildBranch, branch)

	exeArgs = []string{"rebase", parent}
	err = gh.exe.WithGit().WithArgs(exeArgs).Run()
	if err != nil {
		log.Warning("Rebase paused due to conflicts. Resolve them, then run `gt cont` or abort.")
		return log.Error("Rebase paused", err)
	}

	_ = gh.DeletePending(constants.ParentBranch)
	_ = gh.DeletePending(constants.ChildBranch)

	if err := gh.SetParent(parent, branch); err != nil {
		return log.Error("Failed to set parent branch relationship", err)
	}

	log.Success("Branch '" + branch + "' rebased onto '" + parent + "' successfully")

	return nil
}

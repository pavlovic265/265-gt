package helpers

import "fmt"

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

// Package helpers provides git helper utilities for branch management and operations.
package helpers

import (
	"context"

	"github.com/pavlovic265/265-gt/constants"
	"github.com/pavlovic265/265-gt/runner"
)

// GitHelper defines the interface for git helper operations.
type GitHelper interface {
	SetParent(parent string, child string) error
	GetParent(branch string) (string, error)
	DeleteParent(branch string) error
	GetChildren(branch string) []string
	GetCurrentBranch() (string, error)
	GetBranches() ([]string, error)
	GetRemoteBranches() ([]string, error)
	RebaseBranch(branch string, parent string) error
	SetPending(branchType constants.Branch, branch string) error
	GetPending(branchType constants.Branch) (string, error)
	DeletePending(branchType constants.Branch) error
	IsGitRepository() error
	GetGitRoot() (string, error)
	EnsureGitRepository() error
	IsProtectedBranch(ctx context.Context, branch string) bool
	RelinkParentChildren(parent string, branchChildren []string) error
	IsRebaseInProgress() bool
	GetRemoteURL(remoteName string) (string, error)
	ValidateBranchName(name string) error
}

// GitHelperImpl implements GitHelper interface.
type GitHelperImpl struct {
	runner runner.Runner
}

// NewGitHelper creates a new GitHelper instance.
func NewGitHelper(runner runner.Runner) GitHelper {
	return &GitHelperImpl{runner: runner}
}

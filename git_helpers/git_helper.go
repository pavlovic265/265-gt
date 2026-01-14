package helpers

import (
	"context"

	"github.com/pavlovic265/265-gt/constants"
	"github.com/pavlovic265/265-gt/executor"
)

type GitHelper interface {
	SetParent(parent string, child string) error
	GetParent(branch string) (string, error)
	DeleteParent(branch string) error
	GetChildren(branch string) []string
	GetCurrentBranch() (string, error)
	GetBranches() ([]string, error)
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
	CheckGTVersion(ctx context.Context)
}

type GitHelperImpl struct {
	exe executor.Executor
}

func NewGitHelper(exe executor.Executor) GitHelper {
	return &GitHelperImpl{exe: exe}
}

package helpers

import (
	"github.com/pavlovic265/265-gt/executor"
)

// GitHelper defines the interface for all git-related helper operations
type GitHelper interface {
	// Parent/Children operations
	SetParent(exe executor.Executor, parent string, child string) error
	SetChildren(exe executor.Executor, parent string, children string) error
	GetParent(exe executor.Executor, branch string) string
	GetChildren(exe executor.Executor, branch string) string
	DeleteParent(exe executor.Executor, branch string) error
	DeleteChildren(exe executor.Executor, branch string) error
	DeleteFromParentChildren(exe executor.Executor, parent, branch string) error

	// Branch operations
	GetCurrentBranchName(exe executor.Executor) (*string, error)
	GetBranches(exe executor.Executor) ([]string, error)

	// Repository operations
	IsGitRepository(exe executor.Executor) error
	GetGitRoot(exe executor.Executor) (string, error)
	EnsureGitRepository(exe executor.Executor) error

	// Branch protection
	IsProtectedBranch(branch string) bool

	// Utility functions
	RelinkParentChildren(exe executor.Executor, parent string, parentChildren string, branch string, branchChildren string) error
	UnmarshalChildren(children string) []string
}

type GitHelperImpl struct{}

func NewGitHelper() GitHelper {
	return &GitHelperImpl{}
}

package helpers

import (
	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/executor"
)

// GitHelper defines the interface for all git-related helper operations
type GitHelper interface {
	// Parent/Children operations
	SetParent(parent string, child string) error
	SetChildren(parent string, children string) error
	GetParent(branch string) string
	GetChildren(branch string) string
	DeleteParent(branch string) error
	DeleteChildren(branch string) error
	DeleteFromParentChildren(parent, branch string) error

	// Branch operations
	GetCurrentBranchName() (*string, error)
	GetBranches() ([]string, error)

	// Repository operations
	IsGitRepository() error
	GetGitRoot() (string, error)
	EnsureGitRepository() error

	// Branch protection
	IsProtectedBranch(branch string) bool

	// Utility functions
	RelinkParentChildren(
		parent string, parentChildren string, branch string, branchChildren string,
	) error
	UnmarshalChildren(children string) []string

	// Version check
	CheckGTVersion()
}

type GitHelperImpl struct {
	exe           executor.Executor
	configManager config.ConfigManager
}

func NewGitHelper(exe executor.Executor, configManager config.ConfigManager) GitHelper {
	return &GitHelperImpl{exe: exe, configManager: configManager}
}

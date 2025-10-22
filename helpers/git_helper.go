package helpers

import (
	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/executor"
)

// GitHelper defines the interface for all git-related helper operations
type GitHelper interface {
	// Parent operations
	SetParent(parent string, child string) error
	GetParent(branch string) string
	DeleteParent(branch string) error

	// Children operations
	GetChildren(branch string) []string

	// Branch operations
	GetCurrentBranchName() (*string, error)
	GetBranches() ([]string, error)
	RebaseBranch(branch string, parent string) error

	// Repository operations
	IsGitRepository() error
	GetGitRoot() (string, error)
	EnsureGitRepository() error

	// Branch protection
	IsProtectedBranch(branch string) bool

	// Utility functions
	RelinkParentChildren(parent string, branchChildren []string) error

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

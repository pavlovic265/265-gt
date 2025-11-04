package helpers

import (
	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/constants"
	"github.com/pavlovic265/265-gt/executor"
)

// GitHelper defines the interface for all git-related helper operations
type GitHelper interface {
	// Parent operations
	SetParent(parent string, child string) error
	GetParent(branch string) (string, error)
	DeleteParent(branch string) error

	// Children operations
	GetChildren(branch string) []string

	// Branch operations
	GetCurrentBranch() (string, error)
	GetBranches() ([]string, error)
	GetBranchesWithoutCurrent() ([]string, error)
	RebaseBranch(branch string, parent string) error

	// Pending branch
	SetPending(branchType constants.Branch, branch string) error
	GetPending(branchType constants.Branch) (string, error)
	DeletePending(branchType constants.Branch) error

	// Repository operations
	IsGitRepository() error
	GetGitRoot() (string, error)
	EnsureGitRepository() error

	// Branch protection
	IsProtectedBranch(branch string) bool

	// Utility functions
	RelinkParentChildren(parent string, branchChildren []string) error
	IsRebaseInProgress() bool

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

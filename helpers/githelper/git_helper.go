package githelper

import (
	"context"
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/constants"
	"github.com/pavlovic265/265-gt/runner"
	"github.com/pavlovic265/265-gt/utils/log"
)

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

type GitHelperImpl struct {
	runner runner.Runner
}

func NewGitHelper(runner runner.Runner) GitHelper {
	return &GitHelperImpl{runner: runner}
}

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

func (gh *GitHelperImpl) SetPending(branchType constants.Branch, branch string) error {
	return gh.runner.Git("config", "--local", constants.GitConfigPendingPrefix+branchType.String(), branch)
}

func (gh *GitHelperImpl) GetPending(branchType constants.Branch) (string, error) {
	return gh.runner.GitOutput("config", "--local", "--get", constants.GitConfigPendingPrefix+branchType.String())
}

func (gh *GitHelperImpl) DeletePending(branchType constants.Branch) error {
	return gh.runner.Git("config", "--local", "--unset", constants.GitConfigPendingPrefix+branchType.String())
}

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

func (gh *GitHelperImpl) GetRemoteBranches() ([]string, error) {
	output, err := gh.runner.GitOutput("branch", "-r", "--list")
	if err != nil {
		return nil, err
	}

	lines := strings.Split(output, "\n")
	var branches []string
	for _, line := range lines {
		branch := strings.TrimSpace(line)
		if branch == "" || strings.Contains(branch, "HEAD") {
			continue
		}
		branch = strings.TrimPrefix(branch, "origin/")
		branches = append(branches, branch)
	}
	return branches, nil
}

func (gh *GitHelperImpl) RebaseBranch(branch string, parent string) error {
	if err := gh.runner.Git("checkout", branch); err != nil {
		return fmt.Errorf("failed to checkout branch: %w", err)
	}

	_ = gh.SetPending(constants.ParentBranch, parent)
	_ = gh.SetPending(constants.ChildBranch, branch)

	if err := gh.runner.Git("rebase", parent); err != nil {
		log.Warning("Rebase paused due to conflicts. Resolve them, then run `gt cont` or abort.")
		return fmt.Errorf("rebase paused: %w", err)
	}

	_ = gh.DeletePending(constants.ParentBranch)
	_ = gh.DeletePending(constants.ChildBranch)

	if err := gh.SetParent(parent, branch); err != nil {
		return fmt.Errorf("failed to set parent branch relationship: %w", err)
	}

	log.Successf("Branch '%s' rebased onto '%s' successfully", branch, parent)

	return nil
}

func (gh *GitHelperImpl) RelinkParentChildren(parent string, branchChildren []string) error {
	if parent == "" {
		return nil
	}

	for _, child := range branchChildren {
		if err := gh.SetParent(parent, child); err != nil {
			return err
		}
	}

	return nil
}

func (gh *GitHelperImpl) IsRebaseInProgress() bool {
	path, err := gh.runner.GitOutput("rev-parse", "--git-path", "rebase-merge")
	if err == nil {
		if _, err := os.Stat(path); err == nil {
			return true
		}
	}

	path, err = gh.runner.GitOutput("rev-parse", "--git-path", "rebase-apply")
	if err == nil {
		if _, err := os.Stat(path); err == nil {
			return true
		}
	}

	return false
}

func (gh *GitHelperImpl) GetRemoteURL(remoteName string) (string, error) {
	return gh.runner.GitOutput("remote", "get-url", remoteName)
}

func (gh *GitHelperImpl) IsGitRepository() error {
	_, err := gh.runner.GitOutput("rev-parse", "--git-dir")
	if err != nil {
		return fmt.Errorf("not a git repository (or any of the parent directories): .git")
	}
	return nil
}

func (gh *GitHelperImpl) GetGitRoot() (string, error) {
	output, err := gh.runner.GitOutput("rev-parse", "--show-toplevel")
	if err != nil {
		return "", fmt.Errorf("not a git repository (or any of the parent directories): .git")
	}
	return output, nil
}

func (gh *GitHelperImpl) EnsureGitRepository() error {
	err := gh.IsGitRepository()
	if err != nil {
		currentDir, _ := os.Getwd()
		return fmt.Errorf("❌ No git repository found\n\n"+
			"Current directory: %s\n"+
			"Expected: A directory containing a .git folder\n\n"+
			"To fix this:\n"+
			"  1. Navigate to a git repository: cd /path/to/your/repo\n"+
			"  2. Or initialize a new git repository: git init\n"+
			"  3. Or clone an existing repository: git clone <repository-url>\n\n"+
			"Error: %v", currentDir, err)
	}
	return nil
}

var defaultProtectedBranches = []string{"main", "master"}

func (gh *GitHelperImpl) IsProtectedBranch(ctx context.Context, branch string) bool {
	if slices.Contains(defaultProtectedBranches, branch) {
		return true
	}

	cfg, ok := config.GetConfig(ctx)
	if !ok || cfg.Local == nil {
		return false
	}
	return slices.Contains(cfg.Local.Protected, branch)
}

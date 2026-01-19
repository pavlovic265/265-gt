package helpers

import (
	"context"
	"slices"

	"github.com/pavlovic265/265-gt/config"
)

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

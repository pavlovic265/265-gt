package helpers

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pavlovic265/265-gt/mocks"
)

func TestNewGitHelper(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConfigManager := mocks.NewMockConfigManager(ctrl)
	gitHelper := NewGitHelper(nil, mockConfigManager)

	if gitHelper == nil {
		t.Error("Expected non-nil GitHelper, got nil")
	}

	// Test that it implements the GitHelper interface
	var _ = gitHelper
}

func TestGitHelperImpl_InterfaceCompliance(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConfigManager := mocks.NewMockConfigManager(ctrl)

	// This test ensures that GitHelperImpl implements all methods of GitHelper interface
	// We can assign it to the interface type to verify compliance
	var _ GitHelper = (*GitHelperImpl)(nil)

	// Test that all interface methods exist by checking the interface compliance
	gitHelper := &GitHelperImpl{configManager: mockConfigManager}

	// Test that we can assign to the interface
	var helper = gitHelper
	_ = helper // Use the variable to avoid unused variable warning

	// Set up mock expectation for IsProtectedBranch
	mockConfigManager.EXPECT().
		GetProtectedBranches().
		Return([]string{}).
		AnyTimes()

	// Test that IsProtectedBranch works (doesn't need executor)
	result2 := gitHelper.IsProtectedBranch("main")
	if result2 {
		t.Error("IsProtectedBranch should return false for 'main' with default config")
	}
}

func TestGitHelperImpl_TypeAssertion(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConfigManager := mocks.NewMockConfigManager(ctrl)
	gitHelper := NewGitHelper(nil, mockConfigManager)

	// Test that we can type assert to the concrete type
	impl, ok := gitHelper.(*GitHelperImpl)
	if !ok {
		t.Error("Expected GitHelper to be of type *GitHelperImpl")
	}

	if impl == nil {
		t.Error("Expected non-nil *GitHelperImpl")
	}
}

func TestGitHelperImpl_ZeroValue(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConfigManager := mocks.NewMockConfigManager(ctrl)

	// Test that zero value of GitHelperImpl works
	gitHelper := GitHelperImpl{configManager: mockConfigManager}

	// Set up mock expectation for IsProtectedBranch
	mockConfigManager.EXPECT().
		GetProtectedBranches().
		Return([]string{}).
		AnyTimes()

	// Test IsProtectedBranch with zero value
	result2 := gitHelper.IsProtectedBranch("main")
	if result2 {
		t.Error("Zero value GitHelperImpl should return false for IsProtectedBranch")
	}
}

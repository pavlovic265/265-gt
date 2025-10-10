package helpers

import (
	"bytes"
	"errors"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pavlovic265/265-gt/mocks"
)

func TestIsGitRepository(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExecutor := mocks.NewMockExecutor(ctrl)
	gitHelper := &GitHelperImpl{}

	// Set up expectations
	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		WithArgs([]string{"rev-parse", "--git-dir"}).
		Return(mockExecutor).
		Times(1)

	// Create a mock output
	mockOutput := bytes.Buffer{}
	mockOutput.WriteString(".git\n")

	mockExecutor.EXPECT().
		RunWithOutput().
		Return(mockOutput, nil).
		Times(1)

	// Execute the function
	err := gitHelper.IsGitRepository(mockExecutor)

	// Assertions
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestIsGitRepository_NotGitRepo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExecutor := mocks.NewMockExecutor(ctrl)
	gitHelper := &GitHelperImpl{}

	expectedError := errors.New("not a git repository")

	// Set up expectations
	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		WithArgs([]string{"rev-parse", "--git-dir"}).
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		RunWithOutput().
		Return(bytes.Buffer{}, expectedError).
		Times(1)

	// Execute the function
	err := gitHelper.IsGitRepository(mockExecutor)

	// Assertions
	if err == nil {
		t.Error("Expected error, got nil")
	}
	expectedErrMsg := "not a git repository (or any of the parent directories): .git"
	if err.Error() != expectedErrMsg {
		t.Errorf("Expected error '%s', got '%s'", expectedErrMsg, err.Error())
	}
}

func TestGetGitRoot(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExecutor := mocks.NewMockExecutor(ctrl)
	gitHelper := &GitHelperImpl{}

	expectedRoot := "/path/to/git/repo"

	// Set up expectations
	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		WithArgs([]string{"rev-parse", "--show-toplevel"}).
		Return(mockExecutor).
		Times(1)

	// Create a mock output
	mockOutput := bytes.Buffer{}
	mockOutput.WriteString(expectedRoot + "\n")

	mockExecutor.EXPECT().
		RunWithOutput().
		Return(mockOutput, nil).
		Times(1)

	// Execute the function
	result, err := gitHelper.GetGitRoot(mockExecutor)

	// Assertions
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != expectedRoot {
		t.Errorf("Expected '%s', got '%s'", expectedRoot, result)
	}
}

func TestGetGitRoot_WithWhitespace(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExecutor := mocks.NewMockExecutor(ctrl)
	gitHelper := &GitHelperImpl{}

	expectedRoot := "/path/to/git/repo"
	outputWithWhitespace := "  " + expectedRoot + "  \n"

	// Set up expectations
	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		WithArgs([]string{"rev-parse", "--show-toplevel"}).
		Return(mockExecutor).
		Times(1)

	// Create a mock output with whitespace
	mockOutput := bytes.Buffer{}
	mockOutput.WriteString(outputWithWhitespace)

	mockExecutor.EXPECT().
		RunWithOutput().
		Return(mockOutput, nil).
		Times(1)

	// Execute the function
	result, err := gitHelper.GetGitRoot(mockExecutor)

	// Assertions
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != expectedRoot {
		t.Errorf("Expected '%s', got '%s'", expectedRoot, result)
	}
}

func TestGetGitRoot_NotGitRepo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExecutor := mocks.NewMockExecutor(ctrl)
	gitHelper := &GitHelperImpl{}

	expectedError := errors.New("not a git repository")

	// Set up expectations
	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		WithArgs([]string{"rev-parse", "--show-toplevel"}).
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		RunWithOutput().
		Return(bytes.Buffer{}, expectedError).
		Times(1)

	// Execute the function
	result, err := gitHelper.GetGitRoot(mockExecutor)

	// Assertions
	if err == nil {
		t.Error("Expected error, got nil")
	}
	expectedErrMsg := "not a git repository (or any of the parent directories): .git"
	if err.Error() != expectedErrMsg {
		t.Errorf("Expected error '%s', got '%s'", expectedErrMsg, err.Error())
	}
	if result != "" {
		t.Errorf("Expected empty string, got '%s'", result)
	}
}

func TestEnsureGitRepository(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExecutor := mocks.NewMockExecutor(ctrl)
	gitHelper := &GitHelperImpl{}

	// Set up expectations for IsGitRepository
	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		WithArgs([]string{"rev-parse", "--git-dir"}).
		Return(mockExecutor).
		Times(1)

	// Create a mock output
	mockOutput := bytes.Buffer{}
	mockOutput.WriteString(".git\n")

	mockExecutor.EXPECT().
		RunWithOutput().
		Return(mockOutput, nil).
		Times(1)

	// Execute the function
	err := gitHelper.EnsureGitRepository(mockExecutor)

	// Assertions
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestEnsureGitRepository_NotGitRepo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExecutor := mocks.NewMockExecutor(ctrl)
	gitHelper := &GitHelperImpl{}

	expectedError := errors.New("not a git repository")

	// Set up expectations for IsGitRepository
	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		WithArgs([]string{"rev-parse", "--git-dir"}).
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		RunWithOutput().
		Return(bytes.Buffer{}, expectedError).
		Times(1)

	// Execute the function
	err := gitHelper.EnsureGitRepository(mockExecutor)

	// Assertions
	if err == nil {
		t.Error("Expected error, got nil")
	}

	// Check that the error message contains helpful information
	errorMsg := err.Error()
	if !strings.Contains(errorMsg, "❌ No git repository found") {
		t.Errorf("Expected error message to contain '❌ No git repository found', got: %s", errorMsg)
	}
	if !strings.Contains(errorMsg, "Current directory:") {
		t.Errorf("Expected error message to contain 'Current directory:', got: %s", errorMsg)
	}
	if !strings.Contains(errorMsg, "To fix this:") {
		t.Errorf("Expected error message to contain 'To fix this:', got: %s", errorMsg)
	}
}

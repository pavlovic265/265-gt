package helpers

import (
	"testing"

	"github.com/pavlovic265/265-gt/config"
)

func TestIsProtectedBranch(t *testing.T) {
	gitHelper := &GitHelperImpl{}

	// Set up test config
	originalConfig := config.Config
	defer func() {
		config.Config = originalConfig
	}()

	config.Config = config.ConfigStruct{
		LocalConfig: config.LocalConfigStruct{
			Protected: []string{"main", "develop", "master"},
		},
	}

	tests := []struct {
		name     string
		branch   string
		expected bool
	}{
		{
			name:     "Protected branch main",
			branch:   "main",
			expected: true,
		},
		{
			name:     "Protected branch develop",
			branch:   "develop",
			expected: true,
		},
		{
			name:     "Protected branch master",
			branch:   "master",
			expected: true,
		},
		{
			name:     "Non-protected branch feature1",
			branch:   "feature1",
			expected: false,
		},
		{
			name:     "Non-protected branch feature-branch",
			branch:   "feature-branch",
			expected: false,
		},
		{
			name:     "Non-protected branch hotfix",
			branch:   "hotfix",
			expected: false,
		},
		{
			name:     "Empty branch name",
			branch:   "",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := gitHelper.IsProtectedBranch(tt.branch)
			if result != tt.expected {
				t.Errorf("IsProtectedBranch(%s) = %v, expected %v", tt.branch, result, tt.expected)
			}
		})
	}
}

func TestIsProtectedBranch_EmptyProtectedList(t *testing.T) {
	gitHelper := &GitHelperImpl{}

	// Set up test config with empty protected list
	originalConfig := config.Config
	defer func() {
		config.Config = originalConfig
	}()

	config.Config = config.ConfigStruct{
		LocalConfig: config.LocalConfigStruct{
			Protected: []string{},
		},
	}

	tests := []struct {
		name     string
		branch   string
		expected bool
	}{
		{
			name:     "Any branch with empty protected list",
			branch:   "main",
			expected: false,
		},
		{
			name:     "Another branch with empty protected list",
			branch:   "develop",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := gitHelper.IsProtectedBranch(tt.branch)
			if result != tt.expected {
				t.Errorf("IsProtectedBranch(%s) = %v, expected %v", tt.branch, result, tt.expected)
			}
		})
	}
}

func TestIsProtectedBranch_CaseSensitive(t *testing.T) {
	gitHelper := &GitHelperImpl{}

	// Set up test config
	originalConfig := config.Config
	defer func() {
		config.Config = originalConfig
	}()

	config.Config = config.ConfigStruct{
		LocalConfig: config.LocalConfigStruct{
			Protected: []string{"main", "MAIN", "Main"},
		},
	}

	tests := []struct {
		name     string
		branch   string
		expected bool
	}{
		{
			name:     "Exact match lowercase",
			branch:   "main",
			expected: true,
		},
		{
			name:     "Exact match uppercase",
			branch:   "MAIN",
			expected: true,
		},
		{
			name:     "Exact match title case",
			branch:   "Main",
			expected: true,
		},
		{
			name:     "No match with different case",
			branch:   "MainBranch",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := gitHelper.IsProtectedBranch(tt.branch)
			if result != tt.expected {
				t.Errorf("IsProtectedBranch(%s) = %v, expected %v", tt.branch, result, tt.expected)
			}
		})
	}
}

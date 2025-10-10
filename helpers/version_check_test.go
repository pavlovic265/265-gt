package helpers

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/pavlovic265/265-gt/config"
	timeutils "github.com/pavlovic265/265-gt/utils/timeutils"
)

func TestShouldCheckVersion(t *testing.T) {
	// Save original config
	originalConfig := config.Config
	defer func() {
		config.Config = originalConfig
	}()

	tests := []struct {
		name           string
		lastChecked    string
		expectedResult bool
		description    string
	}{
		{
			name:           "No last checked time",
			lastChecked:    "",
			expectedResult: true,
			description:    "Should check when no last checked time is set",
		},
		{
			name:           "Invalid last checked time",
			lastChecked:    "invalid-time",
			expectedResult: true,
			description:    "Should check when last checked time is invalid",
		},
		{
			name:           "Same day as today",
			lastChecked:    timeutils.Now().Format(time.RFC3339),
			expectedResult: false,
			description:    "Should not check when last checked is today",
		},
		{
			name:           "Yesterday",
			lastChecked:    timeutils.Now().AddDate(0, 0, -1).Format(time.RFC3339),
			expectedResult: true,
			description:    "Should check when last checked was yesterday",
		},
		{
			name:           "Tomorrow (edge case)",
			lastChecked:    timeutils.Now().AddDate(0, 0, 1).Format(time.RFC3339),
			expectedResult: true,
			description:    "Should check when last checked is in the future (different day)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set up test config
			config.Config = config.ConfigStruct{
				GlobalConfig: config.GlobalConfigStruct{
					Version: config.Version{
						LastChecked: tt.lastChecked,
					},
				},
			}

			result := shouldCheckVersion()
			if result != tt.expectedResult {
				t.Errorf("shouldCheckVersion() = %v, expected %v. %s", result, tt.expectedResult, tt.description)
			}
		})
	}
}

func TestShowVersionNotification(t *testing.T) {
	// This test mainly ensures the function doesn't panic
	// In a real scenario, you might want to capture stdout to test the output

	tests := []struct {
		name    string
		current string
		latest  string
		url     string
	}{
		{
			name:    "Normal version notification",
			current: "v1.0.0",
			latest:  "v1.1.0",
			url:     "https://github.com/user/repo/releases/tag/v1.1.0",
		},
		{
			name:    "Version without v prefix",
			current: "1.0.0",
			latest:  "1.1.0",
			url:     "https://github.com/user/repo/releases/tag/v1.1.0",
		},
		{
			name:    "Unknown current version",
			current: "unknown",
			latest:  "v1.0.0",
			url:     "https://github.com/user/repo/releases/tag/v1.0.0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// This should not panic
			showVersionNotification(tt.current, tt.latest, tt.url)
		})
	}
}

func TestGetLatestGTVersionWithContext_EnvironmentVariable(t *testing.T) {
	// Save original environment variable
	originalRepo := os.Getenv("GT_REPOSITORY")
	defer func() {
		if originalRepo != "" {
			os.Setenv("GT_REPOSITORY", originalRepo)
		} else {
			os.Unsetenv("GT_REPOSITORY")
		}
	}()

	// Test with missing environment variable
	os.Unsetenv("GT_REPOSITORY")

	ctx := context.Background()
	_, _, err := getLatestGTVersionWithContext(ctx)

	if err == nil {
		t.Error("Expected error when GT_REPOSITORY is not set, got nil")
	}

	expectedErrMsg := "GT_REPOSITORY environment variable not set"
	if err.Error() != expectedErrMsg {
		t.Errorf("Expected error '%s', got '%s'", expectedErrMsg, err.Error())
	}
}

func TestGetLatestGTVersionWithContext_WithRepository(t *testing.T) {
	// Save original environment variable
	originalRepo := os.Getenv("GT_REPOSITORY")
	defer func() {
		if originalRepo != "" {
			os.Setenv("GT_REPOSITORY", originalRepo)
		} else {
			os.Unsetenv("GT_REPOSITORY")
		}
	}()

	// Set a test repository
	os.Setenv("GT_REPOSITORY", "testuser/testrepo")

	ctx := context.Background()
	// This will likely fail due to network request, but we're testing the setup
	_, _, err := getLatestGTVersionWithContext(ctx)

	// We expect an error due to network request, but not the environment variable error
	if err != nil && err.Error() == "GT_REPOSITORY environment variable not set" {
		t.Error("Expected network error, got environment variable error")
	}
}

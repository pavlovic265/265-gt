package sshhelper

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildSSHHost(t *testing.T) {
	tests := []struct {
		name     string
		platform string
		username string
		expected string
	}{
		{
			name:     "GitHub host",
			platform: "github.com",
			username: "testuser",
			expected: "github.com-testuser",
		},
		{
			name:     "GitLab host",
			platform: "gitlab.com",
			username: "myuser",
			expected: "gitlab.com-myuser",
		},
		{
			name:     "Username with special chars",
			platform: "github.com",
			username: "user-name",
			expected: "github.com-user-name",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := BuildSSHHost(tt.platform, tt.username)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestDefaultSSHKeyPath(t *testing.T) {
	tests := []struct {
		name     string
		platform string
		username string
		expected string
	}{
		{
			name:     "GitHub user",
			platform: "github",
			username: "testuser",
			expected: "~/.ssh/github-testuser",
		},
		{
			name:     "GitLab user",
			platform: "gitlab",
			username: "testuser",
			expected: "~/.ssh/gitlab-testuser",
		},
		{
			name:     "Same username different platforms",
			platform: "github",
			username: "myuser",
			expected: "~/.ssh/github-myuser",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := DefaultSSHKeyPath(tt.platform, tt.username)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestExpandPath(t *testing.T) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		t.Skip("Could not get home directory")
	}

	tests := []struct {
		name     string
		path     string
		expected string
	}{
		{
			name:     "Path with tilde",
			path:     "~/.ssh/key",
			expected: filepath.Join(homeDir, ".ssh/key"),
		},
		{
			name:     "Path without tilde",
			path:     "/absolute/path/to/key",
			expected: "/absolute/path/to/key",
		},
		{
			name:     "Relative path",
			path:     "relative/path",
			expected: "relative/path",
		},
		{
			name:     "Just tilde slash",
			path:     "~/",
			expected: homeDir,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := expandPath(tt.path)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestSSHKeyTypeConstants(t *testing.T) {
	assert.Equal(t, SSHKeyType("ed25519"), SSHKeyTypeEd25519)
	assert.Equal(t, SSHKeyType("rsa"), SSHKeyTypeRSA)
}

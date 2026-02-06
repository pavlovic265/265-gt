package helpers

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/pavlovic265/265-gt/runner"
)

type SSHKeyType string

const (
	SSHKeyTypeEd25519 SSHKeyType = "ed25519"
	SSHKeyTypeRSA     SSHKeyType = "rsa"
)

type SSHHelper interface {
	GenerateKey(keyPath, email string, keyType SSHKeyType) error
	AddToAgent(keyPath string) error
	AddHostConfig(host, hostname, identityFile string) error
	GetPublicKey(keyPath string) (string, error)
	HostExists(host string) bool
}

type SSHHelperImpl struct {
	runner runner.Runner
}

func NewSSHHelper(runner runner.Runner) *SSHHelperImpl {
	return &SSHHelperImpl{runner: runner}
}

func (s *SSHHelperImpl) GenerateKey(keyPath, email string, keyType SSHKeyType) error {
	keyPath = expandPath(keyPath)

	dir := filepath.Dir(keyPath)
	if err := os.MkdirAll(dir, 0700); err != nil {
		return fmt.Errorf("failed to create ssh directory: %w", err)
	}

	var args []string
	switch keyType {
	case SSHKeyTypeRSA:
		args = []string{"-t", "rsa", "-b", "4096", "-C", email, "-f", keyPath, "-N", ""}
	default:
		args = []string{"-t", "ed25519", "-C", email, "-f", keyPath, "-N", ""}
	}

	cmd := exec.Command("ssh-keygen", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to generate SSH key: %w", err)
	}

	return nil
}

func (s *SSHHelperImpl) AddToAgent(keyPath string) error {
	keyPath = expandPath(keyPath)

	var args []string
	if runtime.GOOS == "darwin" {
		args = []string{"--apple-use-keychain", keyPath}
	} else {
		args = []string{keyPath}
	}

	cmd := exec.Command("ssh-add", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to add key to ssh-agent: %w", err)
	}

	return nil
}

func (s *SSHHelperImpl) AddHostConfig(host, hostname, identityFile string) error {
	configPath := expandPath("~/.ssh/config")
	identityFile = expandPath(identityFile)

	if err := os.MkdirAll(filepath.Dir(configPath), 0700); err != nil {
		return fmt.Errorf("failed to create ssh directory: %w", err)
	}

	if s.HostExists(host) {
		return nil
	}

	f, err := os.OpenFile(configPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return fmt.Errorf("failed to open ssh config: %w", err)
	}
	defer f.Close()

	entry := fmt.Sprintf(`
Host %s
  HostName %s
  User git
  IdentityFile %s
  AddKeysToAgent yes   # Auto-add key to ssh-agent on first use
  IdentitiesOnly yes   # Only use this key for this host
`, host, hostname, identityFile)

	if runtime.GOOS == "darwin" {
		entry = fmt.Sprintf(`
Host %s
  HostName %s
  User git
  IdentityFile %s
  AddKeysToAgent yes   # Auto-add key to ssh-agent on first use
  UseKeychain yes      # Store passphrase in macOS Keychain
  IdentitiesOnly yes   # Only use this key for this host
`, host, hostname, identityFile)
	}

	if _, err := f.WriteString(entry); err != nil {
		return fmt.Errorf("failed to write ssh config: %w", err)
	}

	return nil
}

func (s *SSHHelperImpl) GetPublicKey(keyPath string) (string, error) {
	keyPath = expandPath(keyPath)
	pubKeyPath := keyPath + ".pub"

	content, err := os.ReadFile(pubKeyPath)
	if err != nil {
		return "", fmt.Errorf("failed to read public key: %w", err)
	}

	return strings.TrimSpace(string(content)), nil
}

func (s *SSHHelperImpl) HostExists(host string) bool {
	configPath := expandPath("~/.ssh/config")

	f, err := os.Open(configPath)
	if err != nil {
		return false
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(strings.ToLower(line), "host ") {
			existingHost := strings.TrimSpace(strings.TrimPrefix(line, "Host "))
			existingHost = strings.TrimSpace(strings.TrimPrefix(existingHost, "host "))
			if existingHost == host {
				return true
			}
		}
	}

	return false
}

func expandPath(path string) string {
	if strings.HasPrefix(path, "~/") {
		home, err := os.UserHomeDir()
		if err != nil {
			return path
		}
		return filepath.Join(home, path[2:])
	}
	return path
}

func BuildSSHHost(platform, username string) string {
	return fmt.Sprintf("%s-%s", platform, username)
}

func DefaultSSHKeyPath(platform, username string) string {
	return fmt.Sprintf("~/.ssh/%s-%s", platform, username)
}

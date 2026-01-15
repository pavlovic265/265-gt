// Package runner provides command execution utilities for git and shell commands.
package runner

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Runner defines the interface for executing commands.
type Runner interface {
	// Git executes a git command with the given arguments.
	Git(args ...string) error
	// GitOutput executes a git command and returns the output.
	GitOutput(args ...string) (string, error)
	// Exec executes an arbitrary command.
	Exec(name string, args ...string) error
	// ExecOutput executes a command and returns the output.
	ExecOutput(name string, args ...string) (string, error)
}

type runnerImpl struct{}

// NewRunner creates a new Runner instance.
func NewRunner() Runner {
	return &runnerImpl{}
}

func (r *runnerImpl) Git(args ...string) error {
	return r.Exec("git", args...)
}

func (r *runnerImpl) GitOutput(args ...string) (string, error) {
	return r.ExecOutput("git", args...)
}

func (r *runnerImpl) Exec(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("%s %s: %w", name, strings.Join(args, " "), err)
	}
	return nil
}

func (r *runnerImpl) ExecOutput(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		errMsg := strings.TrimSpace(stderr.String())
		if errMsg != "" {
			return "", fmt.Errorf("%s %s: %s", name, strings.Join(args, " "), errMsg)
		}
		return "", fmt.Errorf("%s %s: %w", name, strings.Join(args, " "), err)
	}

	return strings.TrimSpace(stdout.String()), nil
}

package executor

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

//go:generate mockgen -source=executor.go -destination=../mocks/mock_executor.go -package=mocks

// Runner executes shell commands
type Runner interface {
	// Git runs a git command with output to terminal
	Git(args ...string) error

	// GitOutput runs a git command and returns stdout
	GitOutput(args ...string) (string, error)

	// Exec runs an arbitrary command with output to terminal
	Exec(name string, args ...string) error

	// ExecOutput runs an arbitrary command and returns stdout
	ExecOutput(name string, args ...string) (string, error)
}

type runner struct{}

func NewRunner() Runner {
	return &runner{}
}

func (r *runner) Git(args ...string) error {
	return r.Exec("git", args...)
}

func (r *runner) GitOutput(args ...string) (string, error) {
	return r.ExecOutput("git", args...)
}

func (r *runner) Exec(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("%s %s: %w", name, strings.Join(args, " "), err)
	}
	return nil
}

func (r *runner) ExecOutput(name string, args ...string) (string, error) {
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

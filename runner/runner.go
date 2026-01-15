package runner

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Runner interface {
	Git(args ...string) error
	GitOutput(args ...string) (string, error)
	Exec(name string, args ...string) error
	ExecOutput(name string, args ...string) (string, error)
}

type runnerImpl struct{}

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

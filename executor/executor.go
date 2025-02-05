package executor

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Executor interface {
	Execute(name string, args ...string) error
	ExecuteWithStdin(name, input string, args ...string) error
	ExecuteWithOutput(name string, args ...string) ([]byte, error)
}

type exe struct{}

func NewExe() Executor {
	return exe{}
}

func (exe exe) Execute(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error executing `%s %s` with err (%v)",
			name,
			strings.Join(args, " "),
			err,
		)
	}

	return nil
}

func (exe exe) ExecuteWithStdin(name, input string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Stdin = strings.NewReader(input)

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error executing `%s %s` with err (%v)",
			name,
			strings.Join(args, " "),
			err,
		)
	}

	return nil
}

func (exe exe) ExecuteWithOutput(name string, args ...string) ([]byte, error) {
	cmd := exec.Command(name, args...)

	output, err := cmd.CombinedOutput()
	if err != nil && string(output) != "" {
		return nil, fmt.Errorf("error executing `%s %s` with err (%v)",
			name,
			strings.Join(args, " "),
			fmt.Errorf("%w", err),
		)
	}

	return output, nil
}

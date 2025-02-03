package executor

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Executor interface {
	Execute(name string, args ...string) error
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

		fmt.Printf("error executing `%s %s` with err (%w)",
			name,
			strings.Join(args, " "),
			err,
		)
		os.Exit(1)
		//		return fmt.Errorf("error executing `%s %s` with err (%w)",
		//			name,
		//			strings.Join(args, " "),
		//			err,
		//		)
	}

	return nil
}

func (exe exe) ExecuteWithOutput(name string, args ...string) ([]byte, error) {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("error executing `%s %s` with err (%w)",
			name,
			strings.Join(args, " "),
			err,
		)
	}

	return output, nil
}

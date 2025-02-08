package executor

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Executor interface {
	Execute(name string, args ...string) ([]byte, error)
}

type exe struct{}

func NewExe() Executor {
	return exe{}
}

func (exe exe) Execute(name string, args ...string) ([]byte, error) {
	var output bytes.Buffer

	cmd := exec.Command(name, args...)
	cmd.Stdout = &output
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("error executing `%s %s`: %v", name, strings.Join(args, " "), err)
	}

	return output.Bytes(), nil
}

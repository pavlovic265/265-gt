package executor

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

type Executor interface {
	Execute(name string, args ...string) (bytes.Buffer, error)
}

type exe struct{}

func NewExe() Executor {
	return exe{}
}

func (exe exe) Execute(name string, args ...string) (bytes.Buffer, error) {
	var output bytes.Buffer

	multiOut := io.MultiWriter(os.Stdout, &output)
	cmd := exec.Command(name, args...)

	cmd.Stdout = multiOut
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return bytes.Buffer{}, fmt.Errorf("error executing `%s %s`: %v", name, strings.Join(args, " "), err)
	}

	return output, nil
}

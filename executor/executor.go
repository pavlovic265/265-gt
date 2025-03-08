package executor

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Executor interface {
	WithGit() Executor
	WithGh() Executor
	WithName(name string) Executor
	WithArgs(args []string) Executor
	WithStdin(stdin string) Executor
	Run() error
	RunWithOutput() (bytes.Buffer, error)
}

type exe struct {
	Name      string
	HasOutput bool
	Args      []string
	Stdin     string
}

func NewExe() Executor {
	return exe{
		Name:      "",
		HasOutput: true,
		Args:      []string{},
		Stdin:     "",
	}
}

func (exe exe) WithName(name string) Executor {
	exe.Name = name
	return exe
}

func (exe exe) WithGit() Executor {
	exe.Name = "git"
	return exe
}

func (exe exe) WithGh() Executor {
	exe.Name = "gh"
	return exe
}

func (exe exe) WithArgs(args []string) Executor {
	exe.Args = args
	return exe
}

func (exe exe) WithStdin(stdin string) Executor {
	exe.Stdin = stdin
	return exe
}

func (exe exe) Run() error {
	exe.HasOutput = false
	_, err := exe.RunWithOutput()
	if err != nil {
		return err
	}

	return nil
}

func (exe exe) RunWithOutput() (bytes.Buffer, error) {
	var output bytes.Buffer
	cmd := exec.Command(exe.Name, exe.Args...)

	if exe.Stdin != "" {
		cmd.Stdin = strings.NewReader(exe.Stdin + "\n")
	}

	if exe.HasOutput {
		cmd.Stdout = os.Stdout
		cmd.Stdout = &output
	} else {
		cmd.Stdout = os.Stdout
	}
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return bytes.Buffer{}, fmt.Errorf("error executing `%s %s`: %v", exe.Name, strings.Join(exe.Args, " "), err)
	}

	return output, nil
}

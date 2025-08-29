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
	RunSilent() error
}

type exe struct {
	Name string

	Args  []string
	Stdin string
}

func NewExe() Executor {
	return exe{
		Name: "",

		Args:  []string{},
		Stdin: "",
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
	cmd := exec.Command(exe.Name, exe.Args...)

	if exe.Stdin != "" {
		cmd.Stdin = strings.NewReader(exe.Stdin + "\n")
	}

	// Set GIT_EDITOR to true to prevent interactive editor from opening
	cmd.Env = append(os.Environ(), "GIT_EDITOR=true")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error executing `%s %s`: %w", exe.Name, strings.Join(exe.Args, " "), err)
	}

	return nil
}

func (exe exe) RunWithOutput() (bytes.Buffer, error) {
	var output bytes.Buffer
	cmd := exec.Command(exe.Name, exe.Args...)

	if exe.Stdin != "" {
		cmd.Stdin = strings.NewReader(exe.Stdin + "\n")
	}

	// Set GIT_EDITOR to true to prevent interactive editor from opening
	cmd.Env = append(os.Environ(), "GIT_EDITOR=true")

	cmd.Stdout = &output
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return bytes.Buffer{}, fmt.Errorf("error executing `%s %s`: %v", exe.Name, strings.Join(exe.Args, " "), err)
	}

	return output, nil
}

func (exe exe) RunSilent() error {
	cmd := exec.Command(exe.Name, exe.Args...)

	if exe.Stdin != "" {
		cmd.Stdin = strings.NewReader(exe.Stdin + "\n")
	}

	// Set GIT_EDITOR to true to prevent interactive editor from opening
	cmd.Env = append(os.Environ(), "GIT_EDITOR=true")

	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error executing `%s %s`: %w", exe.Name, strings.Join(exe.Args, " "), err)
	}

	return nil
}

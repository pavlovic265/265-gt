package utils

import (
	"github.com/pavlovic265/265-gt/executor"
)

func SetParent(exe executor.Executor, parent string, child string) error {
	exeArgs := []string{"config", "gt.branch." + child + ".parent", parent}
	err := exe.WithGit().WithArgs(exeArgs).Run()
	if err != nil {
		return err
	}
	return nil
}

func GetParent(exe executor.Executor, branch string) (*string, error) {
	exeArgs := []string{"config", "--get", "gt.branch." + branch + ".parent"}
	output, err := exe.WithGit().WithArgs(exeArgs).RunWithOutput()
	if err != nil {
		return nil, err
	}
	parent := output.String()
	return &parent, nil
}

func DeleteParent(exe executor.Executor, branch string) error {
	exeArgs := []string{"config", "--unset", "gt.branch." + branch + ".parent"}
	err := exe.WithGit().WithArgs(exeArgs).Run()
	if err != nil {
		return err
	}
	return nil
}

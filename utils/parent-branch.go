package utils

import (
	"strings"

	"github.com/pavlovic265/265-gt/executor"
)

func SetParent(exe executor.Executor, parent string, child string) error {
	exeArgs := []string{"config", "branch." + child + ".parent", parent}
	err := exe.WithGit().WithArgs(exeArgs).Run()
	if err != nil {
		return err
	}
	return nil
}

func GetParent(exe executor.Executor, branch string) string {
	exeArgs := []string{"config", "--get", "branch." + branch + ".parent"}
	output, _ := exe.WithGit().WithArgs(exeArgs).RunWithOutput()
	parent := strings.TrimSuffix(output.String(), "\n")
	return parent
}

func DeleteParent(exe executor.Executor, branch string) error {
	exeArgs := []string{"config", "--unset", "branch." + branch + ".parent"}
	err := exe.WithGit().WithArgs(exeArgs).Run()
	if err != nil {
		return err
	}
	return nil
}

func DeleteFromParentChildren(exe executor.Executor, parent, branch string) error {
	children := GetChildren(exe, parent)
	splitChldren := strings.Split(children, " ")
	if len(splitChldren) == 0 {
		return nil
	}

	var newChildren []string

	for _, child := range splitChldren {
		if child != branch {
			newChildren = append(newChildren, child)
		}
	}
	if len(newChildren) > 0 {
		joinChildren := strings.TrimSpace(strings.Join(newChildren, " "))
		if err := SetChildren(exe, parent, joinChildren); err != nil {
			return err
		}
	} else {
		DeleteChildren(exe, parent)
	}
	return nil
}

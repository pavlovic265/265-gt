package helpers

import (
	"strings"

	"github.com/pavlovic265/265-gt/executor"
)

func SetChildren(exe executor.Executor, parent string, children string) error {
	exeArgs := []string{"config", "branch." + parent + ".children", children}
	err := exe.WithGit().WithArgs(exeArgs).Run()
	if err != nil {
		return err
	}
	return nil
}

func GetChildren(exe executor.Executor, branch string) string {
	exeArgs := []string{"config", "--get", "branch." + branch + ".children"}
	output, _ := exe.WithGit().WithArgs(exeArgs).RunWithOutput()
	parent := strings.TrimSuffix(output.String(), "\n")
	return parent
}

func DeleteChildren(exe executor.Executor, branch string) error {
	exeArgs := []string{"config", "--unset", "branch." + branch + ".children"}
	err := exe.WithGit().WithArgs(exeArgs).Run()
	if err != nil {
		return err
	}
	return nil
}

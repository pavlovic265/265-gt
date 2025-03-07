package components

import (
	"github.com/pavlovic265/265-gt/executor"
)

// git config --global alias.set-parent '!f() { git config branch.$(git rev-parse --abbrev-ref HEAD).parent "$1"; }; f'
func SetParent(exe executor.Executor, parent string, child string) error {
	exeArgs := []string{"config", "gt.branch." + child + ".parent", parent}
	err := exe.WithGit().WithArgs(exeArgs).Run()
	if err != nil {
		return err
	}
	return nil
}

// git config --global alias.get-parent '!git config --get branch.$(git rev-parse --abbrev-ref HEAD).parent 2>/dev/null || echo "No parent set"'
func GetParent(exe executor.Executor, branch string) (*string, error) {
	exeArgs := []string{"config", "--get", "gt.branch." + branch + ".parent"}
	output, err := exe.WithGit().WithArgs(exeArgs).RunWithOutput()
	if err != nil {
		return nil, err
	}
	parent := output.String()
	return &parent, nil
}

// git config --global alias.delete-parent '!git config --unset branch.$(git rev-parse --abbrev-ref HEAD).parent 2>/dev/null || echo "No parent to delete"'
func DeleteParent(exe executor.Executor, branch string) error {
	exeArgs := []string{"config", "--unset", "gt.branch." + branch + ".parent"}
	err := exe.WithGit().WithArgs(exeArgs).Run()
	if err != nil {
		return err
	}
	return nil
}

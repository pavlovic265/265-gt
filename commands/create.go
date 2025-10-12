package commands

import (
	"fmt"
	"strings"

	"github.com/pavlovic265/265-gt/executor"
	"github.com/pavlovic265/265-gt/helpers"
	pointer "github.com/pavlovic265/265-gt/utils/pointer"
	"github.com/spf13/cobra"
)

type createCommand struct {
	exe       executor.Executor
	gitHelper helpers.GitHelper
}

func NewCreateCommand(
	exe executor.Executor,
	gitHelper helpers.GitHelper,
) createCommand {
	return createCommand{
		exe:       exe,
		gitHelper: gitHelper,
	}
}

func (svc createCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:     "create",
		Aliases: []string{"c"},
		Short:   "create branch",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return fmt.Errorf("missing branch name")
			}
			branch := args[0]
			parent, err := svc.gitHelper.GetCurrentBranchName(svc.exe)
			if err != nil {
				return err
			}

			exeArgs := []string{"checkout", "-b", branch}
			err = svc.exe.WithGit().WithArgs(exeArgs).Run()
			if err != nil {
				return err
			}

			if err := svc.gitHelper.SetParent(svc.exe, pointer.Deref(parent), branch); err != nil {
				return err
			}

			if err := svc.setChildrenBranch(pointer.Deref(parent), branch); err != nil {
				return err
			}

			fmt.Println("Branch '" + branch + "' created and switched to successfully")
			return nil
		},
	}
}

func (svc createCommand) setChildrenBranch(parent, child string) error {
	children := svc.gitHelper.GetChildren(svc.exe, parent)

	splitedChildren := strings.Split(children, " ")
	splitedChildren = append(splitedChildren, child)
	joinedChildren := strings.TrimSpace(strings.Join(splitedChildren, " "))

	if err := svc.gitHelper.SetChildren(svc.exe, parent, joinedChildren); err != nil {
		return err
	}

	return nil
}

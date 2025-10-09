package commands

import (
	"fmt"
	"strings"

	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/executor"
	"github.com/pavlovic265/265-gt/utils"
	"github.com/spf13/cobra"
)

type createCommand struct {
	exe executor.Executor
}

func NewCreateCommand(
	exe executor.Executor,
) createCommand {
	return createCommand{
		exe: exe,
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
			parent, err := utils.GetCurrentBranchName(svc.exe)
			if err != nil {
				return err
			}

			exeArgs := []string{"checkout", "-b", branch}
			err = svc.exe.WithGit().WithArgs(exeArgs).Run()
			if err != nil {
				return err
			}

			if err := utils.SetParent(svc.exe, utils.Deref(parent), branch); err != nil {
				return err
			}

			if err := svc.setChildrenBranch(utils.Deref(parent), branch); err != nil {
				return err
			}

			fmt.Println(config.SuccessIndicator("Branch '" + branch + "' created and switched to successfully"))
			return nil
		},
	}
}

func (svc createCommand) setChildrenBranch(parent, child string) error {
	children := utils.GetChildren(svc.exe, parent)

	splitedChildren := strings.Split(children, " ")
	splitedChildren = append(splitedChildren, child)
	joinedChildren := strings.TrimSpace(strings.Join(splitedChildren, " "))

	if err := utils.SetChildren(svc.exe, parent, joinedChildren); err != nil {
		return err
	}

	return nil
}

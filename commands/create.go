package commands

import (
	"github.com/pavlovic265/265-gt/executor"
	"github.com/pavlovic265/265-gt/helpers"
	"github.com/pavlovic265/265-gt/utils/log"
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
				return log.ErrorMsg("Branch name is required")
			}
			branch := args[0]
			parent, err := svc.gitHelper.GetCurrentBranch()
			if err != nil {
				return log.Error("Failed to get current branch name", err)
			}

			exeArgs := []string{"checkout", "-b", branch}
			err = svc.exe.WithGit().WithArgs(exeArgs).Run()
			if err != nil {
				return log.Error("Failed to create and checkout branch", err)
			}

			if err := svc.gitHelper.SetParent(parent, branch); err != nil {
				return log.Error("Failed to set parent branch relationship", err)
			}

			log.Success("Branch '" + branch + "' created and switched to successfully")
			return nil
		},
	}
}

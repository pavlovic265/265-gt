package stack

import (
	"github.com/pavlovic265/265-gt/executor"
	helpers "github.com/pavlovic265/265-gt/git_helpers"
	"github.com/pavlovic265/265-gt/utils/log"
	"github.com/spf13/cobra"
)

type restackCommand struct {
	runner    executor.Runner
	gitHelper helpers.GitHelper
}

func NewRestackCommand(
	runner executor.Runner,
	gitHelper helpers.GitHelper,
) restackCommand {
	return restackCommand{
		runner:    runner,
		gitHelper: gitHelper,
	}
}

func (svc restackCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:     "restack",
		Aliases: []string{"rs"},
		Short:   "Restack branches",
		RunE: func(cmd *cobra.Command, args []string) error {
			branch, err := svc.gitHelper.GetCurrentBranch()
			if err != nil {
				return err
			}
			queue := []string{branch}
			for len(queue) > 0 {
				parent := queue[0]
				queue = queue[1:]

				children := svc.gitHelper.GetChildren(parent)

				for _, child := range children {
					if child == parent {
						continue
					}

					if err := svc.gitHelper.RebaseBranch(child, parent); err != nil {
						return err
					}

					queue = append(queue, child)

				}

			}

			log.Success("Restack completed")
			return nil
		},
	}
}

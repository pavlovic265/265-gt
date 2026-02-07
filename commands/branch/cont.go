package branch

import (
	"github.com/pavlovic265/265-gt/constants"
	helpers "github.com/pavlovic265/265-gt/helpers"
	"github.com/pavlovic265/265-gt/runner"
	"github.com/pavlovic265/265-gt/utils/log"
	"github.com/spf13/cobra"
)

type contCommand struct {
	runner    runner.Runner
	gitHelper helpers.GitHelper
}

func NewContCommand(
	runner runner.Runner,
	gitHelper helpers.GitHelper,
) contCommand {
	return contCommand{
		runner:    runner,
		gitHelper: gitHelper,
	}
}

func (svc contCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:   "cont",
		Short: "short for rebase --continue",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := svc.gitHelper.EnsureGitRepository(); err != nil {
				return err
			}

			if err := svc.runner.Git("rebase", "--continue"); err != nil {
				return log.Error("failed to continue rebase", err)
			}

			if svc.gitHelper.IsRebaseInProgress() {
				return nil
			}

			parent, pErr := svc.gitHelper.GetPending(constants.ParentBranch)
			child, cErr := svc.gitHelper.GetPending(constants.ChildBranch)
			if pErr == nil && cErr == nil {
				if parent != "" && child != "" {
					if err := svc.gitHelper.SetParent(parent, child); err != nil {
						return log.Error("failed to set parent branch relationship", err)
					}

					_ = svc.gitHelper.DeletePending(constants.ParentBranch)
					_ = svc.gitHelper.DeletePending(constants.ChildBranch)

					log.Success("Rebase completed and metadata updated")
				}
			}

			return nil
		},
	}
}

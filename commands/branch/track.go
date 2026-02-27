package branch

import (
	helpers "github.com/pavlovic265/265-gt/helpers"
	"github.com/pavlovic265/265-gt/runner"
	"github.com/pavlovic265/265-gt/ui/components"
	"github.com/pavlovic265/265-gt/utils/log"
	"github.com/spf13/cobra"
)

type trackCommand struct {
	runner    runner.Runner
	gitHelper helpers.GitHelper
}

func NewTrackCommand(
	runner runner.Runner,
	gitHelper helpers.GitHelper,
) trackCommand {
	return trackCommand{
		runner:    runner,
		gitHelper: gitHelper,
	}
}

func (svc trackCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:     "track",
		Aliases: []string{"tr"},
		Short:   "track existing branch",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := svc.gitHelper.EnsureGitRepository(); err != nil {
				return err
			}

			branch, err := svc.gitHelper.GetCurrentBranch()
			if err != nil {
				return log.Error("failed to get current branch name", err)
			}

			branches, err := svc.gitHelper.GetBranches()
			if err != nil {
				return log.Error("failed to get branches", err)
			}

			var branchesWithoutCurrent []string
			for _, b := range branches {
				if b != branch {
					branchesWithoutCurrent = append(branchesWithoutCurrent, b)
				}
			}

			selected, err := components.SelectString(branchesWithoutCurrent)
			if err != nil {
				return log.Error("failed to display branch selection", err)
			}
			if selected == "" {
				return log.ErrorMsg("no branch selected")
			}

			if err := svc.gitHelper.SetParent(selected, branch); err != nil {
				return log.Error("failed to set parent", err)
			}

			log.Successf("Successfully tracking %s", branch)
			return nil
		},
	}
}

package basic

import (
	helpers "github.com/pavlovic265/265-gt/helpers"
	"github.com/pavlovic265/265-gt/runner"
	"github.com/pavlovic265/265-gt/utils/log"
	"github.com/spf13/cobra"
)

type cherryPickCommand struct {
	runner    runner.Runner
	gitHelper helpers.GitHelper
}

func NewCherryPickCommand(
	runner runner.Runner,
	gitHelper helpers.GitHelper,
) cherryPickCommand {
	return cherryPickCommand{
		runner:    runner,
		gitHelper: gitHelper,
	}
}

func (svc cherryPickCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:                "cherry-pick",
		Aliases:            []string{"cp"},
		Short:              "git cherry-pick",
		DisableFlagParsing: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := svc.gitHelper.EnsureGitRepository(); err != nil {
				return err
			}

			if err := svc.runner.Git(append([]string{"cherry-pick"}, args...)...); err != nil {
				return log.Error("failed to cherry-pick commit", err)
			}

			log.Success("Cherry-pick completed successfully")
			return nil
		},
	}
}

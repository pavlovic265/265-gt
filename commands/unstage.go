package commands

import (
	helpers "github.com/pavlovic265/265-gt/helpers"
	"github.com/pavlovic265/265-gt/runner"
	"github.com/pavlovic265/265-gt/utils/log"
	"github.com/spf13/cobra"
)

type unstageCommand struct {
	runner    runner.Runner
	gitHelper helpers.GitHelper
}

func NewUnstageCommand(
	runner runner.Runner,
	gitHelper helpers.GitHelper,
) unstageCommand {
	return unstageCommand{
		runner:    runner,
		gitHelper: gitHelper,
	}
}

func (svc unstageCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:     "unstage",
		Aliases: []string{"us"},
		Short:   "unstage files",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := svc.gitHelper.EnsureGitRepository(); err != nil {
				return err
			}

			gitArgs := append([]string{"restore", "--staged"}, args...)
			if err := svc.runner.Git(gitArgs...); err != nil {
				return log.Error("Failed to unstage files", err)
			}

			if len(args) == 0 {
				log.Success("All staged changes unstaged")
			} else {
				log.Success("Files unstaged successfully")
			}
			return nil
		},
	}
}

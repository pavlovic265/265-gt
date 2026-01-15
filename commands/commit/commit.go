package commit

import (
	helpers "github.com/pavlovic265/265-gt/helpers"
	"github.com/pavlovic265/265-gt/runner"
	"github.com/pavlovic265/265-gt/utils/log"
	timeutils "github.com/pavlovic265/265-gt/utils/timeutils"
	"github.com/spf13/cobra"
)

type commitCommand struct {
	runner    runner.Runner
	gitHelper helpers.GitHelper
}

func NewCommitCommand(
	runner runner.Runner,
	gitHelper helpers.GitHelper,
) commitCommand {
	return commitCommand{
		runner:    runner,
		gitHelper: gitHelper,
	}
}

func (svc commitCommand) Command() *cobra.Command {
	var empty bool

	commitCmd := &cobra.Command{
		Use:     "commit",
		Aliases: []string{"cm"},
		Short:   "create commit",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := svc.gitHelper.EnsureGitRepository(); err != nil {
				return err
			}

			if empty {
				return svc.handleEmptyCommit()
			}

			if len(args) == 0 {
				return log.ErrorMsg("No commit message provided")
			}
			message := string(args[0])
			return svc.handleCommit(message)

		},
	}

	commitCmd.Flags().BoolVarP(&empty, "empty", "e", false, "create an empty commit")

	return commitCmd
}

func (svc commitCommand) handleEmptyCommit() error {
	message := timeutils.Now().Format(timeutils.LayoutUserFriendly)
	if err := svc.runner.Git("commit", "--allow-empty", "-m", message); err != nil {
		return log.Error("Failed to create empty commit", err)
	}

	log.Successf("Empty commit created: %s", message)
	return nil
}

func (svc commitCommand) handleCommit(message string) error {
	if err := svc.runner.Git("commit", "-m", message); err != nil {
		return log.Error("Failed to create commit", err)
	}

	log.Successf("Commit created: %s", message)
	return nil
}

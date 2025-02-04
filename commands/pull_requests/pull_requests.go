package pullrequests

import (
	"github.com/pavlovic265/265-gt/executor"
	"github.com/spf13/cobra"
)

type pullRequestCommand struct {
	exe executor.Executor
}

func NewPullRequestCommand(
	exe executor.Executor,
) pullRequestCommand {
	return pullRequestCommand{
		exe: exe,
	}
}

func (svc pullRequestCommand) Command() *cobra.Command {
	pullRequestCmd := &cobra.Command{
		Use:     "pull_request",
		Short:   "commands for pull request",
		Aliases: []string{"pr"},
	}

	pullRequestCmd.AddCommand(NewCreateCommand(svc.exe).Command())

	return pullRequestCmd
}

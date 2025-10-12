package pullrequests

import (
	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/executor"
	"github.com/spf13/cobra"
)

type pullRequestCommand struct {
	exe           executor.Executor
	configManager config.ConfigManager
}

func NewPullRequestCommand(
	exe executor.Executor,
	configManager config.ConfigManager,
) pullRequestCommand {
	return pullRequestCommand{
		exe:           exe,
		configManager: configManager,
	}
}

func (svc pullRequestCommand) Command() *cobra.Command {
	pullRequestCmd := &cobra.Command{
		Use:     "pull_request",
		Short:   "commands for pull request",
		Aliases: []string{"pr"},
	}

	pullRequestCmd.AddCommand(NewCreateCommand(svc.exe, svc.configManager).Command())
	pullRequestCmd.AddCommand(NewListCommand(svc.exe, svc.configManager).Command())

	return pullRequestCmd
}

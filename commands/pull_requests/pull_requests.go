package pullrequests

import (
	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/executor"
	helpers "github.com/pavlovic265/265-gt/git_helpers"
	"github.com/spf13/cobra"
)

type pullRequestCommand struct {
	runner        executor.Runner
	configManager config.ConfigManager
	gitHelper     helpers.GitHelper
}

func NewPullRequestCommand(
	runner executor.Runner,
	configManager config.ConfigManager,
	gitHelper helpers.GitHelper,
) pullRequestCommand {
	return pullRequestCommand{
		runner:        runner,
		configManager: configManager,
		gitHelper:     gitHelper,
	}
}

func (svc pullRequestCommand) Command() *cobra.Command {
	pullRequestCmd := &cobra.Command{
		Use:     "pull_request",
		Short:   "commands for pull request",
		Aliases: []string{"pr"},
	}

	pullRequestCmd.AddCommand(NewCreateCommand(svc.runner, svc.configManager, svc.gitHelper).Command())
	pullRequestCmd.AddCommand(NewListCommand(svc.runner, svc.configManager, svc.gitHelper).Command())

	return pullRequestCmd
}

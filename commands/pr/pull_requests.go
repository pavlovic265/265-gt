package pr

import (
	"github.com/pavlovic265/265-gt/client"
	"github.com/pavlovic265/265-gt/config"
	helpers "github.com/pavlovic265/265-gt/helpers"
	"github.com/pavlovic265/265-gt/runner"
	"github.com/spf13/cobra"
)

type pullRequestCommand struct {
	runner        runner.Runner
	configManager config.ConfigManager
	gitHelper     helpers.GitHelper
	cliClient     client.CliClient
}

func NewPullRequestCommand(
	runner runner.Runner,
	configManager config.ConfigManager,
	gitHelper helpers.GitHelper,
	cliClient client.CliClient,
) pullRequestCommand {
	return pullRequestCommand{
		runner:        runner,
		configManager: configManager,
		gitHelper:     gitHelper,
		cliClient:     cliClient,
	}
}

func (svc pullRequestCommand) Command() *cobra.Command {
	pullRequestCmd := &cobra.Command{
		Use:     "pull_request",
		Short:   "commands for pull requests",
		Aliases: []string{"pr"},
	}

	pullRequestCmd.AddCommand(
		NewCreateCommand(svc.runner, svc.configManager, svc.gitHelper, svc.cliClient).Command(),
	)
	pullRequestCmd.AddCommand(
		NewListCommand(svc.runner, svc.configManager, svc.gitHelper, svc.cliClient).Command(),
	)

	return pullRequestCmd
}

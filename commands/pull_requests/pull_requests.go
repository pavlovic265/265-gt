package pullrequests

import (
	"fmt"
	"os"

	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/executor"
	helpers "github.com/pavlovic265/265-gt/git_helpers"
	"github.com/spf13/cobra"
)

type pullRequestCommand struct {
	exe           executor.Executor
	configManager config.ConfigManager
	gitHelper     helpers.GitHelper
}

func NewPullRequestCommand(
	exe executor.Executor,
	configManager config.ConfigManager,
	gitHelper helpers.GitHelper,
) pullRequestCommand {
	return pullRequestCommand{
		exe:           exe,
		configManager: configManager,
		gitHelper:     gitHelper,
	}
}

func (svc pullRequestCommand) Command() *cobra.Command {
	pullRequestCmd := &cobra.Command{
		Use:     "pull_request",
		Short:   "commands for pull request",
		Aliases: []string{"pr"},
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if err := svc.gitHelper.EnsureGitRepository(); err != nil {
				fmt.Printf("%v\n", err)
				os.Exit(1)
			}
		},
	}

	pullRequestCmd.AddCommand(NewCreateCommand(svc.exe, svc.configManager).Command())
	pullRequestCmd.AddCommand(NewListCommand(svc.exe, svc.configManager, svc.gitHelper).Command())

	return pullRequestCmd
}

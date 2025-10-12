package pullrequests

import (
	"fmt"

	"github.com/pavlovic265/265-gt/client"
	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/executor"
	"github.com/spf13/cobra"
)

type createCommand struct {
	exe           executor.Executor
	configManager config.ConfigManager
}

func NewCreateCommand(
	exe executor.Executor,
	configManager config.ConfigManager,
) createCommand {
	return createCommand{
		exe:           exe,
		configManager: configManager,
	}
}

var draft bool

func (svc createCommand) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "create",
		Aliases: []string{"c"},
		Short:   "Create a pull request",
		RunE: func(cmd *cobra.Command, args []string) error {
			if draft {
				args = append([]string{"--draft"}, args...)
			}
			account := svc.configManager.GetActiveAccount()
			if !svc.configManager.HasActiveAccount() {
				return fmt.Errorf("no active account found")
			}

			err := client.Client[account.Platform].CreatePullRequest(args)
			if err != nil {
				return err
			}
			fmt.Println("Pull request created successfully")
			return nil
		},
	}

	cmd.Flags().BoolVarP(&draft, "draft", "d", false, "Create a draft pull request")

	return cmd
}

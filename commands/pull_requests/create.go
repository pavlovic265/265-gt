package pullrequests

import (
	"fmt"
	"os"

	"github.com/pavlovic265/265-gt/client"
	"github.com/pavlovic265/265-gt/executor"
	"github.com/spf13/cobra"
)

type createCommand struct {
	exe executor.Executor
}

func NewCreateCommand(
	exe executor.Executor,
) createCommand {
	return createCommand{
		exe: exe,
	}
}

func (svc *createCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:     "create",
		Aliases: []string{"c"},
		Short:   "Create a pull request",
		Run: func(cmd *cobra.Command, args []string) {
			err := client.GlobalClient.CreatePullRequest(args)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error executing git status: %v\n", err)
				os.Exit(1)
			}
		},
	}
}

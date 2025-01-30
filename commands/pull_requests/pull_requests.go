package pullrequests

import (
	"github.com/spf13/cobra"
)

var pullRequestCmd = &cobra.Command{
	Use:                "pullreq",
	Aliases:            []string{"pr"},
	Short:              "pull branch",
	DisableFlagParsing: true,
}

func NewPullRequestCommand() *cobra.Command {
	pullRequestCmd.AddCommand(NewCreatePullRequestCommand())

	return pullRequestCmd
}

package commands

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pavlovic265/265-gt/executor"
	"github.com/spf13/cobra"
)

type versionCommand struct {
	exe executor.Executor
}

func NewVersionCommand(
	exe executor.Executor,
) versionCommand {
	return versionCommand{
		exe: exe,
	}
}

func (svc versionCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "version of current build",
		RunE: func(cmd *cobra.Command, args []string) error {
			url := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases/latest", "pavlovic265", "265-gt")

			resp, err := http.Get(url)
			if err != nil {
				return err
			}
			defer resp.Body.Close()

			var result struct {
				TagName string `json:"tag_name"`
			}

			if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
				return err
			}

			fmt.Println(result.TagName)

			return nil
		},
	}
}

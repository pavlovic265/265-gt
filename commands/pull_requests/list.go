package pullrequests

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/pavlovic265/265-gt/client"
	"github.com/pavlovic265/265-gt/components"
	"github.com/pavlovic265/265-gt/executor"
	"github.com/spf13/cobra"
)

type listCommand struct {
	exe executor.Executor
}

func NewListCommand(
	exe executor.Executor,
) listCommand {
	return listCommand{
		exe: exe,
	}
}

func (svc listCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:     "list",
		Short:   "show list of pull request",
		Aliases: []string{"li"},
		RunE: func(cmd *cobra.Command, args []string) error {
			prs, err := client.GlobalClient.ListPullRequests(args)
			if err != nil {
				return err
			}
			return svc.selectPullRequest(prs)
		},
	}
}

func (svc listCommand) selectPullRequest(
	prs []client.PullRequest,
) error {
	var strPrs []string
	for _, pr := range prs {
		strPrs = append(strPrs, fmt.Sprintf("%d:%s", pr.Number, pr.Title))
	}
	initialModel := components.ListModel{
		AllChoices: strPrs,
		Choices:    strPrs,
		Cursor:     0,
		Query:      "",
	}

	program := tea.NewProgram(initialModel)

	if finalModel, err := program.Run(); err == nil {
		if m, ok := finalModel.(components.ListModel); ok && m.Selected != "" {
			splited := strings.Split(m.Selected, ":")
			prNumber, err := strconv.Atoi(splited[0])
			if err != nil {
				return fmt.Errorf("faild to get pr number id")
			}

			for _, pr := range prs {
				if prNumber == pr.Number {
					exec.Command("open", pr.URL).Start()
				}
			}
		}
	} else {
		return err
	}
	return nil
}

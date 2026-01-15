package commands

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/pavlovic265/265-gt/constants"
	helpers "github.com/pavlovic265/265-gt/helpers"
	"github.com/pavlovic265/265-gt/runner"
	"github.com/spf13/cobra"
)

type statusCommand struct {
	runner    runner.Runner
	gitHelper helpers.GitHelper
}

func NewStatusCommand(
	runner runner.Runner,
	gitHelper helpers.GitHelper,
) statusCommand {
	return statusCommand{
		runner:    runner,
		gitHelper: gitHelper,
	}
}

func (svc statusCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:                "status",
		Aliases:            []string{"st"},
		Short:              "git status",
		DisableFlagParsing: true,
		SilenceUsage:       true,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := svc.gitHelper.EnsureGitRepository(); err != nil {
				return err
			}

			gitArgs := append([]string{"status"}, args...)
			output, err := svc.runner.GitOutput(gitArgs...)
			if err != nil {
				return err
			}

			styledOutput := svc.styleGitStatus(output)
			fmt.Print(styledOutput)

			return nil
		},
	}
}

func (svc statusCommand) styleGitStatus(output string) string {
	lines := strings.Split(output, "\n")
	var styledLines []string

	branchStyle := lipgloss.NewStyle().
		Foreground(constants.Blue).
		Bold(true)

	headerStyle := lipgloss.NewStyle().
		Foreground(constants.Magenta).
		Bold(true)

	stagedHeaderStyle := lipgloss.NewStyle().
		Foreground(constants.Green).
		Bold(true)

	untrackedHeaderStyle := lipgloss.NewStyle().
		Foreground(constants.Red).
		Bold(true)

	modifiedStyle := lipgloss.NewStyle().
		Foreground(constants.Yellow)

	newFileStyle := lipgloss.NewStyle().
		Foreground(constants.Green)

	deletedStyle := lipgloss.NewStyle().
		Foreground(constants.Red)

	untrackedStyle := lipgloss.NewStyle().
		Foreground(constants.Red)

	renamedStyle := lipgloss.NewStyle().
		Foreground(constants.Yellow)

	helpStyle := lipgloss.NewStyle().
		Foreground(constants.BrightBlack)

	cleanStyle := lipgloss.NewStyle().
		Foreground(constants.Green)

	for _, line := range lines {
		styledLine := line

		if strings.Contains(line, "On branch") {
			parts := strings.Split(line, " ")
			if len(parts) >= 3 {
				branchName := strings.Join(parts[2:], " ")
				styledLine = branchStyle.Render(fmt.Sprintf("%s %s %s",
					parts[0], parts[1], branchName))
			}
		} else if strings.Contains(line, "nothing to commit, working tree clean") {
			styledLine = cleanStyle.Render(line)
		} else if strings.Contains(line, "Changes to be committed") {
			styledLine = stagedHeaderStyle.Render(line)
		} else if strings.Contains(line, "Changes not staged for commit") {
			styledLine = headerStyle.Render(line)
		} else if strings.Contains(line, "Untracked files") {
			styledLine = untrackedHeaderStyle.Render(line)
		} else if strings.Contains(line, "modified:") {
			parts := strings.Split(line, ":")
			if len(parts) >= 2 {
				status := parts[0]
				file := strings.TrimSpace(parts[1])
				styledLine = fmt.Sprintf("%s: %s",
					modifiedStyle.Render(status),
					file)
			}
		} else if strings.Contains(line, "new file:") {
			parts := strings.Split(line, ":")
			if len(parts) >= 2 {
				status := parts[0]
				file := strings.TrimSpace(parts[1])
				styledLine = fmt.Sprintf("%s: %s",
					newFileStyle.Render(status),
					file)
			}
		} else if strings.Contains(line, "deleted:") {
			parts := strings.Split(line, ":")
			if len(parts) >= 2 {
				status := parts[0]
				file := strings.TrimSpace(parts[1])
				styledLine = fmt.Sprintf("%s: %s",
					deletedStyle.Render(status),
					file)
			}
		} else if strings.Contains(line, "renamed:") {
			// Handle renamed files (format: "renamed: old -> new")
			parts := strings.Split(line, ":")
			if len(parts) >= 2 {
				status := parts[0]
				fileInfo := strings.TrimSpace(parts[1])
				styledLine = fmt.Sprintf("%s: %s",
					renamedStyle.Render(status),
					fileInfo)
			}
		} else if strings.Contains(line, "use \"git") {
			styledLine = helpStyle.Render(line)
		} else if strings.Contains(line, "\t") && !strings.Contains(line, ":") {
			// This is likely an untracked file (indented with tab, no colon)
			file := strings.TrimSpace(line)
			styledLine = fmt.Sprintf("\t%s", untrackedStyle.Render(file))
		}

		styledLines = append(styledLines, styledLine)
	}

	return strings.Join(styledLines, "\n")
}

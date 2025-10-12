package commands

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/pavlovic265/265-gt/constants"
	"github.com/pavlovic265/265-gt/executor"
	"github.com/spf13/cobra"
)

type statusCommand struct {
	exe executor.Executor
}

func NewStatusCommand(
	exe executor.Executor,
) statusCommand {
	return statusCommand{
		exe: exe,
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
			exeArgs := append([]string{"status"}, args...)

			output, err := svc.exe.WithGit().WithArgs(exeArgs).RunWithOutput()
			if err != nil {
				return err
			}

			// Style the git status output
			styledOutput := svc.styleGitStatus(output.String())
			fmt.Print(styledOutput)

			return nil
		},
	}
}

func (svc statusCommand) styleGitStatus(output string) string {
	lines := strings.Split(output, "\n")
	var styledLines []string

	// Define styles using the new color schema
	branchStyle := lipgloss.NewStyle().
		Foreground(constants.GetAnsiColor(4)). // Blue for branch info
		Bold(true)

	headerStyle := lipgloss.NewStyle().
		Foreground(constants.GetAnsiColor(3)). // Yellow for headers
		Bold(true)

	modifiedStyle := lipgloss.NewStyle().
		Foreground(constants.GetAnsiColor(1)) // Red for modified files

	newFileStyle := lipgloss.NewStyle().
		Foreground(constants.GetAnsiColor(2)) // Green for new files

	deletedStyle := lipgloss.NewStyle().
		Foreground(constants.GetAnsiColor(1)) // Red for deleted files

	untrackedStyle := lipgloss.NewStyle().
		Foreground(constants.GetAnsiColor(6)) // Cyan for untracked files

	helpStyle := lipgloss.NewStyle().
		Foreground(constants.GetAnsiColor(8)) // Dim for help text

	cleanStyle := lipgloss.NewStyle().
		Foreground(constants.GetAnsiColor(2)) // Green for clean status

	for _, line := range lines {
		styledLine := line

		// Style branch information
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
			styledLine = headerStyle.Render(line)
		} else if strings.Contains(line, "Changes not staged for commit") {
			styledLine = headerStyle.Render(line)
		} else if strings.Contains(line, "Untracked files") {
			styledLine = headerStyle.Render(line)
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

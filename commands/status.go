package commands

import (
	"fmt"
	"strings"

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

	for _, line := range lines {
		styledLine := line

		// Style branch information
		if strings.Contains(line, "On branch") {
			parts := strings.Split(line, " ")
			if len(parts) >= 3 {
				branchName := strings.Join(parts[2:], " ")
				styledLine = fmt.Sprintf("%s %s %s",
					parts[0], parts[1], branchName)
			}
		}

		// Style file status indicators
		if strings.Contains(line, "Changes to be committed") {
			styledLine = line
		} else if strings.Contains(line, "Changes not staged for commit") {
			styledLine = line
		} else if strings.Contains(line, "Untracked files") {
			styledLine = line
		} else if strings.Contains(line, "modified:") {
			parts := strings.Split(line, ":")
			if len(parts) >= 2 {
				styledLine = fmt.Sprintf("%s:%s",
					parts[0],
					parts[1])
			}
		} else if strings.Contains(line, "new file:") {
			parts := strings.Split(line, ":")
			if len(parts) >= 2 {
				styledLine = fmt.Sprintf("%s:%s",
					parts[0],
					parts[1])
			}
		} else if strings.Contains(line, "deleted:") {
			parts := strings.Split(line, ":")
			if len(parts) >= 2 {
				styledLine = fmt.Sprintf("%s:%s",
					parts[0],
					parts[1])
			}
		} else if strings.Contains(line, "use \"git add <file>...\" to include in what will be committed") {
			styledLine = line
		} else if strings.Contains(line, "use \"git restore --staged <file>...\" to unstage") {
			styledLine = line
		} else if strings.Contains(line, "use \"git add/rm <file>...\" to update what will be committed") {
			styledLine = line
		} else if strings.Contains(line, "use \"git restore <file>...\" to discard changes in working directory") {
			styledLine = line
		} else if strings.Contains(line, "\t") && !strings.Contains(line, ":") {
			// This is likely an untracked file (indented with tab, no colon)
			styledLine = fmt.Sprintf("\t%s", strings.TrimSpace(line))
		}

		styledLines = append(styledLines, styledLine)
	}

	return strings.Join(styledLines, "\n")
}

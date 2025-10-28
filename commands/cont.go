package commands

import (
	"os"

	"github.com/pavlovic265/265-gt/constants"
	"github.com/pavlovic265/265-gt/executor"
	helpers "github.com/pavlovic265/265-gt/git_helpers"
	"github.com/pavlovic265/265-gt/utils/log"
	"github.com/spf13/cobra"
)

type contCommand struct {
	exe       executor.Executor
	gitHelper helpers.GitHelper
}

func NewContCommand(
	exe executor.Executor,
	gitHelper helpers.GitHelper,
) contCommand {
	return contCommand{
		exe:       exe,
		gitHelper: gitHelper,
	}
}

func (svc contCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:   "cont",
		Short: "short for rebase --continue",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if err := svc.gitHelper.EnsureGitRepository(); err != nil {
				_ = log.Error("Not in a git repository", err)
				os.Exit(1)
			}
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			exeArgs := []string{"rebase", "--continue"}
			err := svc.exe.WithGit().WithArgs(exeArgs).Run()
			if err != nil {
				return log.Error("Failed to continue rebase", err)
			}

			if svc.gitHelper.IsRebaseInProgress() {
				return nil
			}

			parent, pErr := svc.gitHelper.GetPending(constants.ParentBranch)
			child, cErr := svc.gitHelper.GetPending(constants.ChildBranch)
			if pErr == nil && cErr == nil {
				if parent != "" && child != "" {
					if err := svc.gitHelper.SetParent(parent, child); err != nil {
						return log.Error("Failed to set parent branch relationship", err)
					}

					_ = svc.gitHelper.DeletePending(constants.ParentBranch)
					_ = svc.gitHelper.DeletePending(constants.ChildBranch)

					log.Success("Rebase completed and metadata updated")
				}
			}

			// Run stty sane to fix any terminal issues that might have occurred
			// This is especially useful when Git opens an editor (like vim) during rebase
			// that can mess up terminal display settings
			// Side effects: Resets any custom terminal settings to standard defaults
			// _ = exec.Command("stty", "sane").Run() // Ignore stty errors as they're not critical

			return nil
		},
	}
}

package commands

import (
	"fmt"
	"os"

	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/executor"
	"github.com/spf13/cobra"
)

type UpgradeCommand struct {
	exe executor.Executor
}

func NewUpgradeCommand(
	exe executor.Executor,
) UpgradeCommand {
	return UpgradeCommand{
		exe: exe,
	}
}

func (svc UpgradeCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:   "upgrade",
		Short: "upgrade of current build",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Get repository from environment variable
			repository := os.Getenv("GT_REPOSITORY")
			if repository == "" {
				return fmt.Errorf("GT_REPOSITORY environment variable not set")
			}

			installURL := fmt.Sprintf("https://raw.githubusercontent.com/%s/main/scripts/install.sh", repository)
			exeArgs := []string{"-c", fmt.Sprintf("curl -fsSL %s | bash", installURL)}
			err := svc.exe.WithName("bash").WithArgs(exeArgs).Run()
			if err != nil {
				return err
			}
			fmt.Println(config.SuccessIndicator("Tool upgrade successfully"))
			return nil
		},
	}
}

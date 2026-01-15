package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/pavlovic265/265-gt/client"
	"github.com/pavlovic265/265-gt/commands"
	"github.com/pavlovic265/265-gt/commands/account"
	"github.com/pavlovic265/265-gt/commands/auth"
	"github.com/pavlovic265/265-gt/commands/commit"
	createconfig "github.com/pavlovic265/265-gt/commands/create_config"
	pullrequests "github.com/pavlovic265/265-gt/commands/pull_requests"
	"github.com/pavlovic265/265-gt/commands/stack"
	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/executor"
	helpers "github.com/pavlovic265/265-gt/git_helpers"

	"github.com/spf13/cobra"
)

var (
	runner        = executor.NewRunner()
	configManager = config.NewDefaultConfigManager(runner)
	gitHelper     = helpers.NewGitHelper(runner)
)

func init() {
	client.InitCliClient(gitHelper)
}

const UNKNOWN_COMMAND_ERROR = "unknown command"

var rootCmd = &cobra.Command{
	Use: "gt",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		globalCfg, _ := configManager.LoadGlobalConfig()
		localCfg, _ := configManager.LoadLocalConfig()

		ctx := cmd.Context()
		if ctx == nil {
			ctx = context.Background()
		}

		cfg := config.NewConfigContext(globalCfg, localCfg)
		ctx = config.WithConfig(ctx, cfg)
		cmd.SetContext(ctx)

		if globalCfg != nil {
			gitHelper.CheckGTVersion(ctx)
		}
		return nil
	},
	PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
		cfg, ok := config.GetConfig(cmd.Context())
		if !ok {
			return nil
		}

		if cfg.IsDirty() {
			if err := configManager.SaveGlobalConfig(*cfg.Global); err != nil {
				return err
			}
		}

		if cfg.IsLocalDirty() && cfg.Local != nil {
			if err := configManager.SaveLocalConfig(*cfg.Local); err != nil {
				return err
			}
		}

		return nil
	},
	SilenceErrors: true,
	SilenceUsage:  true,
}

func passToGit(args []string) {
	fmt.Printf("Unknown command, passing to git: git %s\n", strings.Join(args, " "))

	if err := runner.Git(args...); err != nil {
		os.Exit(1)
	}
}

func main() {
	// Load .env file if it exists
	_ = godotenv.Load() // Ignore .env loading errors as the file is optional

	rootCmd.AddCommand(commands.NewAddCommand(runner, gitHelper).Command())
	rootCmd.AddCommand(commands.NewStatusCommand(runner, gitHelper).Command())
	rootCmd.AddCommand(commands.NewSwitchCommand(runner, gitHelper).Command())
	rootCmd.AddCommand(commands.NewPushCommand(runner, gitHelper).Command())
	rootCmd.AddCommand(commands.NewPullCommand(runner, gitHelper).Command())
	rootCmd.AddCommand(commands.NewCreateCommand(runner, gitHelper).Command())
	rootCmd.AddCommand(commands.NewContCommand(runner, gitHelper).Command())
	rootCmd.AddCommand(commands.NewCheckoutCommand(runner, gitHelper).Command())
	rootCmd.AddCommand(commands.NewMoveCommand(runner, gitHelper).Command())
	rootCmd.AddCommand(commands.NewDeleteCommand(runner, gitHelper).Command())
	rootCmd.AddCommand(commands.NewCleanCommand(runner, gitHelper).Command())
	rootCmd.AddCommand(commands.NewUpgradeCommand(runner, configManager).Command())
	rootCmd.AddCommand(commands.NewDownCommand(runner, gitHelper).Command())
	rootCmd.AddCommand(commands.NewUpCommand(runner, gitHelper).Command())
	rootCmd.AddCommand(commands.NewTrackCommand(runner, gitHelper).Command())
	rootCmd.AddCommand(commands.NewUnstageCommand(runner, gitHelper).Command())
	rootCmd.AddCommand(stack.NewStackCommand(runner, gitHelper).Command())
	rootCmd.AddCommand(commit.NewCommitCommand(runner, gitHelper).Command())
	rootCmd.AddCommand(pullrequests.NewPullRequestCommand(runner, configManager, gitHelper).Command())

	rootCmd.AddCommand(commands.NewVersionCommand(runner, configManager).Command())
	rootCmd.AddCommand(auth.NewAuthCommand(configManager).Command())
	rootCmd.AddCommand(account.NewAccountCommand(runner, configManager).Command())
	rootCmd.AddCommand(createconfig.NewConfigCommand(runner, configManager).Command())
	rootCmd.AddCommand(commands.NewCompletionCommand().Command())

	// Execute the command and handle unknown commands
	if err := rootCmd.Execute(); err != nil {
		// Check if this is an unknown command error
		if strings.Contains(err.Error(), UNKNOWN_COMMAND_ERROR) {
			// If no arguments provided, show help instead of passing to git
			if len(os.Args) <= 1 {
				if err := rootCmd.Help(); err != nil {
					fmt.Fprintf(os.Stderr, "Error showing help: %v\n", err)
				}
				return
			}
			// Pass the command to git
			passToGit(os.Args[1:])
		} else {
			// For other errors, print and exit
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	}
}

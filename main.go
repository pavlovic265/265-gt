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
	helpers "github.com/pavlovic265/265-gt/helpers"
	"github.com/pavlovic265/265-gt/runner"
	"github.com/pavlovic265/265-gt/version"

	"github.com/spf13/cobra"
)

var (
	run           = runner.NewRunner()
	configManager = config.NewDefaultConfigManager(run)
	gitHelper     = helpers.NewGitHelper(run)
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
			version.CheckGTVersion(ctx)
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

	if err := run.Git(args...); err != nil {
		os.Exit(1)
	}
}

func main() {
	// Load .env file if it exists
	_ = godotenv.Load() // Ignore .env loading errors as the file is optional

	rootCmd.AddCommand(commands.NewAddCommand(run, gitHelper).Command())
	rootCmd.AddCommand(commands.NewStatusCommand(run, gitHelper).Command())
	rootCmd.AddCommand(commands.NewSwitchCommand(run, gitHelper).Command())
	rootCmd.AddCommand(commands.NewPushCommand(run, gitHelper).Command())
	rootCmd.AddCommand(commands.NewPullCommand(run, gitHelper).Command())
	rootCmd.AddCommand(commands.NewCreateCommand(run, gitHelper).Command())
	rootCmd.AddCommand(commands.NewContCommand(run, gitHelper).Command())
	rootCmd.AddCommand(commands.NewCheckoutCommand(run, gitHelper).Command())
	rootCmd.AddCommand(commands.NewMoveCommand(run, gitHelper).Command())
	rootCmd.AddCommand(commands.NewDeleteCommand(run, gitHelper).Command())
	rootCmd.AddCommand(commands.NewCleanCommand(run, gitHelper).Command())
	rootCmd.AddCommand(commands.NewUpgradeCommand(run, configManager).Command())
	rootCmd.AddCommand(commands.NewDownCommand(run, gitHelper).Command())
	rootCmd.AddCommand(commands.NewUpCommand(run, gitHelper).Command())
	rootCmd.AddCommand(commands.NewTrackCommand(run, gitHelper).Command())
	rootCmd.AddCommand(commands.NewUnstageCommand(run, gitHelper).Command())
	rootCmd.AddCommand(stack.NewStackCommand(run, gitHelper).Command())
	rootCmd.AddCommand(commit.NewCommitCommand(run, gitHelper).Command())
	rootCmd.AddCommand(pullrequests.NewPullRequestCommand(run, configManager, gitHelper).Command())

	rootCmd.AddCommand(commands.NewVersionCommand(run, configManager).Command())
	rootCmd.AddCommand(auth.NewAuthCommand(configManager).Command())
	rootCmd.AddCommand(account.NewAccountCommand(run, configManager).Command())
	rootCmd.AddCommand(createconfig.NewConfigCommand(run, configManager).Command())
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

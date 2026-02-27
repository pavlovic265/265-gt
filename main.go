package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/pavlovic265/265-gt/client"
	"github.com/pavlovic265/265-gt/commands/account"
	"github.com/pavlovic265/265-gt/commands/auth"
	"github.com/pavlovic265/265-gt/commands/basic"
	"github.com/pavlovic265/265-gt/commands/branch"
	"github.com/pavlovic265/265-gt/commands/commit"
	createconfig "github.com/pavlovic265/265-gt/commands/createconfig"
	"github.com/pavlovic265/265-gt/commands/pr"
	"github.com/pavlovic265/265-gt/commands/remote"
	"github.com/pavlovic265/265-gt/commands/stack"
	"github.com/pavlovic265/265-gt/commands/utility"
	"github.com/pavlovic265/265-gt/config"
	helpers "github.com/pavlovic265/265-gt/helpers"
	"github.com/pavlovic265/265-gt/runner"
	"github.com/pavlovic265/265-gt/version"

	"github.com/spf13/cobra"
)

const UNKNOWN_COMMAND_ERROR = "unknown command"

type App struct {
	run           runner.Runner
	configManager config.ConfigManager
	gitHelper     helpers.GitHelper
	rootCmd       *cobra.Command
}

func (app *App) passToGit(args []string) {
	fmt.Printf("Unknown command, passing to git: git %s\n", strings.Join(args, " "))

	if err := app.run.Git(args...); err != nil {
		os.Exit(1)
	}
}

func NewApp() *App {
	run := runner.NewRunner()
	configManager := config.NewDefaultConfigManager(run)
	gitHelper := helpers.NewGitHelper(run)

	client.InitCliClient(gitHelper)

	rootCmd := &cobra.Command{
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

	app := &App{
		run:           run,
		configManager: configManager,
		gitHelper:     gitHelper,
		rootCmd:       rootCmd,
	}

	basic.RegisterCommands(app.rootCmd, app.run, app.gitHelper)
	branch.RegisterCommands(app.rootCmd, app.run, app.gitHelper)
	remote.RegisterCommands(app.rootCmd, app.run, app.gitHelper)
	utility.RegisterCommands(app.rootCmd, app.run, app.configManager)
	stack.RegisterCommands(app.rootCmd, app.run, app.gitHelper)
	commit.RegisterCommands(app.rootCmd, app.run, app.gitHelper)
	pr.RegisterCommands(app.rootCmd, app.run, app.configManager, app.gitHelper)
	auth.RegisterCommands(app.rootCmd, app.configManager)
	account.RegisterCommands(app.rootCmd, app.run, app.configManager)
	createconfig.RegisterCommands(app.rootCmd, app.run, app.configManager)

	return app
}

func (app *App) Run() {
	if err := app.rootCmd.Execute(); err != nil {
		// Check if this is an unknown command error
		if strings.Contains(err.Error(), UNKNOWN_COMMAND_ERROR) {
			// If no arguments provided, show help instead of passing to git
			if len(os.Args) <= 1 {
				if err := app.rootCmd.Help(); err != nil {
					fmt.Fprintf(os.Stderr, "Error showing help: %v\n", err)
				}
				return
			}
			// Pass the command to git
			app.passToGit(os.Args[1:])
		} else {
			// For other errors, print and exit
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	}
}

func main() {
	// Load .env file if it exists
	_ = godotenv.Load() // Ignore .env loading errors as the file is optional

	NewApp().Run()
}

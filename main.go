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
	"github.com/pavlovic265/265-gt/constants"
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
	cliClient     client.CliClient
	rootCmd       *cobra.Command
}

func (app *App) passToGit(args []string) {
	fmt.Printf("Unknown command, passing to git: git %s\n", strings.Join(args, " "))

	if err := app.run.Git(args...); err != nil {
		os.Exit(1)
	}
}

func NewApp() (*App, error) {
	run := runner.NewRunner()
	configManager := config.NewDefaultConfigManager(run)
	gitHelper := helpers.NewGitHelper(run)
	platform := constants.GitHubPlatform
	if globalCfg, err := configManager.LoadGlobalConfig(); err == nil &&
		globalCfg != nil &&
		globalCfg.ActiveAccount != nil &&
		globalCfg.ActiveAccount.Platform != "" {
		platform = globalCfg.ActiveAccount.Platform
	}

	cliClient, err := client.NewRestCliClient(platform, gitHelper)
	if err != nil {
		return nil, fmt.Errorf("failed to create cli client: %w", err)
	}

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
		cliClient:     cliClient,
		rootCmd:       rootCmd,
	}

	basic.RegisterCommands(app.rootCmd, app.run, app.gitHelper)
	branch.RegisterCommands(app.rootCmd, app.run, app.gitHelper)
	remote.RegisterCommands(app.rootCmd, app.run, app.gitHelper)
	utility.RegisterCommands(app.rootCmd, app.run, app.configManager)
	stack.RegisterCommands(app.rootCmd, app.run, app.gitHelper, app.cliClient)
	commit.RegisterCommands(app.rootCmd, app.run, app.gitHelper)
	pr.RegisterCommands(app.rootCmd, app.run, app.configManager, app.gitHelper, app.cliClient)
	auth.RegisterCommands(app.rootCmd, app.configManager, app.cliClient)
	account.RegisterCommands(app.rootCmd, app.run, app.configManager)
	createconfig.RegisterCommands(app.rootCmd, app.run, app.configManager)

	return app, nil
}

func (app *App) Run() {
	if err := app.rootCmd.Execute(); err != nil {
		if strings.Contains(err.Error(), UNKNOWN_COMMAND_ERROR) {
			if len(os.Args) <= 1 {
				if err := app.rootCmd.Help(); err != nil {
					fmt.Fprintf(os.Stderr, "Error showing help: %v\n", err)
				}
				return
			}
			app.passToGit(os.Args[1:])
		} else {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	}
}

func main() {
	_ = godotenv.Load()

	app, err := NewApp()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	app.Run()
}

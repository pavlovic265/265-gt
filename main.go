package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/joho/godotenv"
	"github.com/pavlovic265/265-gt/client"
	"github.com/pavlovic265/265-gt/commands"
	"github.com/pavlovic265/265-gt/commands/auth"
	"github.com/pavlovic265/265-gt/commands/commit"
	createconfig "github.com/pavlovic265/265-gt/commands/create_config"
	pullrequests "github.com/pavlovic265/265-gt/commands/pull_requests"
	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/executor"
	"github.com/pavlovic265/265-gt/helpers"

	"github.com/spf13/cobra"
)

var exe = executor.NewExe()
var configManager = config.NewDefaultConfigManager(exe)
var gitHelper = helpers.NewGitHelper(exe, configManager)

const UNKNOWN_COMMAND_ERROR = "unknown command"

var rootCmd = &cobra.Command{
	Use: "gt",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		client.InitCliClient(exe, configManager, gitHelper)

		isConfig := cmd.Parent() != nil && cmd.Parent().Name() == "config"
		isAuth := cmd.Parent() != nil && cmd.Parent().Name() == "auth"
		isVersion := cmd.Name() == "version"
		isCompletion := cmd.Name() == "completion"

		loadLocalConfig := false
		if isVersion || isCompletion || isConfig || isAuth {
			configManager.InitConfig(loadLocalConfig) // Don't load local config for these commands
			return
		}

		// Check if we're in a git repository
		if err := gitHelper.EnsureGitRepository(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		loadLocalConfig = true

		configManager.InitConfig(loadLocalConfig) // Load local config for these commands

		gitHelper.CheckGTVersion()
	},
	// Override the default error handling to pass unknown commands to git
	SilenceErrors: true,
	SilenceUsage:  true,
}

func passToGit(args []string) {
	fmt.Printf("Unknown command, passing to git: git %s\n", strings.Join(args, " "))

	err := exe.WithGit().WithArgs(args).Run()
	if err != nil {
		// If git command fails, exit with the same exit code
		if exitError, ok := err.(*exec.ExitError); ok {
			os.Exit(exitError.ExitCode())
		}
		// For other errors, exit with code 1
		os.Exit(1)
	}
}

func main() {
	// Load .env file if it exists
	_ = godotenv.Load() // Ignore .env loading errors as the file is optional

	rootCmd.AddCommand(commands.NewAddCommand(exe).Command())
	rootCmd.AddCommand(commands.NewStatusCommand(exe).Command())
	rootCmd.AddCommand(commands.NewSwitchCommand(exe).Command())
	rootCmd.AddCommand(commands.NewPushCommand(exe, gitHelper).Command())
	rootCmd.AddCommand(commands.NewPullCommand(exe, gitHelper).Command())
	rootCmd.AddCommand(commands.NewCreateCommand(exe, gitHelper).Command())
	rootCmd.AddCommand(commands.NewContCommand(exe).Command())
	rootCmd.AddCommand(commands.NewCheckoutCommand(exe, gitHelper).Command())
	rootCmd.AddCommand(commands.NewMoveCommand(exe, gitHelper).Command())
	rootCmd.AddCommand(commands.NewDeleteCommand(exe, gitHelper).Command())
	rootCmd.AddCommand(commands.NewCleanCommand(exe, gitHelper).Command())
	rootCmd.AddCommand(commands.NewVersionCommand(exe, configManager).Command())
	rootCmd.AddCommand(commands.NewUpgradeCommand(exe, configManager).Command())
	rootCmd.AddCommand(commands.NewDownCommand(exe, gitHelper).Command())
	rootCmd.AddCommand(commands.NewUpCommand(exe, gitHelper).Command())
	rootCmd.AddCommand(commands.NewUnstageCommand(exe).Command())
	rootCmd.AddCommand(commit.NewCommitCommand(exe).Command())
	rootCmd.AddCommand(pullrequests.NewPullRequestCommand(exe, configManager).Command())
	rootCmd.AddCommand(auth.NewAuthCommand(exe, configManager).Command())
	rootCmd.AddCommand(createconfig.NewConfigCommand(exe, configManager).Command())

	rootCmd.AddCommand(&cobra.Command{
		Use:   "completion [bash|zsh|fish|powershell]",
		Short: "Generate shell completion scripts",
		Long:  "Install auto-completion for Bash, Zsh, Fish, or PowerShell",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			var err error
			switch args[0] {
			case "bash":
				err = cmd.Root().GenBashCompletion(os.Stdout)
			case "zsh":
				err = cmd.Root().GenZshCompletion(os.Stdout)
			case "fish":
				err = cmd.Root().GenFishCompletion(os.Stdout, true)
			case "powershell":
				err = cmd.Root().GenPowerShellCompletion(os.Stdout)
			}
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error generating completion: %v\n", err)
				os.Exit(1)
			}
		},
	})

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

package commands

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
)

type CompletionCommand struct{}

func NewCompletionCommand() CompletionCommand {
	return CompletionCommand{}
}

func (c CompletionCommand) Command() *cobra.Command {
	var (
		install bool
		dir     string // optional custom output dir
		noDesc  bool   // omit descriptions where supported
	)

	cmd := &cobra.Command{
		Use:   "completion [bash|zsh|fish|powershell]",
		Short: "Generate shell completion scripts",
		Long: `Generate shell completion scripts for Bash, Zsh, Fish, or PowerShell.

By default, this prints the script to stdout. You can source it directly or redirect to a file.

Examples:
  # Print to stdout
  gt completion bash

  # Install to a default location for your shell
  gt completion zsh --install

  # Write to a custom directory
  gt completion fish --install --dir ~/.config/fish/completions
`,
		Args: cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
		ValidArgs: []string{
			"bash", "zsh", "fish", "powershell",
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			shell := args[0]
			if install {
				return installCompletion(cmd.Root(), shell, dir, !noDesc)
			}
			return printCompletion(cmd.Root(), shell, !noDesc, os.Stdout)
		},
	}

	cmd.Flags().BoolVar(&install, "install", false, "Write script to a sensible default location for the shell")
	cmd.Flags().StringVar(&dir, "dir", "", "Custom directory to write the completion file (implies --install)")
	cmd.Flags().BoolVar(&noDesc, "no-descriptions", false, "Disable command descriptions in completion where supported")

	return cmd
}

func printCompletion(root *cobra.Command, shell string, withDesc bool, w io.Writer) error {
	switch shell {
	case "bash":
		// V2 supports descriptions toggle
		return root.GenBashCompletionV2(w, withDesc)
	case "zsh":
		// zsh generator already includes descriptions
		return root.GenZshCompletion(w)
	case "fish":
		return root.GenFishCompletion(w, withDesc)
	case "powershell":
		return root.GenPowerShellCompletionWithDesc(w)
	default:
		return fmt.Errorf("unsupported shell: %s", shell)
	}
}

func installCompletion(root *cobra.Command, shell, customDir string, withDesc bool) error {
	name := root.CommandPath()
	tool := strings.Fields(name)
	bin := tool[len(tool)-1] // last token; typically "gt"

	outDir, outFile := defaultInstallLocation(shell, bin)
	if customDir != "" {
		outDir = expandHome(customDir)
	}

	if outDir == "" || outFile == "" {
		return fmt.Errorf("no default install path for shell %q on %s", shell, runtime.GOOS)
	}

	if err := os.MkdirAll(outDir, 0o755); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", outDir, err)
	}

	dst := filepath.Join(outDir, outFile)
	f, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("failed to create %s: %w", dst, err)
	}
	defer f.Close()

	if err := printCompletion(root, shell, withDesc, f); err != nil {
		return err
	}

	fmt.Printf("✓ Installed %s completion to %s\n", shell, dst)
	printPostInstallHint(shell, dst, outDir)
	return nil
}

// defaultInstallLocation returns (directory, filename) for a given shell/bin.
// These locations follow common conventions and do not modify rc files.
//
// Users can always override with --dir.
func defaultInstallLocation(shell, bin string) (string, string) {
	home, _ := os.UserHomeDir()
	switch shell {
	case "bash":
		// Prefer XDG if set, fall back to ~/.config
		cfg := os.Getenv("XDG_CONFIG_HOME")
		if cfg == "" {
			cfg = filepath.Join(home, ".config")
		}
		return filepath.Join(cfg, "bash", "completion"), fmt.Sprintf("%s.bash", bin)
	case "zsh":
		// A safe per-user location that can be added to fpath
		// (_filename, no extension) per zsh convention.
		dir := filepath.Join(home, ".zsh", "completions")
		return dir, "_" + bin
	case "fish":
		return filepath.Join(home, ".config", "fish", "completions"), fmt.Sprintf("%s.fish", bin)
	case "powershell":
		// Install next to the profile directory if possible.
		// We do NOT try to modify $PROFILE; users can dot-source.
		psDir := filepath.Join(home, "Documents", "PowerShell")
		// Windows PowerShell (<=5.1) profile folder:
		legacy := filepath.Join(home, "Documents", "WindowsPowerShell")
		// Prefer modern path if it exists, else legacy.
		if pathExists(psDir) {
			return psDir, fmt.Sprintf("%s-completion.ps1", bin)
		}
		return legacy, fmt.Sprintf("%s-completion.ps1", bin)
	default:
		return "", ""
	}
}

func printPostInstallHint(shell, dst, dir string) {
	switch shell {
	case "bash":
		fmt.Println("\nAdd this to your ~/.bashrc (if not already present):")
		fmt.Printf("  [ -f %q ] && source %q\n", dst, dst)
		fmt.Println("\nThen restart your shell or run:")
		fmt.Println("  source ~/.bashrc")
	case "zsh":
		fmt.Println("\nIf not already set, add your completions dir to fpath in ~/.zshrc:")
		fmt.Printf("  fpath=(%q $fpath)\n", dir)
		fmt.Println("Then enable and reload compsys:")
		fmt.Println("  autoload -Uz compinit && compinit")
	case "fish":
		fmt.Println("\nFish autoloads from ~/.config/fish/completions; just restart your shell.")
	case "powershell":
		fmt.Println("\nTo load completions in PowerShell, add this line to your $PROFILE and restart:")
		fmt.Printf("  . %q\n", dst)
	}
}

func expandHome(p string) string {
	if p == "" {
		return p
	}
	if strings.HasPrefix(p, "~") {
		home, _ := os.UserHomeDir()
		return filepath.Join(home, strings.TrimPrefix(p, "~"))
	}
	return p
}

func pathExists(p string) bool {
	_, err := os.Stat(p)
	return err == nil || !errors.Is(err, os.ErrNotExist)
}

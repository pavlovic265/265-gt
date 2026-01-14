package account

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/utils/log"
	"github.com/spf13/cobra"
)

type attachCommand struct {
	configManager config.ConfigManager
}

func NewAttachCommand(
	configManager config.ConfigManager,
) attachCommand {
	return attachCommand{
		configManager: configManager,
	}
}

func (atc attachCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:     "attach [directory]",
		Aliases: []string{"at"},
		Short:   "Attach active account to a directory",
		Long:    "Configure Git to use the active account's credentials for a specific directory",
		Args:    cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := config.RequireGlobal(cmd.Context())
			if err != nil {
				return err
			}

			if cfg.Global.ActiveAccount == nil || cfg.Global.ActiveAccount.User == "" {
				return log.ErrorMsg("No active account. Run 'gt auth' to set an active account")
			}

			// Get active account
			activeAccount := *cfg.Global.ActiveAccount

			// Determine target directory
			targetDir := "."
			if len(args) > 0 {
				targetDir = args[0]
			}

			// Resolve paths
			absPath, relPath, err := atc.resolvePaths(targetDir)
			if err != nil {
				return err
			}

			// Get home directory for global config
			homeDir, err := os.UserHomeDir()
			if err != nil {
				return log.Error("Failed to get home directory", err)
			}

			// 1. Update ~/.gitconfig with includeIf
			globalGitConfig := filepath.Join(homeDir, ".gitconfig")
			if err := atc.writeToGlobalGitConfig(globalGitConfig, relPath); err != nil {
				return err
			}

			// 2. Create/update .gitconfig in target directory
			localGitConfig := filepath.Join(absPath, ".gitconfig")
			if err := atc.writeToLocalGitConfig(localGitConfig, activeAccount); err != nil {
				return err
			}

			log.Successf("Attached account '%s' to %s", activeAccount.User, absPath)
			log.Infof("Git will use this account for all repositories in: %s", relPath)
			return nil
		},
	}
}

func (atc attachCommand) resolvePaths(targetDir string) (absPath string, relPath string, err error) {
	// Resolve to absolute path
	absPath, err = filepath.Abs(targetDir)
	if err != nil {
		return "", "", log.Error("Failed to resolve directory path", err)
	}

	// Ensure directory exists
	if _, err := os.Stat(absPath); os.IsNotExist(err) {
		return "", "", log.Errorf("Directory does not exist: %s", absPath)
	}

	// Get home directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", "", log.Error("Failed to get home directory", err)
	}

	// Convert to relative path from home if possible
	relPath, err = filepath.Rel(homeDir, absPath)
	if err != nil || strings.HasPrefix(relPath, "..") {
		// Use absolute path if not under home
		relPath = absPath
	} else {
		relPath = filepath.Join("~", relPath)
	}

	return absPath, relPath, nil
}

func (atc attachCommand) writeToGlobalGitConfig(gitConfigPath, dirPath string) error {
	// Prepare the gitdir pattern and local config path
	gitdirPattern := fmt.Sprintf("gitdir:%s/**/*", dirPath)
	localConfigPath := filepath.Join(dirPath, ".gitconfig")
	sectionHeader := fmt.Sprintf("[includeIf \"%s\"]", gitdirPattern)
	pathLine := fmt.Sprintf("\tpath = %s", localConfigPath)

	// Read existing content (empty if file doesn't exist)
	existingContent, err := os.ReadFile(gitConfigPath)
	if err != nil && !os.IsNotExist(err) {
		return log.Error("Failed to read ~/.gitconfig", err)
	}

	fileContent := string(existingContent)

	// Check if the exact section and path already exist
	if strings.Contains(fileContent, sectionHeader) && strings.Contains(fileContent, pathLine) {
		log.Info("includeIf section already exists in ~/.gitconfig")
		return nil
	}

	// Build new content
	var newContent strings.Builder
	newContent.WriteString(fileContent)

	// Add blank line if file has content
	if len(strings.TrimSpace(fileContent)) > 0 {
		newContent.WriteString("\n")
	}

	newContent.WriteString("# Added by gt\n")
	newContent.WriteString(sectionHeader + "\n")
	newContent.WriteString(pathLine + "\n")
	newContent.WriteString("# End of gt additions\n")

	// Write the file
	if err := os.WriteFile(gitConfigPath, []byte(newContent.String()), 0644); err != nil {
		return log.Error("Failed to write ~/.gitconfig", err)
	}

	log.Successf("Added includeIf to ~/.gitconfig")
	return nil
}

func (atc attachCommand) writeToLocalGitConfig(gitConfigPath string, account config.Account) error {
	// Read existing content (empty if file doesn't exist)
	existingContent, err := os.ReadFile(gitConfigPath)
	if err != nil && !os.IsNotExist(err) {
		return log.Error("Failed to read .gitconfig", err)
	}

	fileContent := string(existingContent)

	// Build new content by appending
	var newContent strings.Builder
	newContent.WriteString(fileContent)

	// Add blank line if file has content
	if len(strings.TrimSpace(fileContent)) > 0 {
		newContent.WriteString("\n")
	}

	// Append [user] section
	newContent.WriteString("# Added by gt\n")
	newContent.WriteString("[user]\n")
	if account.Name != "" {
		newContent.WriteString(fmt.Sprintf("\tname = %s\n", account.Name))
	}
	if account.Email != "" {
		newContent.WriteString(fmt.Sprintf("\temail = %s\n", account.Email))
	}
	if account.SigningKey != "" {
		newContent.WriteString(fmt.Sprintf("\tsigningkey = %s\n", account.SigningKey))
	}
	newContent.WriteString("# End of gt additions\n")

	// Write the file
	if err := os.WriteFile(gitConfigPath, []byte(newContent.String()), 0644); err != nil {
		return log.Error("Failed to write .gitconfig", err)
	}

	log.Successf("Created/updated %s", gitConfigPath)
	return nil
}

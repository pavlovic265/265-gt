package accounthelper

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/utils/log"
)

func AttachAccountToDir(account *config.Account, targetDir string) error {
	absPath, relPath, err := resolvePaths(targetDir)
	if err != nil {
		return err
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return log.Error("failed to get home directory", err)
	}

	globalGitConfig := filepath.Join(homeDir, ".gitconfig")
	if err := writeToGlobalGitConfig(globalGitConfig, relPath); err != nil {
		return err
	}

	localGitConfig := filepath.Join(absPath, ".gitconfig")
	if err := writeToLocalGitConfig(localGitConfig, account); err != nil {
		return err
	}

	log.Successf("Attached account '%s' to %s", account.User, absPath)
	return nil
}

func resolvePaths(targetDir string) (absPath string, relPath string, err error) {
	absPath, err = filepath.Abs(targetDir)
	if err != nil {
		return "", "", log.Error("failed to resolve directory path", err)
	}

	if _, err := os.Stat(absPath); os.IsNotExist(err) {
		return "", "", log.ErrorMsg(fmt.Sprintf("directory does not exist: %s", absPath))
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", "", log.Error("failed to get home directory", err)
	}

	relPath, err = filepath.Rel(homeDir, absPath)
	if err != nil || strings.HasPrefix(relPath, "..") {
		relPath = absPath
	} else {
		relPath = filepath.Join("~", relPath)
	}

	return absPath, relPath, nil
}

func writeToGlobalGitConfig(gitConfigPath, dirPath string) error {
	gitdirPattern := fmt.Sprintf("gitdir:%s/**/*", dirPath)
	localConfigPath := filepath.Join(dirPath, ".gitconfig")
	sectionHeader := fmt.Sprintf("[includeIf \"%s\"]", gitdirPattern)
	pathLine := fmt.Sprintf("\tpath = %s", localConfigPath)

	existingContent, err := os.ReadFile(gitConfigPath)
	if err != nil && !os.IsNotExist(err) {
		return log.Error("failed to read ~/.gitconfig", err)
	}

	fileContent := string(existingContent)

	if strings.Contains(fileContent, sectionHeader) && strings.Contains(fileContent, pathLine) {
		return nil
	}

	var newContent strings.Builder
	newContent.WriteString(fileContent)

	if len(strings.TrimSpace(fileContent)) > 0 {
		newContent.WriteString("\n")
	}

	newContent.WriteString("# Added by gt\n")
	newContent.WriteString(sectionHeader + "\n")
	newContent.WriteString(pathLine + "\n")
	newContent.WriteString("# End of gt additions\n")

	if err := os.WriteFile(gitConfigPath, []byte(newContent.String()), 0644); err != nil {
		return log.Error("failed to write ~/.gitconfig", err)
	}

	return nil
}

func writeToLocalGitConfig(gitConfigPath string, account *config.Account) error {
	existingContent, err := os.ReadFile(gitConfigPath)
	if err != nil && !os.IsNotExist(err) {
		return log.Error("failed to read .gitconfig", err)
	}

	fileContent := string(existingContent)

	var newContent strings.Builder
	newContent.WriteString(fileContent)

	if len(strings.TrimSpace(fileContent)) > 0 {
		newContent.WriteString("\n")
	}

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

	if err := os.WriteFile(gitConfigPath, []byte(newContent.String()), 0644); err != nil {
		return log.Error("failed to write .gitconfig", err)
	}

	return nil
}

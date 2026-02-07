package remote

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/constants"
	"github.com/pavlovic265/265-gt/helpers"
	"github.com/pavlovic265/265-gt/runner"
	"github.com/pavlovic265/265-gt/utils/log"
	"github.com/spf13/cobra"
)

type cloneCommand struct {
	runner    runner.Runner
	gitHelper helpers.GitHelper
}

func NewCloneCommand(
	runner runner.Runner,
	gitHelper helpers.GitHelper,
) cloneCommand {
	return cloneCommand{
		runner:    runner,
		gitHelper: gitHelper,
	}
}

func (svc cloneCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:   "clone <repo>",
		Short: "Clone a repository using active account's SSH config",
		Long: `Clone a repository using the active account's SSH configuration.

Examples:
  gt clone owner/repo              # Clone using active account's SSH host
  gt clone github.com/owner/repo   # Clone with explicit platform
  gt clone git@github.com:o/r.git  # Clone with full SSH URL (uses as-is)`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, ok := config.GetConfig(cmd.Context())
			if !ok {
				return log.ErrorMsg("config not loaded")
			}

			if cfg.Global == nil || cfg.Global.ActiveAccount == nil {
				return log.ErrorMsg("no active account - run 'gt auth switch' first")
			}

			account := cfg.Global.ActiveAccount
			repo := args[0]

			cloneURL, err := svc.buildCloneURL(repo, account)
			if err != nil {
				return err
			}

			log.Infof("Cloning %s...", cloneURL)
			if err := svc.runner.Git("clone", cloneURL); err != nil {
				return log.Error("failed to clone repository", err)
			}

			// Get the cloned directory name
			clonedDir := svc.extractRepoName(cloneURL)

			// Attach active account to the cloned directory
			if err := helpers.AttachAccountToDir(account, clonedDir); err != nil {
				log.Warningf("Could not attach account to cloned repo: %v", err)
			}

			log.Success("Repository cloned successfully")
			return nil
		},
	}
}

func (svc cloneCommand) buildCloneURL(repo string, account *config.Account) (string, error) {
	// If it's already a full SSH URL, use as-is
	if strings.HasPrefix(repo, "git@") {
		return repo, nil
	}

	// If it's an HTTPS URL, convert to SSH
	httpsRegex := regexp.MustCompile(`^https?://([^/]+)/(.+?)(?:\.git)?$`)
	if matches := httpsRegex.FindStringSubmatch(repo); matches != nil {
		platform := matches[1]
		path := strings.TrimSuffix(matches[2], ".git")

		sshHost := account.SSHHost
		if sshHost == "" {
			sshHost = platform
		}

		return fmt.Sprintf("git@%s:%s.git", sshHost, path), nil
	}

	// If it contains a platform prefix (github.com/owner/repo)
	platformRegex := regexp.MustCompile(`^(github\.com|gitlab\.com)/(.+?)(?:\.git)?$`)
	if matches := platformRegex.FindStringSubmatch(repo); matches != nil {
		platform := matches[1]
		path := strings.TrimSuffix(matches[2], ".git")

		sshHost := account.SSHHost
		if sshHost == "" {
			sshHost = platform
		}

		return fmt.Sprintf("git@%s:%s.git", sshHost, path), nil
	}

	// Simple format: owner/repo
	ownerRepoRegex := regexp.MustCompile(`^([^/]+)/([^/]+)$`)
	if matches := ownerRepoRegex.FindStringSubmatch(repo); matches != nil {
		owner := matches[1]
		repoName := strings.TrimSuffix(matches[2], ".git")

		sshHost := account.SSHHost
		if sshHost == "" {
			sshHost = constants.GitHubHost
			if account.Platform == constants.GitLabPlatform {
				sshHost = constants.GitLabHost
			}
		}

		return fmt.Sprintf("git@%s:%s/%s.git", sshHost, owner, repoName), nil
	}

	return "", log.ErrorMsg("invalid repository format - use owner/repo or full URL")
}

// extractRepoName extracts the repository name from a clone URL.
// git@github.com:owner/repo.git -> repo
// https://github.com/owner/repo -> repo
func (svc cloneCommand) extractRepoName(cloneURL string) string {
	// Remove .git suffix
	name := strings.TrimSuffix(cloneURL, ".git")

	// Get the last part after / or :
	if idx := strings.LastIndex(name, "/"); idx != -1 {
		name = name[idx+1:]
	} else if idx := strings.LastIndex(name, ":"); idx != -1 {
		name = name[idx+1:]
	}

	return name
}

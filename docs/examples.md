# Usage Examples

## Status Indicators
The tool uses beautiful ASCII icons and Panda Syntax colors to show operation status:

```bash
# Success indicators (✓) - Green
gt create feature/new-feature
# ✓ Branch 'feature/new-feature' created and switched to successfully

gt add
# ✓ All changes staged

gt commit "Add new feature"
# ✓ Commit created 'Add new feature'

gt checkout main
# ✓ Switched to branch 'main'

# Error indicators (✗) - Red
gt delete main
# ✗ Failed to delete branch: Cannot delete protected branch

gt checkout nonexistent
# ✗ Failed to checkout branch 'nonexistent': branch not found

# Version notifications - Blue with colored version numbers
gt status
# 🔄 A new release of gt is available: 0.29.0 → 0.32.0
# To upgrade, run: gt upgrade
```

## Styled Output Examples
```bash
# Git status with colored output
gt status
# On branch main                    # Branch name in purple
# Changes to be committed:          # Section headers in blue
#   new file:   README.md          # New files in green
# Changes not staged for commit:    # Section headers in orange
#   modified:   src/main.go        # Modified files in orange
# Untracked files:                  # Section headers in red
#   temp.txt                       # Untracked files in yellow italic

# Version information with styling
gt version
# Current version: v0.32.0         # Labels in blue, versions in blue bold

gt version --latest
# Latest version: v0.32.0          # Labels in blue, versions in blue bold
```

## Typical Workflow
```bash
# Start a new feature
gt create feature/user-authentication
gt add
gt commit "Add user authentication module"

# Make some changes and commit
gt add
gt commit "Add password validation"

# Oops, need to unstage some files
gt unstage file1.txt file2.txt
gt add file1.txt
gt commit "Add only file1.txt"

# Push to remote
gt push

# Create pull request
gt pr c

# After review, move up to parent branch
gt up

# Clean up merged branches
gt clean
```

## Branch Stack Navigation
```bash
# Current branch: feature/user-auth
gt up    # Move to: feature/auth-system
gt up    # Move to: develop
gt down  # Move to: feature/auth-system
gt down  # Move to: feature/user-auth
```

## Branch Tracking
```bash
# Interactively select a parent branch to track (no rebase)
gt track
# or use the alias
gt tr

# This will:
# 1. Display an interactive list of all branches (excluding current)
# 2. Allow searching/filtering branches by typing
# 3. Select a parent branch with Enter
# 4. Store parent relationship in git config (gt.branch.<branch>.parent)
# Note: This does NOT rebase, only sets the parent relationship

# Example workflow:
# You're on 'feature/new-feature' and want to track 'develop' as parent
gt track
# Search for 'develop', press Enter to select
# ✓ successfully tracking feature/new-feature

# Manual parent configuration (alternative to gt track):
# You can manually set the parent branch in git config
git config --local gt.branch.<branch-name>.parent <parent-branch>

# Example: Set 'main' as parent of current branch
git config --local gt.branch.$(git rev-parse --abbrev-ref HEAD).parent main

# View parent relationship
git config --local --get gt.branch.<branch-name>.parent

# Remove parent relationship
git config --local --unset gt.branch.<branch-name>.parent

# The parent relationship is used by:
# - gt up/down: Navigate branch hierarchy
# - gt stack restack: Rebase child branches onto their parents
# - Branch organization and visualization
```

## Multi-Account Management
```bash
# Add a new account
gt account add
# Follow the interactive prompts:
# 1. Enter username, email, full name
# 2. Select platform (GitHub/GitLab)
# 3. SSH key setup (create new or use existing)
# 4. Token setup (optional - with platform-specific guidance)

# List all configured accounts
gt account list
# * username1 (GitHub) user1@example.com - User One
#   username2 (GitLab) user2@example.com - User Two
# (* indicates active account)

# Edit an existing account
gt account edit
# Select account to edit, then update any field

# Quick update token only
gt account edit -t
# Select account, then enter new token

# Quick update GPG signing key only
gt account edit --gpg
# Select account, then enter GPG key ID

# Remove an account
gt account remove
# Select account to remove

# Switch between accounts
gt auth switch
# Select from available accounts

# Check current auth status
gt auth status

# Attach active account to a directory
gt account attach ~/work/projects
# This will:
# 1. Add an includeIf section to ~/.gitconfig
# 2. Create a .gitconfig in the target directory with active account credentials
# 3. All git repos under that directory will use the attached account

# Attach to current directory
gt account attach .

# The tool automatically manages an active account
# When you switch accounts, the active account is updated
# All operations use the currently active account
```

## Shell Completion
```bash
# Generate completion script for your shell
gt completion bash    # For Bash
gt completion zsh     # For Zsh
gt completion fish    # For Fish
gt completion powershell  # For PowerShell

# Install to default location (recommended)
gt completion bash --install
gt completion zsh --install
gt completion fish --install

# Install to custom directory
gt completion zsh --install --dir ~/.my-completions

# Generate without descriptions (faster, less verbose)
gt completion bash --no-descriptions
```

**Supported shells:**
- **Bash**: Installs to `~/.config/bash/completion/gt.bash`
- **Zsh**: Installs to `~/.zsh/completions/_gt`
- **Fish**: Installs to `~/.config/fish/completions/gt.fish`
- **PowerShell**: Installs to `~/Documents/PowerShell/gt-completion.ps1`

After installation, follow the instructions displayed to activate completions in your shell.

## Stack Restacking
```bash
# Restack all child branches from current branch
gt stack restack
# or use the alias
gt s rs

# This will:
# 1. Get all children of the current branch
# 2. Rebase each child onto its parent
# 3. Process children recursively to maintain branch hierarchy
```

## Enhanced Pull Request Management
```bash
# List all pull requests with visual indicators
gt pr li

# Output shows:
# ✓ 123: Add authentication module (CI passed)
# ✗ 124: Fix bug in parser (CI failed)
# ★ 125: Update dependencies (CI pending)

# Interactive actions in the PR list:
# - Enter: Open selected PR in browser
# - y: Copy PR URL to clipboard
# - m: Merge the selected PR
# - /: Search/filter PRs
# - Esc/q: Exit

# Quick workflow examples:
gt pr li
# Press 'y' on a PR to copy its URL for sharing
# Press 'm' on a PR to merge it directly from the terminal
```

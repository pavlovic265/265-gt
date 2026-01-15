# gt: Git Workflow Utility

[![MIT License](https://img.shields.io/badge/License-MIT-green.svg)](https://github.com/pavlovic265/265-gt/blob/main/LICENSE)
[![Go Version](https://img.shields.io/badge/Go-1.24+-blue.svg)](https://golang.org/)
[![Git Version](https://img.shields.io/badge/Git-2.39.5+-orange.svg)](https://git-scm.com/)

[![CI](https://github.com/pavlovic265/265-gt/actions/workflows/ci.yml/badge.svg)](https://github.com/pavlovic265/265-gt/actions/workflows/ci.yml)
[![Build](https://github.com/pavlovic265/265-gt/actions/workflows/build.yml/badge.svg)](https://github.com/pavlovic265/265-gt/actions/workflows/build.yml)
[![Lint](https://github.com/pavlovic265/265-gt/actions/workflows/lint.yml/badge.svg)](https://github.com/pavlovic265/265-gt/actions/workflows/lint.yml)
[![Test](https://github.com/pavlovic265/265-gt/actions/workflows/test.yml/badge.svg)](https://github.com/pavlovic265/265-gt/actions/workflows/test.yml)
[![Release](https://github.com/pavlovic265/265-gt/actions/workflows/release.yml/badge.svg)](https://github.com/pavlovic265/265-gt/actions/workflows/release.yml)

> A powerful command-line utility designed to simplify and streamline common Git workflows with intelligent branch management and automation.

## 🚀 Features

- **Quick Branch Operations**: Create, checkout, and switch between branches with shorthand commands
- **Branch Stack Navigation**: Seamlessly move up and down your branch hierarchy
- **Stack Management**: Restack branches to keep your branch hierarchy up to date
- **Automated Git Workflows**: Simplify rebasing, syncing, and cleaning up merged branches
- **Enhanced Pull Request Management**: Create, list, and merge PRs with CI/CD status indicators
- **Multi-Platform Support**: Works with GitHub and GitLab
- **Account Management**: Add, edit, list, and remove multiple accounts with full profile information
- **Authentication Management**: Easy account switching and token management
- **Interactive UI**: Beautiful terminal interface with search and selection capabilities
- **Shell Completion**: Auto-completion support for Bash, Zsh, Fish, and PowerShell
- **Panda Syntax Theme**: Stunning color scheme with dark/light theme support
- **Styled Output**: Consistent, beautiful error and success messages with icons
- **Enhanced Logging**: Centralized logging utility with styled output
- **Version Management**: Automatic version checking and upgrade notifications
- **Git Pass-Through**: Unknown commands automatically passed to git for seamless integration

## 🛠️ Installation

### npm (Recommended for cross-platform)
```bash
npm install -g @pavlovic265/gt
```

Supports Linux (x64, ARM64), macOS (Intel, Apple Silicon), and Windows (x64, ARM64).

### Homebrew (macOS)
```bash
brew tap pavlovic265/265-gt
brew install 265-gt
```

### Ubuntu/Debian (.deb package)
```bash
# Download and install the latest .deb package
wget https://github.com/pavlovic265/265-gt/releases/download/v0.3.0/gt_0.3.0_linux_amd64.deb
sudo apt install ./gt_0.3.0_linux_amd64.deb
```

### Fedora / RHEL / CentOS (.rpm package)
```bash
# Download and install the latest .rpm package
wget https://github.com/pavlovic265/265-gt/releases/download/v0.3.0/gt_0.3.0_linux_amd64.rpm
sudo rpm -i gt_0.3.0_linux_amd64.rpm
```



### Quick Install Script
```bash
curl -fsSL https://raw.githubusercontent.com/pavlovic265/265-gt/main/scripts/install.sh | bash
```

### Build from Source (For Developers)
```bash
# Clone the repository
git clone https://github.com/pavlovic265/265-gt.git
cd 265-gt

# Build the binary
go build -o gt main.go

# Install to your PATH (choose one):
# Option 1: User directory (no sudo needed)
mkdir -p ~/.local/bin
mv gt ~/.local/bin/
export PATH="$HOME/.local/bin:$PATH"

# Option 2: System directory (requires sudo)
sudo mv gt /usr/local/bin/
```

## 🎯 Quick Start

1. **Initialize Configuration**
   ```bash
   gt config global
   ```

2. **Add your first account**
   ```bash
   gt account add
   ```

3. **Authenticate with your Git platform**
   ```bash
   gt auth login
   ```

4. **Start using gt commands**
   ```bash
   gt create feature-branch    # Create and switch to a new branch
   gt up                       # Move up in branch stack
   gt down                     # Move down in branch stack
   ```

## 🔄 Git Pass-Through

**gt** automatically passes unknown commands to git, so you can use any git command through gt:

```bash
# These all work seamlessly
gt log --oneline -5
gt branch --list
gt diff --name-only
gt status
gt add .
gt commit -m "My commit"
```

When an unknown command is used, gt will show:
```
Unknown command, passing to git: git [command]
[git output]
```

This means you can use gt as a drop-in replacement for git while getting the benefits of gt's enhanced commands.


## 📖 Command Reference

### Branch Operations

| Command | Alias | Description | Example |
|---------|-------|-------------|---------|
| `create` | `c` | Create a new branch from current branch | `gt create feature-branch` |
| `checkout` | `co` | Checkout/search and switch to branch | `gt checkout main` |
| `delete` | `dl` | Delete a branch | `gt delete old-branch` |
| `clean` | `cl` | Clean merged branches (excludes protected) | `gt clean` |
| `move` | `mv` | Rebase current branch onto another branch | `gt move` |
| `track` | `tr` | Set parent branch relationship (no rebase) | `gt track` |

### Navigation

| Command | Alias | Description | Example |
|---------|-------|-------------|---------|
| `up` | - | Move up in branch stack | `gt up` |
| `down` | - | Move down in branch stack | `gt down` |
| `switch` | `sw` | Switch to previous branch | `gt switch` |
| `cont` | - | Continue rebase after resolving conflicts | `gt cont` |

### Commit Operations

| Command | Alias | Description | Example |
|---------|-------|-------------|---------|
| `add` | - | Stage changes | `gt add` |
| `unstage` | `us` | Unstage changes | `gt unstage` |
| `commit` | `cm` | Create commit with message | `gt commit "Add new feature"` |
| `commit -e` | `cm -e` | Create empty commit | `gt commit -e "WIP"` |

### Remote Operations

| Command | Alias | Description | Example |
|---------|-------|-------------|---------|
| `push` | `pu` | Force-push to remote | `gt push` |
| `pull` | `pl` | Pull latest changes | `gt pull` |

### Pull Request Management

| Command | Alias | Description | Example |
|---------|-------|-------------|---------|
| `pull_request create` | `pr c` | Create a new pull request | `gt pr c` |
| `pull_request create -d` | `pr c -d` | Create a draft pull request | `gt pr c -d` |
| `pull_request list` | `pr li` | List all pull requests with CI/CD status | `gt pr li` |

**Pull Request List Features:**
- **CI/CD Status Indicators**: View build status at a glance
  - `✓` (Green) - Success
  - `✗` (Red) - Failure/Error
  - `★` (Yellow) - Pending/In Progress
- **Interactive Actions**:
  - Press `Enter` to open PR in browser
  - Press `y` to yank (copy) PR URL to clipboard
  - Press `m` to merge the pull request

### Stack Management

| Command | Alias | Description | Example |
|---------|-------|-------------|---------|
| `stack restack` | `s rs` | Restack branches from current branch downward | `gt stack restack` |

### Configuration

| Command | Alias | Description | Example |
|---------|-------|-------------|---------|
| `config global` | `conf gl` | Configure global settings | `gt conf gl` |
| `config local` | `conf lo` | Configure local repository settings | `gt conf lo` |

### Authentication

| Command | Alias | Description | Example |
|---------|-------|-------------|---------|
| `auth login` | `auth lg` | Login with token from config | `gt auth lg` |
| `auth logout` | `auth lo` | Logout from current account | `gt auth lo` |
| `auth status` | `auth st` | Show authentication status | `gt auth st` |
| `auth switch` | `auth sw` | Switch between accounts | `gt auth sw` |

### Account Management

| Command | Alias | Description | Example |
|---------|-------|-------------|---------|
| `account add` | `acc add` | Add a new account interactively | `gt account add` |
| `account list` | `acc ls` | List all configured accounts | `gt account list` |
| `account edit` | `acc edit` | Edit an existing account | `gt account edit` |
| `account remove` | `acc rm` | Remove an account | `gt account remove` |
| `account attach` | `acc at` | Attach active account to a directory | `gt account attach ~/work` |

### Utility Commands

| Command | Alias | Description | Example |
|---------|-------|-------------|---------|
| `version` | `v` | Display version information | `gt version` |
| `version --latest` | `v -l` | Get latest version from repository | `gt version --latest` |
| `upgrade` | - | Update the CLI tool | `gt upgrade` |
| `status` | `st` | Show current repository status | `gt status` |
| `completion` | - | Generate shell completion scripts | `gt completion bash` |
| `completion --install` | - | Install completion script to default location | `gt completion zsh --install` |

## 🔧 Configuration

### Global Configuration
```bash
gt config global
```
This will guide you through setting up:
- Theme preferences (dark/light)
- Version tracking

After creating the global config, add your accounts:
```bash
gt account add
```

The account management commands allow you to:
- **Add accounts**: Interactively add GitHub/GitLab accounts with username, email, full name, platform, API token, and GPG signing key for commit signing
- **List accounts**: View all configured accounts with their details
- **Edit accounts**: Modify existing account information
- **Remove accounts**: Delete accounts you no longer need
- **Switch accounts**: Change between accounts using `gt auth switch`

The configuration supports multiple accounts with automatic active account management:
```yaml
# ~/.gtconfig.yaml
accounts:
  - user: "username1"
    email: "user1@example.com"
    name: "User One"
    token: "ghp_..."
    platform: "GitHub"
    signingkey: "ABCD1234EFGH5678"
  - user: "username2"
    email: "user2@example.com"
    name: "User Two"
    token: "glpat-..."
    platform: "GitLab"
    signingkey: "1234ABCD5678EFGH"
active_account:  # Automatically managed by auth commands
  user: "username1"
  email: "user1@example.com"
  name: "User One"
  token: "ghp_..."
  platform: "GitHub"
  signingkey: "ABCD1234EFGH5678"
theme:
  type: "dark"
version:
  last_checked: "2024-10-27T10:30:00Z"
  current_version: "0.3.0"
```

### Local Configuration
```bash
gt config local
```
Configure repository-specific settings like:
- Protected branches for this repo
- Custom branch naming patterns
- Repository-specific workflows

### Theme Configuration
The tool supports beautiful Panda Syntax theme with dark and light variants:

```yaml
# ~/.gtconfig.yaml
theme:
  type: "dark"  # or "light"
```

**Panda Syntax Colors:**
- 🟢 **Success**: `#A9DC52` (Green) / `#2D5016` (Dark Green - Light Theme)
- 🔴 **Error**: `#FF6188` (Red) / `#8B1538` (Dark Red - Light Theme)
- 🟡 **Warning**: `#FFD866` (Yellow) / `#B8860B` (Dark Orange - Light Theme)
- 🔵 **Info**: `#78DCE8` (Blue) / `#1E3A8A` (Dark Blue - Light Theme)
- 🟣 **Debug**: `#AB9DF2` (Purple) / `#6B46C1` (Dark Purple - Light Theme)
- 🟡 **Highlight**: `#FFD866` (Yellow) / `#B8860B` (Dark Yellow - Light Theme)

## 🎨 Usage Examples

### Status Indicators
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

### Styled Output Examples
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

### Typical Workflow
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

### Branch Stack Navigation
```bash
# Current branch: feature/user-auth
gt up    # Move to: feature/auth-system
gt up    # Move to: develop
gt down  # Move to: feature/auth-system
gt down  # Move to: feature/user-auth
```

### Branch Tracking
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

### Multi-Account Management
```bash
# Add a new account
gt account add
# Follow the interactive prompts to enter:
# - Username
# - Email
# - Full name
# - Platform (GitHub/GitLab)
# - API token
# - GPG signing key for commit signing (optional)

# List all configured accounts
gt account list
# * username1 (GitHub) user1@example.com - User One
#   username2 (GitLab) user2@example.com - User Two
# (* indicates active account)

# Edit an existing account
gt account edit
# Select account to edit, then update any field

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

### Shell Completion
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

### Stack Restacking
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

### Enhanced Pull Request Management
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

## 🚧 Planned Features

- [ ] **Branch Syncing** — Seamlessly synchronize local and remote branches with intelligent conflict handling.
- [ ] **Change Submission** — Streamlined `submit` command for creating pull requests or submitting changes for review.
- [ ] **Advanced Branch Visualization** — Enhanced visualization of branch structures and relationships for easier navigation.
- [x] **Multi-Platform Git Integration** — Full support for GitHub and GitLab with direct API calls (no external CLI dependencies).
- [ ] **Theme Customization** — Flexible theme settings to personalize the CLI experience.
- [ ] **Automated GitHub Setup** — One-command configuration for GitHub authentication, commit signing, tokens, and SSH keys.
- [ ] **User-Aware Repository Checkout** — Automatically clone and manage repositories based on the active user profile.


## 🤝 Contributing

We welcome contributions! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

### Development Setup
```bash
git clone https://github.com/pavlovic265/265-gt.git
cd 265-gt
go mod download
go run main.go
```

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](https://github.com/pavlovic265/265-gt/blob/main/LICENSE) file for details.

## 👨‍💻 Author

**Марко Павловић (Marko Pavlović)** - [@pavlovic265](https://github.com/pavlovic265)

## 🙏 Acknowledgments

- Built with [Cobra](https://github.com/spf13/cobra) for CLI framework
- Beautiful terminal UI powered by [Bubble Tea](https://github.com/charmbracelet/bubbletea)
- Styling with [Lip Gloss](https://github.com/charmbracelet/lipgloss)
- Color scheme inspired by [Panda Syntax](https://github.com/PandaTheme/panda-syntax-vscode) theme

---

⭐ If you find this tool helpful, please consider giving it a star!


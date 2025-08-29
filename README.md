# gt: Git Workflow Utility

[![MIT License](https://img.shields.io/badge/License-MIT-green.svg)](https://github.com/pavlovic265/265-gt/blob/main/LICENSE)
[![Go Version](https://img.shields.io/badge/Go-1.23+-blue.svg)](https://golang.org/)
[![Git Version](https://img.shields.io/badge/Git-2.39.5+-orange.svg)](https://git-scm.com/)

> A powerful command-line utility designed to simplify and streamline common Git workflows with intelligent branch management and automation.

## 🚀 Features

- **Quick Branch Operations**: Create, checkout, and switch between branches with shorthand commands
- **Branch Stack Navigation**: Seamlessly move up and down your branch hierarchy
- **Automated Git Workflows**: Simplify rebasing, syncing, and cleaning up merged branches
- **Pull Request Management**: Create and manage pull requests directly from the command line
- **Multi-Platform Support**: Works with GitHub and GitLab
- **Authentication Management**: Easy account switching and token management
- **Interactive UI**: Beautiful terminal interface with search and selection capabilities

## 📋 Requirements

- **Go**: 1.23 or higher
- **Git**: 2.39.5 or higher
- **GitHub CLI**: 2.66.1 or higher (for GitHub integration)

## 🛠️ Installation

### Quick Install
```bash
curl -fsSL https://raw.githubusercontent.com/pavlovic265/265-gt/main/scripts/install.sh | bash
```

### Manual Installation
```bash
# Clone the repository
git clone https://github.com/pavlovic265/265-gt.git
cd 265-gt

# Build and install
go build -o gt main.go
sudo mv gt /usr/local/bin/
```

## 🎯 Quick Start

1. **Initialize Configuration**
   ```bash
   gt config global
   ```

2. **Authenticate with your Git platform**
   ```bash
   gt auth login
   ```

3. **Start using gt commands**
   ```bash
   gt create feature-branch    # Create and switch to a new branch
   gt up                       # Move up in branch stack
   gt down                     # Move down in branch stack
   ```

## 📖 Command Reference

### Branch Operations

| Command | Alias | Description | Example |
|---------|-------|-------------|---------|
| `create` | `c` | Create a new branch from current branch | `gt create feature-branch` |
| `checkout` | `co` | Checkout/search and switch to branch | `gt checkout main` |
| `delete` | `dl` | Delete a branch | `gt delete old-branch` |
| `clean` | `cl` | Clean merged branches (excludes protected) | `gt clean` |
| `move` | `mv` | Rebase current branch | `gt move` |

### Navigation

| Command | Description | Example |
|---------|-------------|---------|
| `up` | Move up in branch stack | `gt up` |
| `down` | Move down in branch stack | `gt down` |
| `switch` | `sw` | Switch to previous branch | `gt switch` |

### Commit Operations

| Command | Alias | Description | Example |
|---------|-------|-------------|---------|
| `add` | - | Stage changes | `gt add` |
| `unstage` | `us` | Unstage changes | `gt unstage` |
| `commit` | `cm` | Create commit with message | `gt commit "Add new feature"` |
| `empty` | `em` | Create empty commit | `gt empty "WIP"` |

### Remote Operations

| Command | Alias | Description | Example |
|---------|-------|-------------|---------|
| `push` | `pu` | Force-push to remote | `gt push` |
| `pull` | `pl` | Pull latest changes | `gt pull` |

### Pull Request Management

| Command | Alias | Description | Example |
|---------|-------|-------------|---------|
| `pull_request create` | `pr c` | Create a new pull request | `gt pr c` |
| `pull_request list` | `pr li` | List all pull requests | `gt pr li` |

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

### Utility Commands

| Command | Alias | Description | Example |
|---------|-------|-------------|---------|
| `version` | `v` | Display version information | `gt version` |
| `update` | - | Update the CLI tool | `gt update` |
| `status` | - | Show current repository status | `gt status` |

## 🔧 Configuration

### Global Configuration
```bash
gt config global
```
This will guide you through setting up:
- Default Git platform (GitHub/GitLab)
- API tokens
- Default branch naming conventions
- Protected branches

### Local Configuration
```bash
gt config local
```
Configure repository-specific settings like:
- Protected branches for this repo
- Custom branch naming patterns
- Repository-specific workflows

## 🎨 Usage Examples

### Status Indicators
The tool uses beautiful ASCII icons to show operation status:

```bash
# Success indicators (✓)
gt create feature/new-feature
# ✓ Branch 'feature/new-feature' created and switched to successfully

gt auth login
# ✓ Successfully authenticated with username

gt auth status
# ✓ Authentication successful

# Error indicators (✗)
gt auth status
# ✗ Authentication failed

gt delete main
# ✗ Error: Cannot delete protected branch
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

### Multi-Account Management
```bash
# Switch between different accounts
gt auth switch
# Select from available accounts

# Check current auth status
gt auth status
```

## 🚧 Planned Features

- [ ] `sync` - Sync local/remote branches
- [ ] `submit` - Submit changes for review
- [ ] Enhanced branch visualization
- [ ] Integration with more Git platforms

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

**Marko Pavlović** - [@pavlovic265](https://github.com/pavlovic265)

## 🙏 Acknowledgments

- Built with [Cobra](https://github.com/spf13/cobra) for CLI framework
- Beautiful terminal UI powered by [Bubble Tea](https://github.com/charmbracelet/bubbletea)
- Styling with [Lip Gloss](https://github.com/charmbracelet/lipgloss)

---

⭐ If you find this tool helpful, please consider giving it a star!


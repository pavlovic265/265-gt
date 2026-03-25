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
- **Enhanced Pull Request Management**: Create, list, and merge PRs with CI/CD status, review approval, and merge conflict indicators
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

## 📖 Documentation

| Guide | Description |
|-------|-------------|
| [Command Reference](docs/commands.md) | All commands, aliases, flags, and options |
| [Configuration](docs/configuration.md) | Global/local config, themes, YAML reference |
| [Usage Examples](docs/examples.md) | Workflows, branch stacking, multi-account, shell completion |
| [Integrations](docs/integrations.md) | Powerlevel10k prompt setup |

## ⌨️ Interactive Keys

Interactive screens use modal navigation:

- `q` quits from normal mode
- `/` enters search mode in searchable lists
- `i` enters insert mode in text-input forms
- `Esc` exits search or insert mode
- `j` / `k` move down and up
- `h` / `l` move left and right where horizontal selection is available
- Arrow keys and `Tab` / `Shift+Tab` continue to work
- `Ctrl+C` remains available as an emergency quit

## 🧱 Code Structure

```text
.
├── commands/              # Cobra command implementations by domain
├── client/                # GitHub/GitLab API clients
├── config/                # Global/local/public config management
├── constants/             # Domain constants (platforms, config keys, enums)
├── helpers/               # Facade exports for helper subpackages
│   ├── githelper/         # Git workflow and branch relationship operations
│   ├── sshhelper/         # SSH key and ssh-config utilities
│   └── accounthelper/     # Per-directory git account attachment helpers
├── ui/
│   ├── components/        # Reusable TUI components (lists, prompts, inputs)
│   └── theme/             # UI colors, icons, and style helpers
├── runner/                # Shell/git command execution abstraction
├── utils/                 # Small utility packages (log, validate, pointer, time)
└── version/               # Version check and release notification logic
```

## 🚧 Planned Features

- [ ] **Branch Syncing** — Seamlessly synchronize local and remote branches with intelligent conflict handling.
- [ ] **Change Submission** — Streamlined `submit` command for creating pull requests or submitting changes for review.
- [ ] **Advanced Branch Visualization** — Enhanced visualization of branch structures and relationships for easier navigation.
- [x] **Multi-Platform Git Integration** — Full support for GitHub and GitLab with direct API calls (no external CLI dependencies).
- [ ] **Theme Customization** — Flexible theme settings to personalize the CLI experience.
- [x] **Automated GitHub Setup** — One-command configuration for GitHub authentication, commit signing, tokens, and SSH keys.
- [x] **User-Aware Repository Checkout** — Automatically clone and manage repositories based on the active user profile.


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

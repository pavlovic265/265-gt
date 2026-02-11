# Command Reference

## Branch Operations

| Command | Alias | Description | Example |
|---------|-------|-------------|---------|
| `create` | `c` | Create a new branch from current branch | `gt create feature-branch` |
| `checkout` | `co` | Checkout/search and switch to branch | `gt checkout main` |
| `checkout -r` | `co -r` | Checkout remote branch and track it | `gt co -r feature-branch` |
| `delete` | `dl` | Delete a branch | `gt delete old-branch` |
| `clean` | `cl` | Clean merged branches (excludes protected) | `gt clean` |
| `move` | `mv` | Rebase current branch onto another branch | `gt move` |
| `track` | `tr` | Set parent branch relationship (no rebase) | `gt track` |

## Navigation

| Command | Alias | Description | Example |
|---------|-------|-------------|---------|
| `up` | - | Move up in branch stack | `gt up` |
| `down` | - | Move down in branch stack | `gt down` |
| `switch` | `sw` | Switch to previous branch | `gt switch` |
| `cont` | - | Continue rebase after resolving conflicts | `gt cont` |

## Commit Operations

| Command | Alias | Description | Example |
|---------|-------|-------------|---------|
| `add` | - | Stage changes | `gt add` |
| `unstage` | `us` | Unstage changes | `gt unstage` |
| `commit` | `cm` | Create commit with message | `gt commit "Add new feature"` |
| `commit -e` | `cm -e` | Create empty commit | `gt commit -e "WIP"` |

## Remote Operations

| Command | Alias | Description | Example |
|---------|-------|-------------|---------|
| `push` | `pu` | Force-push to remote | `gt push` |
| `pull` | `pl` | Pull latest changes | `gt pull` |
| `pull --all` | `pl -a` | Pull from all remotes | `gt pl -a` |

## Pull Request Management

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
- **Review Approval Status**: See review state per PR
  - `●` (Green) - Approved
  - `●` (Orange) - No reviews yet
  - `●` (Red) - Changes requested
- **Merge Conflict Indicator**: `⚠` shown when the PR has merge conflicts
- **Interactive Actions**:
  - Press `Enter` to open PR in browser
  - Press `y` to yank (copy) PR URL to clipboard
  - Press `m` to merge the pull request

## Stack Management

| Command | Alias | Description | Example |
|---------|-------|-------------|---------|
| `stack restack` | `s rs` | Restack branches from current branch downward | `gt stack restack` |
| `submit-stack` | `ss` | Push and create PRs for the entire stack | `gt ss` |
| `submit-stack -d` | `ss -d` | Push and create draft PRs for the entire stack | `gt ss -d` |
| `submit-stack -i` | `ss -i` | Interactively choose per-branch action | `gt ss -i` |

## Configuration

| Command | Alias | Description | Example |
|---------|-------|-------------|---------|
| `config global` | `conf gl` | Configure global settings | `gt conf gl` |
| `config local` | `conf lo` | Configure local repository settings | `gt conf lo` |

## Authentication

| Command | Alias | Description | Example |
|---------|-------|-------------|---------|
| `auth login` | `auth lg` | Login with token from config | `gt auth lg` |
| `auth logout` | `auth lo` | Logout from current account | `gt auth lo` |
| `auth status` | `auth st` | Show authentication status | `gt auth st` |
| `auth switch` | `auth sw` | Switch between accounts | `gt auth sw` |

## Account Management

| Command | Alias | Description | Example |
|---------|-------|-------------|---------|
| `account add` | `acc add` | Add a new account interactively | `gt account add` |
| `account list` | `acc ls` | List all configured accounts | `gt account list` |
| `account edit` | `acc edit` | Edit an existing account | `gt account edit` |
| `account edit -t [token]` | `acc edit -t [token]` | Update token only (interactive if no value given) | `gt account edit -t ghp_xxx` |
| `account edit --gpg [key]` | `acc edit --gpg [key]` | Update GPG signing key only (interactive if no value given) | `gt account edit --gpg ABC123` |
| `account remove` | `acc rm` | Remove an account | `gt account remove` |
| `account attach` | `acc at` | Attach active account to a directory | `gt account attach ~/work` |

## Utility Commands

| Command | Alias | Description | Example |
|---------|-------|-------------|---------|
| `version` | `v` | Display version information | `gt version` |
| `version --latest` | `v -l` | Get latest version from repository | `gt version --latest` |
| `upgrade` | - | Update the CLI tool | `gt upgrade` |
| `status` | `st` | Show current repository status | `gt status` |
| `completion` | - | Generate shell completion scripts | `gt completion bash` |
| `completion --install` | - | Install completion script to default location | `gt completion zsh --install` |

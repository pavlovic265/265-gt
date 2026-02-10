# Configuration

## Global Configuration
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
- **Add accounts**: Interactively add GitHub/GitLab accounts with SSH key setup and optional token
- **List accounts**: View all configured accounts with their details
- **Edit accounts**: Modify existing account information or quickly update token (`-t [value]`) or GPG key (`--gpg [value]`). Pass a value directly to skip the interactive prompt
- **Remove accounts**: Delete accounts you no longer need
- **Switch accounts**: Change between accounts using `gt auth switch`

The configuration supports multiple accounts with automatic active account management:
```yaml
# ~/.config/gt/config.yml
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

## Local Configuration
```bash
gt config local
```
Configure repository-specific settings like:
- Protected branches for this repo
- Custom branch naming patterns
- Repository-specific workflows

## Theme Configuration
The tool supports beautiful Panda Syntax theme with dark and light variants:

```yaml
# ~/.config/gt/config.yml
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

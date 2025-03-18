
# gt: Git Workflow Utility
gt is a command-line utility designed to simplify and streamline common Git workflows. It provides shortcuts and automation for frequently used Git commands, making branch management, rebasing, and repository synchronization faster and more intuitive.

## Installation

Requirments:
- Go lang 1.23 version

Tested on
- Go lang 1.23 version
- Git 2.39.5  version
- Github cli 2.66.1 version

```bash
 curl -fsSL https://raw.githubusercontent.com/pavlovic265/265-gt/main/scripts/install.sh | bash
```
    

## Badges

TODO add more fancy badges

Add badges from somewhere like: [shields.io](https://shields.io/)

[![MIT License](https://img.shields.io/badge/License-MIT-green.svg)](https://github.com/pavlovic265/265-gt/blob/main/LICENSE)


## Features
**Create and switch branches quickly:** Use shorthand commands to create, checkout, and switch between branches.

**Branch stack navigation:** Move up and down your branch stack with ease.

**Automated Git workflows:** Simplify tasks like rebasing, syncing, and cleaning up merged branches.

**Force push and pull:** Safely force-push branches and pull updates from the remote repository.

**Customizable commands:** Add your own shortcuts or workflows as needed.

## Usage/Examples

### Command List



| Command/Alias               | Status | Description                                  |
|-----------------------------|--------|----------------------------------------------|
| **Branch Operations**       |        |                                              |
| `create` (`c`)              | ✅     | Create a new branch                          |
| `checkout` (`co`)           | ✅     | Checkout/search branch                       |
| `commit` (`cm`)             | ❌     | Create commit                                |
| `add`                       | ✅     | Stage changes (`git add`)                   |
|  `empty` (`em`)            | ❌     | Create empty commit                          |
|   `branch delete` (`branch dl`)           | ✅     | Delete branch                                |
|   `branch clean` (`branch cl`)            | ✅     | Interactive branch cleanup                  |
| `move` (`mv`)               | ✅     | Rebase branch                                |
| **Navigation**              |        |                                              |
| `bottom` (`bt`)             | ❌     | Move to bottom of stack                     |
| `up`                        | ❌     | Move up in stack                            |
| `switch` (`sw`)             | ✅     | Switch to previous branch                   |
| **Remote Operations**       |        |                                              |
| `push` (`pu`)               | ✅     | Force-push to remote                        |
| `pull` (`pl`)               | ✅     | Pull latest changes                         |
| `sync` (`sy`)               | ❌     | Sync local/remote branches                  |
| `submit` (`s`)              | ❌     | Submit changes                              |
| **Pull Requests**           |        | PR management                               |
|   `pull_request create` (`pr c`)            | ✅     | Create PR                                   |
|   `pull_request list` (`pr li`)             | ✅     | List PRs                                    |
| **Configuration**           |        |Config generation                          |
|   `config global` (`conf gl`)           | ✅     | Global config                               |
|   `config local` (`conf lo`)            | ✅     | Local config                                |
| **Auth Management**         |        |Account management                          |
|   `auth login` (`auth lg`)            | ✅     | Login user with token from config            |
|   `auth logout` (`auth lo`)           | ✅     | Logout user from config                     |
|   `auth status` (`auth st`)           | ✅     | Show auth status                            |
|   `auth switch` (`auth sw`)           | ✅     | Switch accounts                             |
| **Miscellaneous**           |        |                                              |
| `version` (`v`)             | ✅     | Check version                               |
| `update`                    | ✅     | Update CLI                                  |

## License

[MIT](https://github.com/pavlovic265/265-gt/blob/main/LICENSE)


## Authors

[@pavlovic265](https://github.com/pavlovic265)


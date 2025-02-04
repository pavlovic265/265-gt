
# gt: Git Workflow Utility
gt is a command-line utility designed to simplify and streamline common Git workflows. It provides shortcuts and automation for frequently used Git commands, making branch management, rebasing, and repository synchronization faster and more intuitive.


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

- <line color="green">✓</line> **`create` (`c`)**  
  Create a new branch.

- <line color="green">✓</line> **`checkout` (`co`)**  
  Check out a branch or search and checkout.

- <line color="red">✗</line> **`bottom` (`bt`)**  
  Move to the bottom branch in the stack.

- <line color="red">✗</line> **`up`**  
  Move to the upper branch in the stack.

- <line color="green">✓</line> **`status` (`st`)**  
  Run `git status`.

- <line color="green">✓</line> **`switch` (`sw`)**  
  Switch to the previous branch.

- <line color="green">✓</line> **`add`**  
  Stage changes (`git add`).

- <line color="green">✓</line> **`move` (`mv`)**  
  Rebase the current branch onto another branch.

- <line color="red">✗</line> **`sync` (`sy`)**  
  Sync local and remote branches.

- <line color="red">✗</line> **`submit` (`s`)**  
  Submit changes.

- <line color="green">✓</line> **`pull_request` (`pr`)**  
  Prefix for pull requests
  - <line color="green">✓</line> **`create` (`c`)**  
    Create a pull request

- <line color="red">✗</line> **`clean` (`cl`)**  
  Delete all branches that have been merged.

- <line color="green">✓</line> **`push` (`pu`)**  
  Force-push the current branch to the remote repository.

- <line color="green">✓</line> **`pull` (`pl`)**  
  Pull the latest changes for the current branch from the remote repository.

- <line color="green">✓</line> **`auth`**  
  Prefix to manage auth accounts.
    - <line color="green">✓</line> **`status` (`st`)**  
      Info about the current authenticated user.
    - <line color="green">✓</line> **`switch` (`sw`)**  
      Switch account to work with (e.g., personal or work).

- <line color="red">✗</line> **`conf`**  
  Generate global or local config
    - <line color="red">✗</line> **`global` (`gl`)**  
      Info about the current authenticated user.
    - <line color="red">✗</line> **`local` (`lo`)**  
      Switch account to work with (e.g., personal or work).
- <line color="red">✗</line> **`version` (`v`)**  
  Check current version

## License

[MIT](https://github.com/pavlovic265/265-gt/blob/main/LICENSE)


## Authors

[@pavlovic265](https://github.com/pavlovic265)


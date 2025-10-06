# Git Conflict Testing Scripts

This folder contains scripts to create and manage git conflicts for testing and learning purposes.

## Scripts

### `create_conflicts.sh`
Creates a rebase conflict scenario with the following structure:

1. **Create root branch** - Creates a `root-branch` with initial content
2. **Create sub-root branch** - Creates `sub-root-branch` from root with additional content
3. **Create sub-sub-root branch** - Creates `sub-sub-root-branch` from sub-root, modifying lines 2 and 3
4. **Modify root branch** - Modifies `root-branch` to change the same lines 2 and 3 (creating conflicts)
5. **Rebase sub-sub-root to root** - Attempts to rebase `sub-sub-root-branch` onto `root-branch`, causing conflicts

### `cleanup_conflicts.sh`
Cleans up all test branches and conflicts, returning the repository to a clean state.

## Usage

### Create Conflicts
```bash
./test/create_conflicts.sh
```

This will:
- Create the branch structure described above
- Generate rebase conflicts when trying to rebase `sub-sub-root-branch` onto `root-branch`
- Show instructions for resolving the conflicts

### Clean Up
```bash
./test/cleanup_conflicts.sh
```

This will:
- Abort any ongoing rebase or merge operations
- Delete all test branches
- Remove test files
- Reset the repository to a clean state

## Conflict Resolution Steps

When conflicts occur during rebase:

1. **Check status**: `git status`
2. **Open conflicted file**: Edit `test_file.txt`
3. **Look for conflict markers**:
   ```
   <<<<<<< HEAD
   content from root-branch
   =======
   content from sub-sub-root-branch
   >>>>>>> commit-hash
   ```
4. **Resolve conflicts**: Edit the file to keep desired content
5. **Stage resolved file**: `git add test_file.txt`
6. **Continue rebase**: `git rebase --continue`
7. **Repeat** until rebase is complete

## Branch Structure

### Original Hierarchy (before merge)
```
root-branch (base branch)
└── sub-root-branch (base branch of sub-sub-root-branch)
    └── sub-sub-root-branch
```

### After Merge and During Rebase
```
main (original)
├── root-branch (merged with sub-root content)
│   └── sub-root-branch (merged into root)
└── sub-sub-root-branch (currently rebasing onto root)
```

## Example Conflict Scenario

The script creates a scenario where:
- `root-branch` modifies lines 2 and 3 with "ROOT MODIFIED" content
- `sub-sub-root-branch` modifies the same lines 2 and 3 with "SUB-SUB-ROOT MODIFIED" content
- When you rebase `sub-sub-root-branch` onto `root-branch`, Git will detect conflicts on lines 2 and 3

## Testing with gt move

After running the script, you can test conflicts with the `gt move` command:

```bash
# Make sure you're on the sub-sub-root-branch
git checkout sub-sub-root-branch

# This will cause conflicts when rebasing onto root-branch
gt move root-branch
```

This is a realistic scenario that can occur in real development workflows when feature branches diverge and need to be rebased onto updated main branches.

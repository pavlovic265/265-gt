# Git Conflict Testing Scripts

This folder contains scripts to create and manage git conflicts for testing and learning purposes.

## Scripts

### `create_conflicts.sh`
Creates a rebase conflict scenario with the following structure:

1. **Create root-branch** - Creates a `root-branch` with 3 commits, adds file and content
2. **Create sub-root-branch** - Creates `sub-root-branch`, adds additional text to file
3. **Add conflicting commit** - Goes back to `root-branch`, adds 4th commit that modifies same lines (creating conflicts)
4. **Ready for testing** - Goes back to `sub-root-branch` ready for `gt move` command

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
```
main (original)
├── root-branch (modified lines 2 and 3)
└── sub-root-branch (original content on lines 2 and 3)
```

### During Rebase Conflict
```
main (original)
├── root-branch (target for rebase)
└── sub-root-branch (currently rebasing onto root-branch - CONFLICTS!)
```

## Example Conflict Scenario

The script creates a scenario where:
- `root-branch` has 4 commits total, with the 4th commit modifying lines 2 and 3 with "ROOT MODIFIED" content
- `sub-root-branch` has 1 commit with the original content on lines 2 and 3
- When you run `gt move root-branch` from `sub-root-branch`, Git will detect conflicts on lines 2 and 3 during the rebase

## Testing with gt move

After running the script, you can test conflicts with the `gt move` command:

```bash
# The script already puts you on sub-root-branch
# Just run the move command
gt move root-branch
```

This is a realistic scenario that can occur in real development workflows when feature branches diverge and need to be rebased onto updated main branches.

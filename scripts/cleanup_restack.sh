#!/bin/bash

# Script to clean up restack test scenario and reset to clean state

set -e

echo "🧹 Cleaning up restack test scenario..."

# Check if we're in a rebase state
if [ -d ".git/rebase-apply" ] || [ -d ".git/rebase-merge" ]; then
    echo "Aborting current rebase..."
    git rebase --abort 2>/dev/null || true
fi

# Check if we're in a merge state
if git status --porcelain | grep -q "^UU\|^AA\|^DD" 2>/dev/null; then
    echo "Aborting current merge..."
    git merge --abort 2>/dev/null || true
fi

# Switch to main/master before deleting branches
echo "Switching to main branch..."
git checkout main 2>/dev/null || git checkout master 2>/dev/null || true

# Delete test branches if they exist
echo "Deleting test branches..."
git branch -D feature-a 2>/dev/null || true
git branch -D feature-b 2>/dev/null || true
git branch -D feature-c 2>/dev/null || true

# Clean up git config entries for test branches
echo "Cleaning up git config..."
git config --local --unset gt.branch.feature-a.parent 2>/dev/null || true
git config --local --unset gt.branch.feature-b.parent 2>/dev/null || true
git config --local --unset gt.branch.feature-c.parent 2>/dev/null || true

# Remove test file if it exists
if [ -f "restack_test.txt" ]; then
    echo "Removing test file..."
    git rm restack_test.txt 2>/dev/null || rm restack_test.txt
    git commit -m "Remove restack test file" 2>/dev/null || true
fi

# Note: We do NOT reset uncommitted changes as they may be unrelated work
# If you have uncommitted test changes, manually discard them with: git checkout -- restack_test.txt

echo "✅ Cleanup completed! Test branches and config removed."
echo "ℹ️  Note: Any uncommitted changes have been preserved."

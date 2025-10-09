#!/bin/bash

# Script to clean up test conflicts and reset to clean state

set -e

echo "🧹 Cleaning up test conflicts..."

# Check if we're in a rebase state
if [ -d ".git/rebase-apply" ] || [ -d ".git/rebase-merge" ]; then
    echo "Aborting current rebase..."
    git rebase --abort
fi

# Check if we're in a merge state
if git status --porcelain | grep -q "^UU\|^AA\|^DD"; then
    echo "Aborting current merge..."
    git merge --abort
fi

# Delete test branches if they exist
echo "Deleting test branches..."
git branch -D root-branch 2>/dev/null || true
git branch -D sub-root-branch 2>/dev/null || true
git branch -D sub-sub-root-branch 2>/dev/null || true
git branch -D feature-branch 2>/dev/null || true

# Remove test file if it exists
if [ -f "test_file.txt" ]; then
    echo "Removing test file..."
    rm test_file.txt
    git add test_file.txt 2>/dev/null || true
    git commit -m "Remove test file" 2>/dev/null || true
fi

# Reset to clean state
echo "Resetting to clean state..."
git reset --hard HEAD

echo "✅ Cleanup completed! Repository is back to clean state."

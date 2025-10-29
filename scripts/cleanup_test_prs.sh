#!/bin/bash

# Script to clean up test PRs and branches

set -e

echo "Cleaning up test PRs and branches..."

# Get current branch to ensure we're not on a test branch
CURRENT_BRANCH=$(git rev-parse --abbrev-ref HEAD)

# Switch to main/master if currently on a test branch
if [[ "$CURRENT_BRANCH" == test-pr-* ]]; then
  echo "Currently on test branch, switching to main..."
  git checkout main 2>/dev/null || git checkout master 2>/dev/null
fi

# Delete local test branches
echo "Deleting local test branches..."
for i in 1 2 3 4; do
  BRANCH="test-pr-$i"
  if git show-ref --verify --quiet refs/heads/$BRANCH; then
    git branch -D $BRANCH
    echo "  ✓ Deleted local branch: $BRANCH"
  fi
done

# Delete remote test branches
echo "Deleting remote test branches..."
for i in 1 2 3 4; do
  BRANCH="test-pr-$i"
  if git ls-remote --exit-code --heads origin $BRANCH >/dev/null 2>&1; then
    git push origin --delete $BRANCH
    echo "  ✓ Deleted remote branch: $BRANCH"
  fi
done

# Delete test files if they exist
echo "Cleaning up test files..."
for i in 1 2 3 4; do
  FILE="test-file-$i.txt"
  if [ -f "$FILE" ]; then
    git rm -f $FILE 2>/dev/null || rm -f $FILE
    echo "  ✓ Deleted test file: $FILE"
  fi
done

# Check if there are changes to commit (deleted test files)
if ! git diff --quiet HEAD 2>/dev/null; then
  echo "Committing cleanup changes..."
  git add -A
  git commit -m "Clean up test files"
  echo "  ✓ Committed cleanup"
fi

echo ""
echo "✓ Cleanup complete!"
echo ""
echo "Test PRs, branches, and files have been removed."

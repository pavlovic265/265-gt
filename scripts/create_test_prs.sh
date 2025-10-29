#!/bin/bash

# Script to create 2-4 test PRs for testing merge and update functionality

set -e

# Get current branch to return to later
CURRENT_BRANCH=$(git rev-parse --abbrev-ref HEAD)

echo "Creating test PRs..."

# Create PR 1
echo "Creating PR 1..."
git checkout -b test-pr-1
echo "# Test PR 1" > test-file-1.txt
echo "This is a test file for PR 1" >> test-file-1.txt
git add test-file-1.txt
git commit -m "Add test file 1"
git push -u origin test-pr-1
gh pr create --title "Test PR 1: Add test file 1" --body "This is a test PR to verify merge functionality" --base "$CURRENT_BRANCH"

# Create PR 2
echo "Creating PR 2..."
git checkout "$CURRENT_BRANCH"
git checkout -b test-pr-2
echo "# Test PR 2" > test-file-2.txt
echo "This is a test file for PR 2" >> test-file-2.txt
git add test-file-2.txt
git commit -m "Add test file 2"
git push -u origin test-pr-2
gh pr create --title "Test PR 2: Add test file 2" --body "This is a test PR to verify merge functionality" --base "$CURRENT_BRANCH"

# Create PR 3 (this one will need an update)
echo "Creating PR 3..."
git checkout "$CURRENT_BRANCH"
git checkout -b test-pr-3
echo "# Test PR 3" > test-file-3.txt
echo "This is a test file for PR 3" >> test-file-3.txt
git add test-file-3.txt
git commit -m "Add test file 3"
git push -u origin test-pr-3
gh pr create --title "Test PR 3: Add test file 3" --body "This is a test PR to verify update functionality" --base "$CURRENT_BRANCH"

# Make a commit to base branch so PR 3 will be out of date
echo "Making PR 3 out of date..."
git checkout "$CURRENT_BRANCH"
echo "# Base change" > base-change.txt
echo "This commit makes PR 3 out of date" >> base-change.txt
git add base-change.txt
git commit -m "Add base change to make PR 3 out of date"
git push origin "$CURRENT_BRANCH"

# Create PR 4
echo "Creating PR 4..."
git checkout "$CURRENT_BRANCH"
git checkout -b test-pr-4
echo "# Test PR 4" > test-file-4.txt
echo "This is a test file for PR 4" >> test-file-4.txt
git add test-file-4.txt
git commit -m "Add test file 4"
git push -u origin test-pr-4
gh pr create --title "Test PR 4: Add test file 4" --body "This is a test PR to verify update functionality" --base "$CURRENT_BRANCH"

# Return to original branch
git checkout "$CURRENT_BRANCH"

echo ""
echo "✓ Successfully created 4 test PRs!"
echo ""
echo "You can now test:"
echo "  - Run 'gt pr li' to list the PRs"
echo "  - Press Ctrl+M on a PR to merge it"
echo "  - Press Ctrl+U on a PR to update its branch"
echo ""
echo "To clean up test branches later, run:"
echo "  git branch -D test-pr-1 test-pr-2 test-pr-3 test-pr-4"
echo "  git push origin --delete test-pr-1 test-pr-2 test-pr-3 test-pr-4"

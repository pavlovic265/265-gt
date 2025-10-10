#!/bin/bash

# Script to create rebase conflicts for testing conflict resolution
# This script creates a specific scenario:
# 1. Create root-branch, add file, add text to file
# 2. Create sub-root-branch, add additional text to file
# 3. Go back to root-branch, add also same file with same line
# 4. Go back to sub-root-branch and run move command to root-branch

set -e

echo "🔧 Creating rebase conflicts for testing..."

# Clean up any existing test branches
echo "Cleaning up existing test branches..."
git branch -D root-branch 2>/dev/null || true
git branch -D sub-root-branch 2>/dev/null || true

# Create a test file with initial content
echo "Creating test file with initial content..."
cat > test_file.txt << 'EOF'
This is the initial content of the test file.
Line 2: Some important information
Line 3: More content here
Line 4: Final line of content
EOF

git add test_file.txt
git commit -m "Add initial test file"

# Step 1: Create root-branch, add file, add text to file
echo "Step 1: Creating root-branch and adding multiple commits..."
git checkout -b root-branch

# First commit - Add some content to root branch
cat > test_file.txt << 'EOF'
This is the initial content of the test file.
Line 2: Some important information
Line 3: More content here
Line 4: Final line of content
Line 5: ROOT BRANCH - Additional content
Line 6: ROOT BRANCH - More modifications
EOF

git add test_file.txt
git commit -m "Add initial content to root branch"

# Second commit - Add more functionality
cat > test_file.txt << 'EOF'
This is the initial content of the test file.
Line 2: Some important information
Line 3: More content here
Line 4: Final line of content
Line 5: ROOT BRANCH - Additional content
Line 6: ROOT BRANCH - More modifications
Line 7: ROOT BRANCH - New feature added
Line 8: ROOT BRANCH - Bug fixes included
EOF

git add test_file.txt
git commit -m "Add new features and bug fixes to root branch"

# Third commit - Add configuration
cat > test_file.txt << 'EOF'
This is the initial content of the test file.
Line 2: Some important information
Line 3: More content here
Line 4: Final line of content
Line 5: ROOT BRANCH - Additional content
Line 6: ROOT BRANCH - More modifications
Line 7: ROOT BRANCH - New feature added
Line 8: ROOT BRANCH - Bug fixes included
Line 9: ROOT BRANCH - Configuration settings
Line 10: ROOT BRANCH - Performance improvements
EOF

git add test_file.txt
git commit -m "Add configuration and performance improvements to root branch"

echo "✅ root-branch created with 3 commits"

# Step 2: Create sub-root-branch, add additional text to file
echo "Step 2: Creating sub-root-branch and adding additional content..."
git checkout -b sub-root-branch

# Add additional content to sub-root branch
cat > test_file.txt << 'EOF'
This is the initial content of the test file.
Line 2: Some important information
Line 3: More content here
Line 4: Final line of content
Line 5: ROOT BRANCH - Additional content
Line 6: ROOT BRANCH - More modifications
Line 7: SUB-ROOT BRANCH - New functionality
Line 8: SUB-ROOT BRANCH - Additional features
EOF

git add test_file.txt
git commit -m "Add additional content to sub-root branch"

echo "✅ sub-root-branch created with additional content"

# Step 3: Go back to root-branch, add also same file with same line
echo "Step 3: Going back to root-branch and adding conflicting commit..."
git checkout root-branch

# Fourth commit - Modify the SAME lines that exist in sub-root branch - this creates conflicts
cat > test_file.txt << 'EOF'
This is the initial content of the test file.
Line 2: ROOT MODIFIED - Some important information
Line 3: ROOT MODIFIED - More content here
Line 4: Final line of content
Line 5: ROOT BRANCH - Additional content
Line 6: ROOT BRANCH - More modifications
Line 7: ROOT BRANCH - New feature added
Line 8: ROOT BRANCH - Bug fixes included
Line 9: ROOT BRANCH - Configuration settings
Line 10: ROOT BRANCH - Performance improvements
Line 11: ROOT BRANCH - Conflict resolution changes
Line 12: ROOT BRANCH - Final modifications
EOF

git add test_file.txt
git commit -m "Modify lines 2 and 3 in root branch (conflicts with sub-root)"

echo "✅ root-branch now has 4 commits with conflicting changes on lines 2 and 3"

# Step 4: Go back to sub-root-branch and run move command to root-branch
echo "Step 4: Going back to sub-root-branch..."
git checkout sub-root-branch

echo ""
echo "🎯 Conflict scenario created successfully!"
echo ""
echo "📋 Summary of what was created:"
echo "   - root-branch: 4 commits total, modified lines 2 and 3 with 'ROOT MODIFIED'"
echo "   - sub-root-branch: 1 commit, has original content on lines 2 and 3"
echo ""
echo "🔧 Now you can test conflicts with gt move:"
echo "   gt move root-branch"
echo ""
echo "⚠️  This will cause conflicts on lines 2 and 3 of test_file.txt"
echo "   because both branches modified the same lines differently"
echo ""
echo "🧹 To clean up: ./test/cleanup_conflicts.sh"

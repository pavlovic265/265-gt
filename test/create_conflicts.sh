#!/bin/bash

# Script to create rebase conflicts for testing conflict resolution
# This script creates a specific scenario:
# 1. Create root branch
# 2. Create sub-root branch (parent: root)
# 3. Create sub-sub-root branch (parent: sub-root)
# 4. Merge sub-root branch to root branch
# 5. Move sub-sub-root to root -> here we have conflicts

set -e

echo "🔧 Creating rebase conflicts for testing..."

# Clean up any existing test branches
echo "Cleaning up existing test branches..."
git branch -D root-branch 2>/dev/null || true
git branch -D sub-root-branch 2>/dev/null || true
git branch -D sub-sub-root-branch 2>/dev/null || true

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

# Step 1: Create root branch
echo "Step 1: Creating root branch..."
git checkout -b root-branch

# Add some content to root branch
cat > test_file.txt << 'EOF'
This is the initial content of the test file.
Line 2: Some important information
Line 3: More content here
Line 4: Final line of content
Line 5: ROOT BRANCH - Additional content
Line 6: ROOT BRANCH - More modifications
EOF

git add test_file.txt
git commit -m "Add content to root branch"

echo "✅ root-branch created (base branch)"

# Step 2: Create sub-root branch (parent: root)
echo "Step 2: Creating sub-root branch from root..."
git checkout -b sub-root-branch

# Add content to sub-root branch
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
git commit -m "Add content to sub-root branch"

echo "✅ sub-root-branch created (base branch of sub-sub-root-branch)"

# Step 3: Create sub-sub-root branch (parent: sub-root)
echo "Step 3: Creating sub-sub-root branch from sub-root..."
git checkout -b sub-sub-root-branch

# Add content to sub-sub-root branch
cat > test_file.txt << 'EOF'
This is the initial content of the test file.
Line 2: Some important information
Line 3: More content here
Line 4: Final line of content
Line 5: ROOT BRANCH - Additional content
Line 6: ROOT BRANCH - More modifications
Line 7: SUB-ROOT BRANCH - New functionality
Line 8: SUB-ROOT BRANCH - Additional features
Line 9: SUB-SUB-ROOT BRANCH - Advanced features
Line 10: SUB-SUB-ROOT BRANCH - Complex modifications
EOF

git add test_file.txt
git commit -m "Add content to sub-sub-root branch"

echo "✅ sub-sub-root-branch created"
echo ""
echo "🌳 Current branch hierarchy:"
echo "   root-branch (base branch)"
echo "   └── sub-root-branch (base branch of sub-sub-root-branch)"
echo "       └── sub-sub-root-branch"
echo ""

# Step 4: Merge sub-root branch to root branch
echo "Step 4: Merging sub-root branch to root branch..."
git checkout root-branch
git merge sub-root-branch --no-ff -m "Merge sub-root branch into root"

# Now root branch has the content from sub-root
echo "Root branch now contains merged content from sub-root"

# Step 5: Move sub-sub-root to root -> this will cause rebase conflicts
echo "Step 5: Attempting to rebase sub-sub-root onto root (this will cause conflicts)..."
git checkout sub-sub-root-branch

echo "⚠️  This will create rebase conflicts that need to be resolved!"

# The rebase command that will cause conflicts
git rebase root-branch || {
    echo "✅ Rebase conflicts created successfully!"
    echo ""
    echo "🔍 To resolve rebase conflicts:"
    echo "1. Check git status: git status"
    echo "2. Open the conflicted file: test_file.txt"
    echo "3. Look for conflict markers: <<<<<<< ======= >>>>>>>"
    echo "4. Edit the file to resolve conflicts"
    echo "5. Add resolved file: git add test_file.txt"
    echo "6. Continue rebase: git rebase --continue"
    echo "7. Repeat until rebase is complete"
    echo ""
    echo "📝 Current git status:"
    git status
    echo ""
    echo "📄 Conflicted file content:"
    cat test_file.txt
    echo ""
    echo "🌳 Branch structure after merge:"
    echo "   main (original)"
    echo "   ├── root-branch (merged with sub-root content)"
    echo "   │   └── sub-root-branch (merged into root)"
    echo "   └── sub-sub-root-branch (currently rebasing onto root)"
    echo ""
    echo "📋 Original hierarchy was:"
    echo "   root-branch (base branch)"
    echo "   └── sub-root-branch (base branch of sub-sub-root-branch)"
    echo "       └── sub-sub-root-branch"
}

echo ""
echo "🎯 Rebase conflict creation script completed!"
echo "You now have rebase conflicts to practice resolving."

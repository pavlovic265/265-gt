#!/bin/bash

# Script to create a restack scenario for testing
# This script creates a specific scenario:
# 1. Create a stacked chain of branches: main -> feature-a -> feature-b -> feature-c
# 2. Each branch has commits building on the previous
# 3. Make changes to main that should be restacked through the chain
# 4. Use gt restack to rebase the entire stack

set -e

echo "🔧 Creating restack scenario for testing..."

# Save current branch
ORIGINAL_BRANCH=$(git branch --show-current)

# Clean up any existing test branches
echo "Cleaning up existing test branches..."
git checkout main 2>/dev/null || git checkout master 2>/dev/null || true
git branch -D feature-a 2>/dev/null || true
git branch -D feature-b 2>/dev/null || true
git branch -D feature-c 2>/dev/null || true

# Create a base file on main
echo "Creating base file on main..."
cat > restack_test.txt << 'EOF'
Base File for Restack Testing
==============================
Line 1: Initial content
Line 2: Base functionality
Line 3: Core features
EOF

git add restack_test.txt
git commit -m "Add base file for restack testing"

# Create feature-a from main
echo "Step 1: Creating feature-a..."
git checkout -b feature-a

cat > restack_test.txt << 'EOF'
Base File for Restack Testing
==============================
Line 1: Initial content
Line 2: Base functionality
Line 3: Core features
Line 4: Feature A - New functionality
Line 5: Feature A - Additional code
EOF

git add restack_test.txt
git commit -m "Feature A: Add new functionality"

# Set parent relationship for feature-a
git config --local gt.branch.feature-a.parent main

echo "✅ feature-a created (based on main)"

# Create feature-b from feature-a
echo "Step 2: Creating feature-b..."
git checkout -b feature-b

cat > restack_test.txt << 'EOF'
Base File for Restack Testing
==============================
Line 1: Initial content
Line 2: Base functionality
Line 3: Core features
Line 4: Feature A - New functionality
Line 5: Feature A - Additional code
Line 6: Feature B - Extended features
Line 7: Feature B - More improvements
EOF

git add restack_test.txt
git commit -m "Feature B: Add extended features"

# Set parent relationship for feature-b
git config --local gt.branch.feature-b.parent feature-a

echo "✅ feature-b created (based on feature-a)"

# Create feature-c from feature-b
echo "Step 3: Creating feature-c..."
git checkout -b feature-c

cat > restack_test.txt << 'EOF'
Base File for Restack Testing
==============================
Line 1: Initial content
Line 2: Base functionality
Line 3: Core features
Line 4: Feature A - New functionality
Line 5: Feature A - Additional code
Line 6: Feature B - Extended features
Line 7: Feature B - More improvements
Line 8: Feature C - Final enhancements
Line 9: Feature C - Polish and cleanup
EOF

git add restack_test.txt
git commit -m "Feature C: Add final enhancements"

# Set parent relationship for feature-c
git config --local gt.branch.feature-c.parent feature-b

echo "✅ feature-c created (based on feature-b)"

# Go back to main and add a conflicting change
echo "Step 4: Adding new commits to main..."
git checkout main

cat > restack_test.txt << 'EOF'
Base File for Restack Testing
==============================
Line 1: MAIN BRANCH - Updated initial content
Line 2: MAIN BRANCH - Updated base functionality
Line 3: MAIN BRANCH - Updated core features
Line 4: MAIN BRANCH - New baseline feature
EOF

git add restack_test.txt
git commit -m "Main: Update baseline and add new features"

# Add another commit to main
cat > restack_test.txt << 'EOF'
Base File for Restack Testing
==============================
Line 1: MAIN BRANCH - Updated initial content
Line 2: MAIN BRANCH - Updated base functionality
Line 3: MAIN BRANCH - Updated core features
Line 4: MAIN BRANCH - New baseline feature
Line 5: MAIN BRANCH - Critical bug fix
Line 6: MAIN BRANCH - Performance improvements
EOF

git add restack_test.txt
git commit -m "Main: Add critical bug fix and performance improvements"

echo "✅ main branch updated with 2 new commits"

# Return to feature-a
git checkout feature-a

echo ""
echo "🎯 Restack scenario created successfully!"
echo ""
echo "📋 Branch structure created:"
echo "   main (2 new commits)"
echo "    └─ feature-a (1 commit)"
echo "        └─ feature-b (1 commit)"
echo "            └─ feature-c (1 commit)"
echo ""
echo "🔧 To test restack:"
echo "   1. On feature-a: gt restack (or gt stack restack)"
echo "      - This should rebase feature-a onto main"
echo "      - Then automatically rebase feature-b onto feature-a"
echo "      - Then automatically rebase feature-c onto feature-b"
echo ""
echo "💡 You can also test from any branch in the stack"
echo "   - From feature-b: gt restack → rebases b, then c"
echo "   - From feature-c: gt restack → rebases c only"
echo ""
echo "⚠️  The rebase will likely have conflicts that need resolution"
echo ""
echo "🧹 To clean up: ./scripts/cleanup_restack.sh"

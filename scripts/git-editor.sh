#!/bin/bash

# Smart Git Editor Script
# This script provides options for editing or skipping

# Check if we're in an interactive terminal
if [ -t 1 ]; then
    # We're in an interactive terminal, ask user what to do
    echo "Git wants to edit a commit message."
    echo "Options:"
    echo "1) Edit with nano (recommended)"
    echo "2) Edit with vim"
    echo "3) Skip editing (use default message)"
    echo "4) Edit with VS Code (if available)"
    echo ""
    read -p "Choose option (1-4): " choice
    
    case $choice in
        1)
            nano "$1"
            ;;
        2)
            vim "$1"
            ;;
        3)
            # Just exit without editing
            exit 0
            ;;
        4)
            if command -v code &> /dev/null; then
                code --wait "$1"
            else
                echo "VS Code not found, using nano instead"
                nano "$1"
            fi
            ;;
        *)
            echo "Invalid choice, using nano"
            nano "$1"
            ;;
    esac
else
    # We're not in an interactive terminal, use nano (safer than vim)
    nano "$1"
fi

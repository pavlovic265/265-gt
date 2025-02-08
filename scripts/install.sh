#!/bin/bash

# Set the repository and binary name
REPO="github.com/pavlovic265/265-gt"
BINARY="gt"
GITHUB_API_URL="https://api.github.com/repos/pavlovic265/265-gt/releases/latest"

# Fetch the latest GitHub release version using curl
LATEST_VERSION=$(curl -s $GITHUB_API_URL | grep '"tag_name"' | awk -F ': "' '{print $2}' | tr -d '",')

# Fallback if no version is found
if [ -z "$LATEST_VERSION" ]; then
  echo "Warning: Could not fetch latest version, defaulting to @latest"
  VERSION="@latest"
else
  VERSION="@$LATEST_VERSION"
fi

# Install the binary
echo "Installing $BINARY from $REPO version $VERSION..."
go install "$REPO$VERSION"

# Set symlink to executable
if [ ! -L /usr/local/bin/$BINARY ]; then
  ln -s ~/go/bin/265-gt /usr/local/bin/$BINARY
fi

# Verify installation
if command -v $BINARY &>/dev/null; then
  echo "$BINARY installed successfully."
  $BINARY version
else
  echo "Installation failed or $BINARY not found in PATH."
fi

#!/bin/bash

# Set the repository and binary name
REPO="github.com/pavlovic265/265-gt"
BINARY="gt"

# Set the version or use latest
VERSION="@latest"

# Install the binary
echo "Installing $BINARY from $REPO..."
go install "$REPO$VERSION"

# Set soft lick to executable
if [ ! -L /usr/local/bin/gt ]; then
  ln -s ~/go/bin/265-gt /usr/local/bin/gt
fi

# Verify installation
if command -v $BINARY &>/dev/null; then
  echo "$BINARY installed successfully."
  $BINARY --help
else
  echo "Installation failed or $BINARY not found in PATH."
fi

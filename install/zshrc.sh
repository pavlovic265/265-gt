#!/bin/bash

# Set the repository and binary name
REPO="github.com/pavlovic265/265-gt"
BINARY="gt"

# Set the version or use latest
VERSION="@latest"

# Install the binary
echo "Installing $BINARY from $REPO..."
go install "$REPO$VERSION"

# Ensure GOPATH/bin is in the PATH
GOBIN=$(go env GOPATH)/bin
if [[ ":$PATH:" != *":$GOBIN:"* ]]; then
  echo "Adding $GOBIN to PATH..."
  export PATH=$PATH:$GOBIN
  echo 'export PATH=$PATH:'"$GOBIN" >>~/.zahrc # Add to bashrc or use .zshrc for zsh
fi

# Verify installation
if command -v $BINARY &>/dev/null; then
  echo "$BINARY installed successfully."
  $BINARY --help
else
  echo "Installation failed or $BINARY not found in PATH."
fi

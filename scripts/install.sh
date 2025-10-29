#!/bin/bash

# Set the repository and binary name
REPO="github.com/pavlovic265/265-gt"
BINARY="gt"
GITHUB_API_URL="https://api.github.com/repos/pavlovic265/265-gt/releases/latest"

# Check if version/tag is provided as argument
if [ -n "$1" ]; then
  VERSION="@$1"
  echo "Installing specific version: $VERSION"
else
  # Fetch the latest GitHub release version using curl
  LATEST_VERSION=$(curl -s $GITHUB_API_URL | grep '"tag_name"' | awk -F ': "' '{print $2}' | tr -d '",')

  # Fallback if no version is found
  if [ -z "$LATEST_VERSION" ]; then
    echo "Warning: Could not fetch latest version, defaulting to @latest"
    VERSION="@latest"
  else
    VERSION="@$LATEST_VERSION"
  fi
fi

# Install the binary
echo "Installing $BINARY from $REPO version $VERSION..."
go install "$REPO$VERSION"

# Set symlink to executable in user's local bin directory
USER_BIN_DIR="$HOME/.local/bin"
mkdir -p "$USER_BIN_DIR"

if [ ! -L "$USER_BIN_DIR/$BINARY" ]; then
  ln -sf ~/go/bin/265-gt "$USER_BIN_DIR/$BINARY"
fi

# Add to PATH if not already there
if [[ ":$PATH:" != *":$USER_BIN_DIR:"* ]]; then
  echo "Adding $USER_BIN_DIR to PATH..."

  # Add to common shell config files
  for config_file in ~/.bashrc ~/.zshrc ~/.profile ~/.bash_profile; do
    if [ -f "$config_file" ]; then
      if ! grep -q 'export PATH="$HOME/.local/bin:$PATH"' "$config_file"; then
        echo 'export PATH="$HOME/.local/bin:$PATH"' >>"$config_file"
        echo "Added to $config_file"
      fi
    fi
  done

  echo "Please restart your shell or run: source ~/.bashrc (or ~/.zshrc)"
fi

# Verify installation
if command -v $BINARY &>/dev/null; then
  echo "$BINARY installed successfully."
  # Try to show version, but don't fail if it doesn't work (e.g., no config yet)
  if $BINARY version 2>/dev/null; then
    echo "Version check completed."
  else
    echo "Installation successful. Run '$BINARY config global' to set up configuration."
  fi
else
  echo "Installation failed or $BINARY not found in PATH."
fi

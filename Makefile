# Name of the executable
BINARY_NAME=gt

# Directory to store the build output
BUILD_DIR=build

# Load environment variables from .env file if it exists
ifneq (,$(wildcard .env))
    include .env
    export
endif

.PHONY: all build clean run lint lint-fix release patch minor major

# Default target, builds the application
all: build

# Build the project for the host platform
build:
	@echo "Building $(BINARY_NAME)..."
	mkdir -p $(BUILD_DIR)
	@if [ -n "$(GT_REPOSITORY)" ]; then \
		echo "Building with repository: $(GT_REPOSITORY)"; \
		GT_REPOSITORY=$(GT_REPOSITORY) go build -o $(BUILD_DIR)/$(BINARY_NAME) .; \
	else \
		echo "Building with default repository"; \
		go build -o $(BUILD_DIR)/$(BINARY_NAME) .; \
	fi
	@echo "Build complete."

# Run the built binary
run: build
	@echo "Running $(BINARY_NAME) with arguments $(args)..."
	./$(BUILD_DIR)/$(BINARY_NAME) $(args)

# Clean the build directory
clean:
	@echo "Cleaning up..."
	rm -rf $(BUILD_DIR)
	@echo "Cleanup complete."

# Run golangci-lint
lint:
	@echo "Running golangci-lint..."
	golangci-lint run
	@echo "Linting complete."

# Run golangci-lint with auto-fix
lint-fix:
	@echo "Running golangci-lint with auto-fix..."
	golangci-lint run --fix
	@echo "Linting with auto-fix complete."

release: build
	@echo "🔍 Fetching latest version from repository..."
	@LATEST_TAG=$$(./$(BUILD_DIR)/$(BINARY_NAME) version -l 2>/dev/null || echo "v0.0.0"); \
	echo "📋 Current version: $$LATEST_TAG"; \
	VERSION=$$(echo $$LATEST_TAG | sed 's/v//'); \
	MAJOR=$$(echo $$VERSION | cut -d. -f1); \
	MINOR=$$(echo $$VERSION | cut -d. -f2); \
	PATCH=$$(echo $$VERSION | cut -d. -f3); \
	if [ -n "$(bump)" ]; then \
		if [ "$(bump)" = "major" ]; then \
			NEW_VERSION="$$(( $$MAJOR + 1 )).0.0"; \
		elif [ "$(bump)" = "minor" ]; then \
			NEW_VERSION="$$MAJOR.$$(( $$MINOR + 1 )).0"; \
		elif [ "$(bump)" = "patch" ]; then \
			NEW_VERSION="$$MAJOR.$$MINOR.$$(( $$PATCH + 1 ))"; \
		else \
			echo "❌ Invalid bump type: $(bump)"; \
			echo "   Valid options: major, minor, patch"; \
			exit 1; \
		fi; \
		echo "🔄 Bumping $(bump) version..."; \
		echo "📦 New version: v$$NEW_VERSION"; \
		NEW_TAG="v$$NEW_VERSION"; \
	elif [ -n "$(v)" ]; then \
		echo "📦 Using specified version: $(v)"; \
		NEW_TAG="$(v)"; \
	else \
		echo "❌ Error: Please specify either:"; \
		echo "   make release bump=patch  # Patch version (0.0.1)"; \
		echo "   make release bump=minor  # Minor version (0.1.0)"; \
		echo "   make release bump=major  # Major version (1.0.0)"; \
		echo "   make release v=v1.2.3    # Specific version"; \
		exit 1; \
	fi; \
	echo "🚀 Creating release tag: $$NEW_TAG"; \
	git tag $$NEW_TAG; \
	git push origin $$NEW_TAG; \
	echo "✅ Release $$NEW_TAG created and pushed!"; \
	echo "🔗 Check GitHub Actions for build progress"

# Convenient aliases for common release types
patch:
	@$(MAKE) release bump=patch

minor:
	@$(MAKE) release bump=minor

major:
	@$(MAKE) release bump=major

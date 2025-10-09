# Name of the executable
BINARY_NAME=gt

# Directory to store the build output
BUILD_DIR=build

# Load environment variables from .env file if it exists
ifneq (,$(wildcard .env))
    include .env
    export
endif

.PHONY: all build clean run lint lint-fix

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

release:
	@if [ -z "$(v)" ]; then \
		echo "❌ Error: Tag version is required. Usage: make release v=v0.30.0"; \
		exit 1; \
	fi
	@echo "🚀 Starting release for tag $(v)..."
	git tag $(v)
	git push origin $(v)
	GITHUB_TOKEN=$(GH_TOKEN_PERSONAL) goreleaser release --clean
	@echo "✅ Release complete."

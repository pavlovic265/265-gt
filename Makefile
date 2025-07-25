# Name of the executable
BINARY_NAME=gt

# Directory to store the build output
BUILD_DIR=build

.PHONY: all build clean run

# Default target, builds the application
all: build

# Build the project for the host platform
build:
	@echo "Building $(BINARY_NAME)..."
	mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) .
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

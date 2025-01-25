# Name of the executable
BINARY_NAME=gt

# Build options
BUILD_DIR=build
SOURCES=$(shell find . -type f -name '*.go')

.PHONY: all build clean run

# Default target, builds the application
all: build

# Build the project
build: $(SOURCES)
	@echo "Building $(BINARY_NAME)..."
	GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME) .
	@echo "Build complete."

# Run the built binary
run: build
	@echo "Running $(BINARY_NAME)..."
	./$(BUILD_DIR)/$(BINARY_NAME)

# Clean the build directory
clean:
	@echo "Cleaning up..."
	rm  $(BUILD_DIR)/gt
	@echo "Cleanup complete."

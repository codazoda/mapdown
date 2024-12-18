# Makefile in src/

# Define the output directory
OUT_DIR=../bin

# Define the binary name
BINARY_NAME=mapdown

# Define the suffix for Windows binaries
WINDOWS_SUFFIX=.exe

# Default target: build all binaries
all: build

# Build all binaries
build:
	@mkdir -p $(OUT_DIR)
	@echo "Building for current OS/Architecture..."
	go build -o $(OUT_DIR)/$(BINARY_NAME) .
	@echo "Building for macOS/arm64..."
	GOOS=darwin GOARCH=arm64 go build -o $(OUT_DIR)/$(BINARY_NAME)-darwin-arm64 .
	@echo "Building for macOS/amd64..."
	GOOS=darwin GOARCH=amd64 go build -o $(OUT_DIR)/$(BINARY_NAME)-darwin-amd64 .
	@echo "Building for Windows/amd64..."
	GOOS=windows GOARCH=amd64 go build -o $(OUT_DIR)/$(BINARY_NAME)-windows-amd64$(WINDOWS_SUFFIX) .
	@echo "Building for Windows/arm64..."
	GOOS=windows GOARCH=arm64 go build -o $(OUT_DIR)/$(BINARY_NAME)-windows-arm64$(WINDOWS_SUFFIX) .
	@echo "Building for Windows/386..."
	GOOS=windows GOARCH=386 go build -o $(OUT_DIR)/$(BINARY_NAME)-windows-386$(WINDOWS_SUFFIX) .
	@echo "Building for Windows/arm..."
	GOOS=windows GOARCH=arm go build -o $(OUT_DIR)/$(BINARY_NAME)-windows-arm32$(WINDOWS_SUFFIX) .
	@echo "Building for Linux/amd64..."
	GOOS=linux GOARCH=amd64 go build -o $(OUT_DIR)/$(BINARY_NAME)-linux-amd64 .
	@echo "Building for Linux/amd64..."
	GOOS=linux GOARCH=amd64 go build -o $(OUT_DIR)/$(BINARY_NAME)-linux-arm64 .
	@echo "Building for Linux/amd32..."
	GOOS=linux GOARCH=386 go build -o $(OUT_DIR)/$(BINARY_NAME)-linux-386 .
	@echo "Building for Linux/arm32..."
	GOOS=linux GOARCH=arm go build -o $(OUT_DIR)/$(BINARY_NAME)-linux-arm32 .

	@echo "All builds completed successfully."

# Clean target: remove all built binaries
clean:
	@echo "Cleaning binaries from $(OUT_DIR)..."
	rm -f $(OUT_DIR)/*
	@echo "Clean complete."

.PHONY: all build clean

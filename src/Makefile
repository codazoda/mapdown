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
	GOOS=darwin GOARCH=arm64 go build -o $(OUT_DIR)/darwin-arm64/$(BINARY_NAME) .
	@echo "Building for macOS/amd64..."
	GOOS=darwin GOARCH=amd64 go build -o $(OUT_DIR)/darwin-amd64/$(BINARY_NAME) .
	@echo "Building for Windows/amd64..."
	GOOS=windows GOARCH=amd64 go build -o $(OUT_DIR)/windows-amd64/$(BINARY_NAME)$(WINDOWS_SUFFIX) .
	@echo "Building for Windows/arm64..."
	GOOS=windows GOARCH=arm64 go build -o $(OUT_DIR)/windows-arm64/$(BINARY_NAME)$(WINDOWS_SUFFIX) .
	@echo "Building for Windows/386..."
	GOOS=windows GOARCH=386 go build -o $(OUT_DIR)/windows-386/$(BINARY_NAME)$(WINDOWS_SUFFIX) .
	@echo "Building for Windows/arm..."
	GOOS=windows GOARCH=arm go build -o $(OUT_DIR)/windows-arm32/$(BINARY_NAME)$(WINDOWS_SUFFIX) .
	@echo "Building for Linux/amd64..."
	GOOS=linux GOARCH=amd64 go build -o $(OUT_DIR)/linux-amd64/$(BINARY_NAME) .
	@echo "Building for Linux/amd64..."
	GOOS=linux GOARCH=amd64 go build -o $(OUT_DIR)/linux-arm64/$(BINARY_NAME) .
	@echo "Building for Linux/amd32..."
	GOOS=linux GOARCH=386 go build -o $(OUT_DIR)/linux-386/$(BINARY_NAME) .
	@echo "Building for Linux/arm32..."
	GOOS=linux GOARCH=arm go build -o $(OUT_DIR)/linux-arm32/$(BINARY_NAME) .

	@echo "All builds completed successfully."

# Clean target: remove all built binaries
clean:
	@echo "Cleaning binaries from $(OUT_DIR)..."
	rm -rf -- $(OUT_DIR)/*
	@echo "Clean complete."

.PHONY: all build clean

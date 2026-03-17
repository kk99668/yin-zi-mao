# Makefile for yin-zi-mao

# Variables
BINARY_NAME=yin-zi-mao
VERSION?=v1.0.0
BUILD_TIME=$(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
LDFLAGS=-ldflags "-X main.version=$(VERSION) -X main.buildTime=$(BUILD_TIME)"
GO=go
GOFLAGS=-v

# Build directory
BUILD_DIR=build
RELEASE_DIR=release

# Platform-specific binaries
WINDOWS_AMD64=$(BINARY_NAME)-windows-amd64.exe
DARWIN_AMD64=$(BINARY_NAME)-darwin-amd64
DARWIN_ARM64=$(BINARY_NAME)-darwin-arm64
LINUX_AMD64=$(BINARY_NAME)-linux-amd64

.PHONY: all build clean test install release help

# Default target
all: build

## build: Compile binary for current platform
build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	$(GO) build $(GOFLAGS) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) .
	@echo "Build complete: $(BUILD_DIR)/$(BINARY_NAME)"
	@echo "Version: $(VERSION)"
	@echo "Build time: $(BUILD_TIME)"

## clean: Remove binary and release artifacts
clean:
	@echo "Cleaning..."
	@rm -rf $(BUILD_DIR)
	@rm -rf $(RELEASE_DIR)
	@echo "Clean complete"

## test: Run Go tests
test:
	@echo "Running tests..."
	$(GO) test -v ./...

## install: Install binary to GOPATH/bin
install: build
	@echo "Installing to $(GOPATH)/bin..."
	@cp $(BUILD_DIR)/$(BINARY_NAME) $(GOPATH)/bin/
	@echo "Install complete"

## release: Build cross-platform releases
release: clean
	@echo "Building releases..."
	@mkdir -p $(RELEASE_DIR)

	@echo "Building for Windows (amd64)..."
	@GOOS=windows GOARCH=amd64 $(GO) build $(LDFLAGS) -o $(RELEASE_DIR)/$(WINDOWS_AMD64) .

	@echo "Building for macOS (amd64)..."
	@GOOS=darwin GOARCH=amd64 $(GO) build $(LDFLAGS) -o $(RELEASE_DIR)/$(DARWIN_AMD64) .

	@echo "Building for macOS (arm64)..."
	@GOOS=darwin GOARCH=arm64 $(GO) build $(LDFLAGS) -o $(RELEASE_DIR)/$(DARWIN_ARM64) .

	@echo "Building for Linux (amd64)..."
	@GOOS=linux GOARCH=amd64 $(GO) build $(LDFLAGS) -o $(RELEASE_DIR)/$(LINUX_AMD64) .

	@echo "Release builds complete in $(RELEASE_DIR)/"
	@ls -lh $(RELEASE_DIR)/

## version: Display version information
version:
	@echo "Version: $(VERSION)"
	@echo "Build time: $(BUILD_TIME)"

## help: Display this help message
help:
	@echo "Available targets:"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

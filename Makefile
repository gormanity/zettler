# Define project variables
BINARY_NAME = zettler
BUILD_DIR = bin
VERSION := $(shell git describe --tags --always --dirty)
COMMIT := $(shell git rev-parse --short HEAD)
DATE := $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

# Go build settings
GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)
LDFLAGS = -ldflags "-X main.version=$(VERSION) -X main.commit=$(COMMIT) -X main.date=$(DATE) -s -w"

# Default target
.PHONY: all
all: build

## 🛠️ Build the project
.PHONY: build
build:
	@echo "🔨 Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	@go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) .

## 🧪 Run tests
.PHONY: test
test:
	@echo "🧪 Running tests..."
	@go test ./...

## 🎨 Format and lint the code
.PHONY: fmt
fmt:
	@echo "🎨 Formatting code..."
	@go fmt ./...
	@go vet ./...

## 🚀 Install the binary
.PHONY: install
install: build
	@echo "🚀 Installing $(BINARY_NAME)..."
	@mv $(BUILD_DIR)/$(BINARY_NAME) $$GOPATH/bin/

## 🗑️ Clean up build artifacts
.PHONY: clean
clean:
	@echo "🧹 Cleaning up..."
	@rm -rf $(BUILD_DIR)

## 📦 Build for multiple platforms
.PHONY: cross-build
cross-build:
	@echo "🌍 Building for all platforms..."
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-linux
	GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-macos
	GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME).exe

## 🏗️ Release (Build + Test + Package)
.PHONY: release
release: fmt test build

## 🔍 Help command
.PHONY: help
help:
	@echo "Makefile commands for $(BINARY_NAME):"
	@echo "  make build       - Build the binary"
	@echo "  make test        - Run tests"
	@echo "  make fmt         - Format and lint code"
	@echo "  make install     - Install binary to \$GOPATH/bin"
	@echo "  make clean       - Remove build artifacts"
	@echo "  make cross-build - Build for multiple platforms"
	@echo "  make release     - Format, test, and build"

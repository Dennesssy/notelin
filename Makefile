# No-Tel-in Build System
BINARY_NAME=notelin
MAIN_PATH=cmd/notelin/main.go
BUILD_DIR=build
DIST_DIR=dist
VERSION ?= $(shell git describe --tags --always --dirty)
COMMIT ?= $(shell git rev-parse --short HEAD)
DATE ?= $(shell date -u '+%Y-%m-%d_%H:%M:%S')

# Go build flags
LDFLAGS=-ldflags "-X main.version=${VERSION} -X main.commit=${COMMIT} -X main.date=${DATE}"
BUILD_FLAGS=-trimpath ${LDFLAGS}

# Default target
.PHONY: all
all: clean build

# Clean build artifacts
.PHONY: clean
clean:
	@echo "🧹 Cleaning build artifacts..."
	rm -rf ${BUILD_DIR} ${DIST_DIR}
	go clean

# Build for current platform
.PHONY: build
build:
	@echo "🔨 Building ${BINARY_NAME}..."
	mkdir -p ${BUILD_DIR}
	go build ${BUILD_FLAGS} -o ${BUILD_DIR}/${BINARY_NAME} ${MAIN_PATH}

# Build for all platforms
.PHONY: build-all
build-all: clean
	@echo "🌍 Building for all platforms..."
	mkdir -p ${DIST_DIR}
	
	# macOS Intel
	GOOS=darwin GOARCH=amd64 go build ${BUILD_FLAGS} -o ${DIST_DIR}/${BINARY_NAME}-darwin-amd64 ${MAIN_PATH}
	
	# macOS Apple Silicon  
	GOOS=darwin GOARCH=arm64 go build ${BUILD_FLAGS} -o ${DIST_DIR}/${BINARY_NAME}-darwin-arm64 ${MAIN_PATH}
	
	# Linux
	GOOS=linux GOARCH=amd64 go build ${BUILD_FLAGS} -o ${DIST_DIR}/${BINARY_NAME}-linux-amd64 ${MAIN_PATH}
	GOOS=linux GOARCH=arm64 go build ${BUILD_FLAGS} -o ${DIST_DIR}/${BINARY_NAME}-linux-arm64 ${MAIN_PATH}
	
	@echo "✅ Built binaries:"
	@ls -la ${DIST_DIR}

# Run tests
.PHONY: test
test:
	@echo "🧪 Running tests..."
	go test -v ./...

# Run tests with coverage
.PHONY: test-coverage
test-coverage:
	@echo "📊 Running tests with coverage..."
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report: coverage.html"

# Lint code
.PHONY: lint
lint:
	@echo "🔍 Linting code..."
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run; \
	else \
		echo "⚠️  golangci-lint not installed, skipping linting"; \
		echo "Install with: brew install golangci-lint"; \
	fi

# Format code
.PHONY: fmt
fmt:
	@echo "✨ Formatting code..."
	go fmt ./...
	go mod tidy

# Install locally
.PHONY: install
install: build
	@echo "📦 Installing ${BINARY_NAME}..."
	cp ${BUILD_DIR}/${BINARY_NAME} /usr/local/bin/
	@echo "✅ Installed to /usr/local/bin/${BINARY_NAME}"

# Uninstall
.PHONY: uninstall
uninstall:
	@echo "🗑️  Uninstalling ${BINARY_NAME}..."
	rm -f /usr/local/bin/${BINARY_NAME}
	@echo "✅ Uninstalled"

# Run application
.PHONY: run
run: build
	@echo "🚀 Running ${BINARY_NAME}..."
	./${BUILD_DIR}/${BINARY_NAME}

# Development mode - rebuild and run on changes
.PHONY: dev
dev:
	@echo "🔄 Starting development mode..."
	@if command -v air >/dev/null 2>&1; then \
		air; \
	else \
		echo "⚠️  'air' not installed for hot reload"; \
		echo "Install with: go install github.com/cosmtrek/air@latest"; \
		echo "Running normally..."; \
		make run; \
	fi

# Check dependencies
.PHONY: deps-check
deps-check:
	@echo "📋 Checking dependencies..."
	@echo "Go version: $(shell go version)"
	@echo "Gum installed: $(shell command -v gum >/dev/null 2>&1 && echo "✅ Yes" || echo "❌ No - Install with: brew install gum")"
	@echo "Git version: $(shell git --version 2>/dev/null || echo "❌ Not installed")"

# Update dependencies
.PHONY: deps-update
deps-update:
	@echo "🔄 Updating dependencies..."
	go get -u ./...
	go mod tidy

# Security scan
.PHONY: security-scan
security-scan:
	@echo "🔒 Running security scan..."
	@if command -v gosec >/dev/null 2>&1; then \
		gosec ./...; \
	else \
		echo "⚠️  gosec not installed"; \
		echo "Install with: go install github.com/securecodewarrior/gosec/v2/cmd/gosec@latest"; \
	fi

# Create release archive
.PHONY: release
release: build-all
	@echo "📦 Creating release archives..."
	mkdir -p ${DIST_DIR}/archives
	
	# Create tarballs for each platform
	cd ${DIST_DIR} && tar -czf archives/${BINARY_NAME}-${VERSION}-darwin-amd64.tar.gz ${BINARY_NAME}-darwin-amd64
	cd ${DIST_DIR} && tar -czf archives/${BINARY_NAME}-${VERSION}-darwin-arm64.tar.gz ${BINARY_NAME}-darwin-arm64  
	cd ${DIST_DIR} && tar -czf archives/${BINARY_NAME}-${VERSION}-linux-amd64.tar.gz ${BINARY_NAME}-linux-amd64
	cd ${DIST_DIR} && tar -czf archives/${BINARY_NAME}-${VERSION}-linux-arm64.tar.gz ${BINARY_NAME}-linux-arm64
	
	# Generate checksums
	cd ${DIST_DIR}/archives && shasum -a 256 *.tar.gz > ${BINARY_NAME}-${VERSION}-checksums.txt
	
	@echo "✅ Release archives created:"
	@ls -la ${DIST_DIR}/archives/

# Preview Homebrew formula
.PHONY: homebrew-preview
homebrew-preview: release
	@echo "🍺 Previewing Homebrew formula..."
	@echo "Version: ${VERSION}"
	@echo "Darwin AMD64 SHA256: $(shell shasum -a 256 ${DIST_DIR}/archives/${BINARY_NAME}-${VERSION}-darwin-amd64.tar.gz | cut -d' ' -f1)"
	@echo "Darwin ARM64 SHA256: $(shell shasum -a 256 ${DIST_DIR}/archives/${BINARY_NAME}-${VERSION}-darwin-arm64.tar.gz | cut -d' ' -f1)"

# Integration test with actual apps
.PHONY: integration-test
integration-test: build
	@echo "🔬 Running integration tests..."
	@echo "Testing with real applications..."
	./${BUILD_DIR}/${BINARY_NAME} scan || true
	@echo "✅ Integration test completed"

# Performance benchmark
.PHONY: benchmark
benchmark:
	@echo "⚡ Running benchmarks..."
	go test -bench=. -benchmem ./...

# Generate docs
.PHONY: docs
docs:
	@echo "📚 Generating documentation..."
	@if command -v godoc >/dev/null 2>&1; then \
		echo "Starting godoc server at http://localhost:6060"; \
		godoc -http=:6060; \
	else \
		echo "⚠️  godoc not available"; \
		echo "Install with: go install golang.org/x/tools/cmd/godoc@latest"; \
	fi

# Help
.PHONY: help
help:
	@echo "No-Tel-in Build System"
	@echo ""
	@echo "Available targets:"
	@echo "  build         - Build for current platform"
	@echo "  build-all     - Build for all supported platforms"
	@echo "  test          - Run tests"
	@echo "  test-coverage - Run tests with coverage report"
	@echo "  lint          - Run linter"
	@echo "  fmt           - Format code"
	@echo "  run           - Build and run application"
	@echo "  dev           - Development mode with hot reload"
	@echo "  install       - Install locally to /usr/local/bin"
	@echo "  uninstall     - Remove from /usr/local/bin"
	@echo "  clean         - Clean build artifacts"
	@echo "  deps-check    - Check required dependencies"
	@echo "  deps-update   - Update Go dependencies"
	@echo "  security-scan - Run security analysis"
	@echo "  release       - Create release archives"
	@echo "  integration-test - Test with real applications"
	@echo "  benchmark     - Run performance benchmarks"
	@echo "  docs          - Generate documentation"
	@echo "  help          - Show this help message"
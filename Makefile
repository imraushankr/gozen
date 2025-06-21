# GoZen Application Makefile
APP_NAME=gozen
BUILD_DIR=bin
MAIN_FILE=src/cmd/main.go
CONFIG_FILE=config.yaml

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod

# Build flags
LDFLAGS=-ldflags "-X main.Version=1.0.0 -X main.BuildTime=$(shell date -u '+%Y-%m-%d_%H:%M:%S')"

.PHONY: all build clean test coverage deps fmt lint vet run dev help

# Default target
all: clean deps test build

# Display help
help:
	@echo "Available commands:"
	@echo "  run          - Run the application in development mode"
	@echo "  dev          - Run with live reload (requires air)"
	@echo "  build        - Build the application"
	@echo "  test         - Run tests"
	@echo "  coverage     - Run tests with coverage"
	@echo "  clean        - Clean build artifacts"
	@echo "  deps         - Download and tidy dependencies"
	@echo "  fmt          - Format code"
	@echo "  lint         - Run linter (requires golangci-lint)"
	@echo "  vet          - Run go vet"
	@echo "  migrate      - Run database migrations"
	@echo "  seed         - Seed database with sample data"
	@echo "  docker-build - Build Docker image"
	@echo "  docker-run   - Run Docker container"

# Run the application
run:
	@echo "🚀 Starting GoZen application..."
	@if [ ! -f $(CONFIG_FILE) ]; then echo "❌ config.yaml not found!"; exit 1; fi
	@mkdir -p tmp
	$(GOBUILD) -o tmp/$(APP_NAME) $(MAIN_FILE) && ./tmp/$(APP_NAME)

# Development mode with live reload
dev:
	@echo "🔄 Starting GoZen in development mode with live reload..."
	@mkdir -p tmp
	@which air > /dev/null || (echo "Installing air..." && $(GOCMD) install github.com/air-verse/air@latest)
	air

# Build the application
build:
	@echo "🔨 Building GoZen..."
	@mkdir -p $(BUILD_DIR)
	$(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_FILE)
	@echo "✅ Build complete: $(BUILD_DIR)/$(APP_NAME)"

# Build for different platforms
build-linux:
	@echo "🔨 Building for Linux..."
	@mkdir -p $(BUILD_DIR)
	GOOS=linux GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(APP_NAME)-linux $(MAIN_FILE)

build-windows:
	@echo "🔨 Building for Windows..."
	@mkdir -p $(BUILD_DIR)
	GOOS=windows GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(APP_NAME)-windows.exe $(MAIN_FILE)

build-mac:
	@echo "🔨 Building for macOS..."
	@mkdir -p $(BUILD_DIR)
	GOOS=darwin GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(APP_NAME)-mac $(MAIN_FILE)

# Build for all platforms
build-all: build-linux build-windows build-mac

# Run tests
test:
	@echo "🧪 Running tests..."
	$(GOTEST) -v ./...

# Run tests with coverage
coverage:
	@echo "🧪 Running tests with coverage..."
	$(GOTEST) -v -coverprofile=coverage.out -covermode=atomic ./...
	$(GOCMD) tool cover -html=coverage.out -o coverage.html
	@echo "📊 Coverage report generated: coverage.html"

# Benchmark tests
benchmark:
	@echo "⚡ Running benchmarks..."
	$(GOTEST) -bench=. -benchmem ./...

# Clean build artifacts
clean:
	@echo "🧹 Cleaning..."
	$(GOCLEAN)
	rm -rf $(BUILD_DIR)
	rm -rf tmp
	rm -f coverage.out coverage.html
	@echo "✅ Clean complete"

# Download dependencies
deps:
	@echo "📦 Downloading dependencies..."
	$(GOMOD) download
	$(GOMOD) tidy
	@echo "✅ Dependencies updated"

# Format code
fmt:
	@echo "🎨 Formatting code..."
	$(GOCMD) fmt ./...
	@echo "✅ Code formatted"

# Vet code
vet:
	@echo "🔍 Vetting code..."
	$(GOCMD) vet ./...
	@echo "✅ Code vetted"

# Lint code
lint:
	@echo "🔍 Linting code..."
	@which golangci-lint > /dev/null || (echo "Installing golangci-lint..." && $(GOCMD) install github.com/golangci/golangci-lint/cmd/golangci-lint@latest)
	golangci-lint run
	@echo "✅ Code linted"

# Security scan
security:
	@echo "🔒 Running security scan..."
	@which gosec > /dev/null || (echo "Installing gosec..." && $(GOCMD) install github.com/securecodewarrior/gosec/v2/cmd/gosec@latest)
	gosec ./...

# Database migrations (assuming you have a migration system)
migrate:
	@echo "🗃️ Running database migrations..."
	@# Add your migration command here
	@echo "✅ Migrations complete"

# Seed database
seed:
	@echo "🌱 Seeding database..."
	@# Add your seeding command here
	@echo "✅ Database seeded"

# Install the application
install: build
	@echo "📥 Installing GoZen..."
	sudo cp $(BUILD_DIR)/$(APP_NAME) /usr/local/bin/
	@echo "✅ GoZen installed to /usr/local/bin/"

# Uninstall the application
uninstall:
	@echo "🗑️ Uninstalling GoZen..."
	sudo rm -f /usr/local/bin/$(APP_NAME)
	@echo "✅ GoZen uninstalled"

# Docker commands
docker-build:
	@echo "🐳 Building Docker image..."
	docker build -t $(APP_NAME):latest .

docker-run:
	@echo "🐳 Running Docker container..."
	docker run -p 3400:3400 -v $(PWD)/config.yaml:/app/config.yaml $(APP_NAME):latest

docker-compose-up:
	@echo "🐳 Starting services with docker-compose..."
	docker-compose up -d

docker-compose-down:
	@echo "🐳 Stopping services with docker-compose..."
	docker-compose down

# Generate documentation
docs:
	@echo "📚 Generating documentation..."
	@which godoc > /dev/null || (echo "Installing godoc..." && $(GOCMD) install golang.org/x/tools/cmd/godoc@latest)
	@echo "📚 Documentation server starting at http://localhost:6060"
	godoc -http=:6060

# Pre-commit checks
pre-commit: fmt vet lint test
	@echo "✅ Pre-commit checks passed"

# Release preparation
release: clean deps test build-all
	@echo "🚀 Release artifacts ready in $(BUILD_DIR)/"

# Setup environment (for new developers)
setup:
	@echo "🔧 Setting up development environment..."
	@mkdir -p tmp bin
	@echo "📦 Installing development tools..."
	$(GOCMD) install github.com/air-verse/air@latest
	$(GOCMD) install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@echo "✅ Environment setup complete"
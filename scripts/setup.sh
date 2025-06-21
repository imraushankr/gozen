#!/bin/bash

# scripts/setup.sh - Setup script for new developers
echo "🔧 Setting up GoZen development environment..."

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "❌ Go is not installed. Please install Go first."
    exit 1
fi

echo "✅ Go is installed: $(go version)"

# Download dependencies
echo "📦 Downloading dependencies..."
go mod download
go mod tidy

# Install development tools
echo "🛠️ Installing development tools..."

# Air for live reload
if ! command -v air &> /dev/null; then
    echo "Installing air..."
    go install github.com/cosmtrek/air@latest
fi

# golangci-lint for linting
if ! command -v golangci-lint &> /dev/null; then
    echo "Installing golangci-lint..."
    go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
fi

# gosec for security scanning
if ! command -v gosec &> /dev/null; then
    echo "Installing gosec..."
    go install github.com/securecodewarrior/gosec/v2/cmd/gosec@latest
fi

# Create necessary directories
echo "📁 Creating directories..."
mkdir -p logs
mkdir -p tmp
mkdir -p bin

# Copy example config if config.yaml doesn't exist
if [ ! -f "config.yaml" ]; then
    if [ -f "config.example.yaml" ]; then
        echo "📋 Copying example config..."
        cp config.example.yaml config.yaml
        echo "⚠️ Please update config.yaml with your settings"
    else
        echo "⚠️ config.yaml not found. Please create one based on the provided structure."
    fi
fi

# Create config.yaml from template if it doesn't exist
if [ ! -f "config.yaml" ]; then
    echo "📝 Creating config.yaml file..."
    cat > config.yaml << 'EOF'
app:
  name: "GoZen"
  version: "1.0.0"
  environment: "development" # development, staging, production
  debug: true

server:
  host: "localhost"
  port: 3400
  read_timeout: 10s
  write_timeout: 10s
  shutdown_timeout: 15s

database:
  type: sqlite # sqlite, mysql, postgresql, mongodb
  host: localhost
  port: 5432
  username: ""
  password: ""
  database: gozen.db
  ssl_mode: disable
  max_open_conns: 25
  max_idle_conns: 5
  conn_max_lifetime: 5m

jwt:
  access_token_secret: change-this-secret-in-production
  access_token_expiry: 15m
  refresh_token_secret: change-this-refresh-secret-in-production
  refresh_token_expiry: 168h
  issuer: "GoZen"

email:
  provider: smtp
  smtp:
    host: smtp.gmail.com
    port: 587
    username: your_email@gmail.com
    password: your_app_password
    from_email: your_email@gmail.com
    from_name: "GoZen App"
    use_tls: true

logger:
  level: debug        # debug, info, warn, error
  format: console     # json, console
  output: stdout      # stdout, stderr, file
  file_path: logs/app.log
  max_size: 100       # MB
  max_backups: 3
  max_age: 28         # days
  compress: true

cors:
  allow_origins: ["http://localhost:3000", "http://localhost:3001", "http://127.0.0.1:3000"]
  allow_methods: ["GET", "POST", "PUT", "DELETE", "OPTIONS"]
  allow_headers: ["*"]
  allow_credentials: true
  max_age: 300
EOF
    echo "⚠️ Please update config.yaml with your actual settings (especially JWT secrets!)"
fi

# Create .gitignore if it doesn't exist
if [ ! -f ".gitignore" ]; then
    echo "📝 Creating .gitignore..."
    cat > .gitignore << 'EOF'
# Binaries
bin/
*.exe
*.exe~
*.dll
*.so
*.dylib

# Test binary, built with `go test -c`
*.test

# Output of the go coverage tool
*.out
coverage.html

# Dependency directories
vendor/

# Go workspace file
go.work

# Environment variables
config.yaml

# IDE
.vscode/
.idea/
*.swp
*.swo

# OS
.DS_Store
Thumbs.db

# Logs
logs/
*.log

# Temporary files
tmp/
temp/

# Database files
*.db
*.sqlite
*.sqlite3

# Air
.air.toml
EOF
fi

echo "🎉 Setup complete! You can now run:"
echo "  make run       # Run the application"
echo "  make dev       # Run with live reload"
echo "  make test      # Run tests"
echo "  make build     # Build the application"

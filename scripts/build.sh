#!/bin/bash

# scripts/build.sh - Build script
echo "🔨 Building GoZen application..."

# Create bin directory if it doesn't exist
mkdir -p bin

# Build the application
go build -ldflags "-X main.Version=1.0.0 -X main.BuildTime=$(date -u '+%Y-%m-%d_%H:%M:%S')" -o bin/gozen src/cmd/main.go

if [ $? -eq 0 ]; then
    echo "✅ Build successful: bin/gozen"
else
    echo "❌ Build failed"
    exit 1
fi

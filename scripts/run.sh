#!/bin/bash

# scripts/run.sh - Simple run script
echo "🚀 Starting GoZen application..."

# Check if config file exists
if [ ! -f "config.yaml" ]; then
    echo "❌ config.yaml not found!"
    exit 1
fi

# Run the application
go run src/cmd/main.go
#!/bin/bash

# scripts/test.sh - Test script
echo "🧪 Running tests..."

# Run tests with verbose output
go test -v ./...

# Check if tests passed
if [ $? -eq 0 ]; then
    echo "✅ All tests passed"
else
    echo "❌ Some tests failed"
    exit 1
fi

#!/usr/bin/env pwsh

# scripts/run.ps1 - PowerShell script for Windows
Write-Host "🚀 Starting GoZen application..." -ForegroundColor Green

# Check if config file exists
if (-not (Test-Path "config.yaml")) {
    Write-Host "❌ config.yaml not found!" -ForegroundColor Red
    exit 1
}

# Run the application
go run src/cmd/main.go

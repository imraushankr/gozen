
#!/usr/bin/env pwsh

# scripts/build.ps1 - PowerShell build script for Windows
Write-Host "🔨 Building GoZen application..." -ForegroundColor Green

# Create bin directory if it doesn't exist
if (-not (Test-Path "bin")) {
    New-Item -ItemType Directory -Path "bin"
}

# Build the application
$buildTime = Get-Date -Format "yyyy-MM-dd_HH:mm:ss"
go build -ldflags "-X main.Version=1.0.0 -X main.BuildTime=$buildTime" -o bin/gozen.exe src/cmd/main.go

if ($LASTEXITCODE -eq 0) {
    Write-Host "✅ Build successful: bin/gozen.exe" -ForegroundColor Green
} else {
    Write-Host "❌ Build failed" -ForegroundColor Red
    exit 1
}
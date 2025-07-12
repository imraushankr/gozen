package main

import (
	"os"

	"github.com/imraushankr/gozen/src/configs"
	"github.com/imraushankr/gozen/src/internal/app"
	"github.com/imraushankr/gozen/src/internal/pkg/logger"
	"go.uber.org/zap"
)

func main() {
	// Initialize basic logger
	logger.InitBasic()
	defer logger.Sync()

	// Load configuration
	cfg, err := configs.LoadConfig()
	if err != nil {
		logger.Fatal("Failed to load config", 
			zap.Error(err),
			zap.String("hint", "Ensure app.yaml exists in configs/ directory"))
		os.Exit(1)
	}

	// Initialize full logger
	logger.Init(&cfg.Logger)
	defer logger.Sync()

	// Create and run server
	server := app.NewServer(cfg)
	server.Run()
}
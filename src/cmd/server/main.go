// package main

// import (
// 	"gozen/src/configs"
// 	"gozen/src/internal/app"
// 	"gozen/src/internal/pkg/logger"

// 	"go.uber.org/zap"
// )

// func main() {
// 	// Load configuration
// 	cfg, err := configs.LoadConfig()
// 	if err != nil {
// 		logger.Fatal("Failed to load config", zap.Error(err))
// 	}

// 	// Initialize logger
// 	logger.Init(&cfg.Logger)
// 	defer logger.Sync()

// 	// Create and run server
// 	server := app.NewServer(cfg)
// 	server.Run()
// }

// package main

// import (
// 	"gozen/src/configs"
// 	"gozen/src/internal/app"
// 	"gozen/src/internal/pkg/logger"
// 	"os"

// 	"go.uber.org/zap"
// )

// func main() {
// 	// Initialize basic logger first for startup messages
// 	logger.InitBasic()
// 	defer logger.Sync()

// 	// Load configuration
// 	cfg, err := configs.LoadConfig()
// 	if err != nil {
// 		logger.Fatal("Failed to load config", zap.Error(err))
// 		os.Exit(1)
// 	}

// 	// Reinitialize logger with full configuration
// 	logger.Init(&cfg.Logger)
// 	defer logger.Sync()

// 	// Create and run server
// 	server := app.NewServer(cfg)
// 	server.Run()
// }

// package main

// import (
// 	"gozen/src/configs"
// 	"gozen/src/internal/app"
// 	"gozen/src/internal/pkg/logger"
// 	"os"

// 	"go.uber.org/zap"
// )

// func main() {
// 	// Initialize basic logger first for startup messages
// 	logger.InitBasic()
// 	defer logger.Sync()

// 	// Load configuration
// 	cfg, err := configs.LoadConfig()
// 	if err != nil {
// 		logger.Fatal("Failed to load config", zap.Error(err))
// 		os.Exit(1)
// 	}

// 	// Reinitialize logger with full configuration
// 	logger.Init(&cfg.Logger)
// 	defer logger.Sync()

// 	// Create and run server
// 	server := app.NewServer(cfg)
// 	server.Run()
// }


package main

import (
	"gozen/src/configs"
	"gozen/src/internal/app"
	"gozen/src/internal/pkg/logger"
	"os"

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
	// if err := server.Run(); err != nil {
	// 	logger.Fatal("Server failed", zap.Error(err))
	// 	os.Exit(1)
	// }
	server.Run()
}
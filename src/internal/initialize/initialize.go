package initialize

import (
  "github.com/imraushankr/gozen/src/internal/config"
  "github.com/imraushankr/gozen/src/internal/db"
  "github.com/imraushankr/gozen/src/pkg/logger"
)

// Initialize sets up all applications dependencies
func Initialize(cfg *config.Config) error {
  // Initialize logger
  if err := logger.InitLogger(&cfg.Logger); err != nil {
    return err
  }
  
  // Initialize database
  if err := db.InitDatabase(cfg.Database); err != nil {
    return err
  }
  
  // Run migrations for SQL databases
  if cfg.Database.Type != "mongodb" {
    if err := db.AutoMigrate(); err != nil {
      return err
    }
  }
  
  return nil
}
package db

import (
  "fmt"
  "github.com/imraushankr/gozen/src/internal/config"
  "gorm.io/driver/sqlite"
  "gorm.io/gorm"
  "gorm.io/gorm/logger"
)

func initSQLite(cfg config.DatabaseConfig) error {
  logLevel := logger.Info
  
  db, err := gorm.Open(sqlite.Open(cfg.Database), &gorm.Config{
    Logger: logger.Default.LogMode(logLevel),
  })
  if err != nil {
    return err
  }
  
  sqlDB, err := db.DB()
  if err != nil {
    return err
  }
  
  // Set connection pool settings
  sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
  sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
  sqlDB.SetConnMaxLifetime(cfg.ConnMaxLifetime)
  
  SQLInstance = db
  fmt.Println("✅ Connected to SQLite")
  return nil
}
package db

import (
  "fmt"
  "github.com/imraushankr/gozen/src/internal/config"
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
  "gorm.io/gorm/logger"
)

func initPostgreSQL(cfg config.DatabaseConfig) error {
  logLevel := logger.Info // Default for now
  
  // Build DSN manually since cfg is DatabaseConfig, not Config
  dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
    cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.Database, cfg.SSLMode)
  
  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
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
  fmt.Println("✅ Connected to PostgreSQL")
  return nil
}
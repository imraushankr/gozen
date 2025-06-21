package db

import (
  "fmt"
  "github.com/imraushankr/gozen/src/internal/config"
  "gorm.io/driver/mysql"
  "gorm.io/gorm"
  "gorm.io/gorm/logger"
)

func initMySQL(cfg config.DatabaseConfig) error {
  logLevel := logger.Info // Default for now
  
  // Build DSN manually since cfg is DatabaseConfig, not Config
  dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
    cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
  
  db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
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
  fmt.Println("✅ Connected to MySQL")
  return nil
}
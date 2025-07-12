package database

import (
	"database/sql"
	"fmt"

	"github.com/imraushankr/gozen/src/configs"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func newPostgresDB(cfg *configs.DatabaseConfig) (*DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.Database, cfg.SSLMode,
	)

	sqlDB, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return &DB{SQL: sqlDB, GORM: gormDB}, nil
}
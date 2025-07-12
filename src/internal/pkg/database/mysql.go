package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gozen/src/configs"
)

func newMySQLDB(cfg *configs.DatabaseConfig) (*DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database,
	)

	sqlDB, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return &DB{SQL: sqlDB, GORM: gormDB}, nil
}
package database

import (
	"context"
	"database/sql"
	"time"
)

func (db *DB) Close() error {
	sqlDB, err := db.GORM.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func (db *DB) HealthCheck() error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	return db.SQL.PingContext(ctx)
}

func (db *DB) IsAlive() bool {
	return db.HealthCheck() == nil
}

func (db *DB) Stats() sql.DBStats {
	return db.SQL.Stats()
}
package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	driverName = "postgres"
)

func Open(cfg Config) (*sql.DB, error) {
	db, err := sql.Open(driverName, cfg.OpenConfig.FormatDSN())
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}

	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetConnMaxLifetime(cfg.ConnMaxLifetime)
	db.SetConnMaxIdleTime(cfg.ConnMaxIdleTime)

	return db, nil
}

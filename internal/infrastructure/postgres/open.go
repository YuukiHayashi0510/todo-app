package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	defaultConnectTimeout = time.Second * 5
)

func Open(cfg Config) (*DB, error) {
	// pgx設定 参考: https://github.com/jackc/pgx/discussions/1989
	config, err := pgxpool.ParseConfig(cfg.OpenConfig.FormatDSN())
	if err != nil {
		return nil, err
	}

	config.MaxConns = int32(cfg.MaxOpenConns)
	config.MinConns = int32(cfg.MaxIdleConns)
	config.MaxConnLifetime = cfg.ConnMaxLifetime
	config.MaxConnIdleTime = cfg.ConnMaxIdleTime
	config.ConnConfig.ConnectTimeout = defaultConnectTimeout

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}

	return pool, nil
}

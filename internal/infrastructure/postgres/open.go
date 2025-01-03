package postgres

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/tracelog"
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
	if cfg.Trace {
		config.ConnConfig.Tracer = &tracelog.TraceLog{
			Logger: tracelog.LoggerFunc(func(ctx context.Context, level tracelog.LogLevel, msg string, data map[string]interface{}) {
				if sql, ok := data["sql"].(string); ok {
					if duration, ok := data["time"].(time.Duration); ok {
						sql = strings.Join(strings.Fields(sql), " ")
						log.Printf("SQL: %s [%.2fms]", sql, float64(duration.Microseconds())/1000)
					}
				}
			}),
			LogLevel: tracelog.LogLevelDebug,
		}
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}

	return pool, nil
}

package postgres

import (
	"fmt"
	"time"
)

type OpenConfig struct {
	Host     string
	Port     int
	DBName   string
	User     string
	Password string
	SslMode  string
	Trace    bool
}

func (cfg OpenConfig) FormatDSN() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SslMode,
	)
}

type ConnectionConfig struct {
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
	ConnMaxIdleTime time.Duration
}

type Config struct {
	OpenConfig
	ConnectionConfig
}

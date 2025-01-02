package postgres

import (
	"time"
)

type OpenConfig struct {
	Host     string
	Port     int
	DBName   string
	User     string
	Password string
	SslMode  string
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

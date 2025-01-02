package config

import (
	"embed"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/YuukiHayashi0510/todo-app/internal/infrastructure/secrets"
	"github.com/YuukiHayashi0510/todo-app/pkg/empty"
	"gopkg.in/yaml.v3"
)

const (
	configFileName = "config.yml"
)

//go:embed config.yml
var configFile embed.FS

var AppConfig Config

type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
	Logging  LoggingConfig  `yaml:"logging"`
}

type ServerConfig struct {
	Port int `yaml:"port"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DBName   string `yaml:"db_name"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	SslMode  string `yaml:"ssl_mode"`

	SecretConfig SecretConfig `yaml:"secret"`

	MaxOpenConnections int           `yaml:"max_open_connections"`
	MaxIdleConnections int           `yaml:"max_idle_connections"`
	ConnMaxLifetime    time.Duration `yaml:"conn_max_lifetime"`
	ConnMaxIdleTime    time.Duration `yaml:"conn_max_idle_time"`
}

type SecretConfig struct {
	Region     string `yaml:"region"`
	SecretName string `yaml:"secret_name"`
}

type LoggingConfig struct {
	Level  string `yaml:"level"`
	Format string `yaml:"format"`
	Path   string `yaml:"path"`
}

type DatabaseSecret struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
}

func init() {
	if err := initConfig(); err != nil {
		log.Fatal(err)
	}
}

func initConfig() error {
	// yamlファイルの読み込み
	data, err := configFile.ReadFile(configFileName)
	if err != nil {
		return fmt.Errorf("設定ファイルの読み込みに失敗: %w", err)
	}
	if err := yaml.Unmarshal(data, &AppConfig); err != nil {
		return fmt.Errorf("設定のパースに失敗: %w", err)
	}

	// SecretsManagerの設定値がある場合、上書き
	if !empty.Is(AppConfig.Database.SecretConfig.SecretName) &&
		!empty.Is(AppConfig.Database.SecretConfig.Region) {
		overrideDBConfig()
	}

	return nil
}

// Secretsの値でDBの設定値を上書きする
func overrideDBConfig() {
	secretValue, err := secrets.GetSecrets(
		AppConfig.Database.SecretConfig.Region,
		AppConfig.Database.SecretConfig.SecretName,
	)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to get secrets: %w", err))
	}

	var dbSecret DatabaseSecret
	if err := json.Unmarshal([]byte(secretValue), &dbSecret); err != nil {
		log.Fatal(fmt.Errorf("failed to unmarshal secrets value: %w", err))
	}

	// DB設定の上書き
	AppConfig.Database.Host = dbSecret.Host
	AppConfig.Database.User = dbSecret.User
	AppConfig.Database.Password = dbSecret.Password
}

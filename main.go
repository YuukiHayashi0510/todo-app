package main

import (
	"fmt"
	"log"
	"os"

	"github.com/YuukiHayashi0510/todo-app/config"
	"github.com/YuukiHayashi0510/todo-app/internal/infrastructure/postgres"
	"github.com/YuukiHayashi0510/todo-app/internal/logger"
	"github.com/YuukiHayashi0510/todo-app/internal/server"
	"github.com/YuukiHayashi0510/todo-app/internal/web"
	"github.com/YuukiHayashi0510/todo-app/internal/web/handler"
	"github.com/gin-gonic/gin"
)

const (
	logBaseDirKey = "LOG_BASE_DIR"
)

func main() {
	if err := run(fmt.Sprintf(":%d", config.AppConfig.Server.Port)); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}

func run(addr string) (runErr error) {
	err := logger.Init(config.AppConfig.Logging, os.Getenv(logBaseDirKey))
	if err != nil {
		return fmt.Errorf("failed to sync logger: %w", err)
	}

	db, err := postgres.Open(postgres.Config{
		OpenConfig: postgres.OpenConfig{
			Host:     config.AppConfig.Database.Host,
			Port:     config.AppConfig.Database.Port,
			DBName:   config.AppConfig.Database.DBName,
			User:     config.AppConfig.Database.User,
			Password: config.AppConfig.Database.Password,
			SslMode:  config.AppConfig.Database.SslMode,
			Trace:    config.AppConfig.Database.Trace,
		},
		ConnectionConfig: postgres.ConnectionConfig{
			MaxIdleConns:    config.AppConfig.Database.MaxIdleConnections,
			MaxOpenConns:    config.AppConfig.Database.MaxOpenConnections,
			ConnMaxLifetime: config.AppConfig.Database.ConnMaxLifetime,
			ConnMaxIdleTime: config.AppConfig.Database.ConnMaxIdleTime,
		},
	})
	if err != nil {
		return fmt.Errorf("failed to get db: %w", err)
	}
	// サーバ終了時にDB接続を閉じる
	defer db.Close()

	// Ginのインスタンス初期化
	gin.DisableConsoleColor()
	gin.SetMode(config.AppConfig.Server.Mode)
	r := gin.New()

	// panic recovery
	r.Use(gin.Recovery())

	// ルーティング
	router := web.NewRouter(
		web.Handlers{
			Organizations: handler.NewOrganizationHandler(db),
		},
	)
	router.Routing(r.Group(""))

	if err := server.Run(r, addr); err != nil {
		return fmt.Errorf("failed to run server: %w", err)
	}

	return
}

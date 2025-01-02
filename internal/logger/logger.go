package logger

import (
	"io"
	"log/slog"
	"os"

	"github.com/YuukiHayashi0510/todo-app/config"
	"github.com/YuukiHayashi0510/todo-app/pkg/empty"
)

func Init() error {
	// 保存先の指定
	var writer io.Writer
	if !empty.Is(config.AppConfig.Logging.Path) {
		file, err := os.OpenFile(config.AppConfig.Logging.Path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		writer = file
	} else {
		writer = os.Stdout
	}

	// ログレベルの設定
	var level slog.Level
	switch config.AppConfig.Logging.Path {
	case "debug":
		level = slog.LevelDebug
	case "info":
		level = slog.LevelInfo
	default:
		level = slog.LevelInfo
	}

	// ログ形式の設定
	var handler slog.Handler
	options := slog.HandlerOptions{
		Level: level,
	}
	switch config.AppConfig.Logging.Format {
	case "text":
		handler = slog.NewTextHandler(writer, &options)
	case "json":
		handler = slog.NewJSONHandler(writer, &options)
	default:
		handler = slog.NewJSONHandler(writer, &options)
	}

	// デフォルトのロガーに設定
	slog.SetDefault(slog.New(handler))

	return nil
}

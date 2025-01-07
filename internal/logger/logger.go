package logger

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/YuukiHayashi0510/todo-app/config"
	"github.com/YuukiHayashi0510/todo-app/pkg/empty"
)

func Init(cfg config.LoggingConfig, logBaseDir string) error {
	// 保存先の指定
	var writer io.Writer
	if !empty.Is(cfg.Path) && !empty.Is(logBaseDir) {
		logFilePath := filepath.Join(logBaseDir, cfg.Path)

		// ディレクトリを作成（既に存在する場合は何もしない）
		if err := os.MkdirAll(filepath.Dir(logFilePath), 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", logBaseDir, err)
		}

		file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		writer = file
	} else {
		writer = os.Stdout
	}

	// ログレベルの設定
	var level slog.Level
	switch cfg.Path {
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
	switch cfg.Format {
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

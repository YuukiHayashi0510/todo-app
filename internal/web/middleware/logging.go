package middleware

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/YuukiHayashi0510/todo-app/internal/web/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequestLogger(lgr *slog.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()

		// リクエストIDの付与
		baseLogger := lgr.With(
			slog.String("request_id", uuid.New().String()),
		)

		// アクセスログ
		requestLogger := baseLogger.With(
			slog.String("proto", ctx.Request.Proto),
			slog.String("host", ctx.Request.Host),
			slog.String("method", ctx.Request.Method),
			slog.String("uri", ctx.Request.RequestURI),
			slog.String("raw_query", ctx.Request.URL.RawQuery),
			slog.Any("structured_params", ctx.Request.URL.Query()),
			slog.String("remote_addr", ctx.Request.RemoteAddr),
			slog.String("client_ip", ctx.ClientIP()),
			slog.String("user_agent", ctx.Request.UserAgent()),
		)
		requestLogger.Info("request started")

		// アプリケーションログ
		appLogger := baseLogger.With(
			slog.String("handler_name", ctx.HandlerName()),
		)

		ctx.Set(LoggerContextKey, appLogger)
		ctx.Next()

		res := ctx.MustGet(ResponseContextKey).(*response.Response)

		// レスポンス情報の付与
		requestLogger = requestLogger.With(
			slog.Int("status", res.HttpStatus),
			slog.Duration("duration", time.Since(start)),
		)

		// ステータスコードでのハンドリング
		switch {
		case res.HttpStatus < http.StatusBadRequest:
			requestLogger.Info("request complete success")
		// 4xx
		case res.HttpStatus >= http.StatusBadRequest && res.HttpStatus < http.StatusInternalServerError:
			requestLogger.Info("request complete client_error")
		// 5xx
		case res.HttpStatus >= http.StatusInternalServerError:
			requestLogger.Error("request complete server_error", slog.String("error", res.Data.(response.ServerError).Parent.Error()))
		}
	}
}

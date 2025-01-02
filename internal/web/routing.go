package web

import (
	"log/slog"
	"net/http"

	"github.com/YuukiHayashi0510/todo-app/internal/web/middleware"
	"github.com/gin-gonic/gin"
)

func Routing(group *gin.RouterGroup) {
	// サンプルエンドポイント, ヘルスチェックに使用
	group.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	group.Use(middleware.RequestLogger(slog.Default()), middleware.CreateResponse())
}

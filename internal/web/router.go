package web

import (
	"log/slog"
	"net/http"

	"github.com/YuukiHayashi0510/todo-app/internal/web/middleware"
	"github.com/YuukiHayashi0510/todo-app/internal/web/request"
	"github.com/gin-gonic/gin"
)

type Router struct {
	Handlers Handlers
}

func NewRouter(handlers Handlers) *Router {
	return &Router{handlers}
}

func (r *Router) Routing(group *gin.RouterGroup) {
	// サンプルエンドポイント, ヘルスチェックに使用
	group.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	group.Use(middleware.RequestLogger(slog.Default()), middleware.CreateResponse())

	r.routingOrgs(group)
}

// TODO: ハンドラの設定
func (r *Router) routingOrgs(group *gin.RouterGroup) {
	group = group.Group("/organizations")
	group.Use(middleware.Validate[request.OrganizationRequest]())

	group.GET("", r.Handlers.Organizations.List)
}

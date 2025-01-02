package middleware

import (
	"github.com/YuukiHayashi0510/todo-app/internal/web/response"
	"github.com/gin-gonic/gin"
)

func CreateResponse() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		res := c.MustGet(ResponseContextKey).(*response.Response)
		if res.RedirectPath != nil {
			c.Redirect(res.HttpStatus, *res.RedirectPath)
		} else if res.TemplatePath != nil {
			c.File(*res.TemplatePath)
		} else {
			c.JSON(res.HttpStatus, res.Data)
		}
	}
}

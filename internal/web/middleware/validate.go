package middleware

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/YuukiHayashi0510/todo-app/internal/web/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Validate[T any]() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req T

		if err := bindRequest(&req, c); err != nil {
			var (
				validationErrors validator.ValidationErrors
				jsonSyntaxError  *json.SyntaxError
			)

			if errors.Is(err, io.EOF) {
				c.Set(ResponseContextKey, &response.Response{
					Data: response.ServerError{
						Parent:  err,
						Message: "missing required params",
					},
				})
			} else if errors.As(err, &validationErrors) {
				c.Set(ResponseContextKey, &response.Response{
					HttpStatus: http.StatusBadRequest,
					Data: response.ServerError{
						Parent:  err,
						Message: "validation error",
					},
				})
			} else if errors.As(err, &jsonSyntaxError) {
				c.Set(ResponseContextKey, &response.Response{
					Data: response.ServerError{
						Parent:  err,
						Message: "invalid JSON format",
					},
				})
			} else {
				c.Set(ResponseContextKey, &response.Response{
					Data: response.ServerError{
						Parent:  err,
						Message: "server error",
					},
				})
			}
			c.Abort()
		}

		c.Set(ValidationContextKey, &req)
		c.Next()
	}
}

func bindRequest[T any](req *T, c *gin.Context) error {
	switch c.Request.Method {
	case http.MethodGet:
		return c.ShouldBindQuery(req)
	case http.MethodPost, http.MethodPut, http.MethodPatch:
		return c.ShouldBindJSON(req)
	}

	// URIとHTTPメソッドはGinのルーティングが判定するため
	// HTTPメソッドのサポートがない場合の考慮は不要
	return nil
}
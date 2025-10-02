package gin

import (
	"context"

	"github.com/gin-gonic/gin"
)

type ginContextKey struct{}

func GinContext() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), ginContextKey{}, c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

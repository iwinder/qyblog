package middleware

import (
	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
	"github.com/iwinder/qyblog/internal/pkg/core"
	"github.com/iwinder/qyblog/internal/pkg/errors"
	"net/http"
	"time"
)

func TimeoutMiddleware(time time.Duration) gin.HandlerFunc {
	return timeout.New(
		timeout.WithTimeout(time),
		timeout.WithHandler(func(c *gin.Context) {
			c.Next()
		}),
		timeout.WithResponse(func(c *gin.Context) {
			core.WriteResponse(c, errors.WithCode(http.StatusGatewayTimeout, "timeout"), nil)
		}),
	)
}

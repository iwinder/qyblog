package qycms_middleware

import (
	log "gitee.com/windcoder/qingyucms/internal/pkg/qygo-log"
	"github.com/gin-gonic/gin"
)

func Context() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(log.KeyRequestID, c.GetString(KeyXRequestId))
		c.Set(log.KeyUsername, c.GetString(KeyUsername))
		c.Next()
	}
}

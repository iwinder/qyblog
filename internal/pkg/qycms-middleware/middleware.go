package qycms_middleware

import (
	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
	"net/http"
	"time"
)

var Middlewares = defaultMiddlewares()

func NoCache(c *gin.Context) {
	c.Header(KeyCacheControl, "no-cache, no-store, max-age=0, must-revalidate, value")
	c.Header(KeyExpires, "Thu, 01 Jan 1970 00:00:00 GMT")
	c.Header(KeyLastModified, time.Now().UTC().Format(http.TimeFormat))
	c.Next()
}

func Options(c *gin.Context) {
	if c.Request.Method != "OPTIONS" {
		c.Next()
	} else {
		c.Header(KeyAccessControlAllowOrigin, "*")
		c.Header(KeyAccessControlAllowMethods, "GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Header(KeyAccessControlAllowHeadres, "authorization, origin, content-type, accept")
		c.Header(KeyAllow, "HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Header(KeyContentType, "application/json")
		c.AbortWithStatus(http.StatusOK)
	}
}

func Secure(c *gin.Context) {
	c.Header(KeyAccessControlAllowOrigin, "*")
	c.Header(KeyXFrameOptions, "DENY")
	c.Header(KeyXContentTypeOptions, "nosniff")
	c.Header(KeyXXSSProtection, "1; mode=block")

	if c.Request.TLS != nil {
		c.Header(KeyStrictTransportSecurity, "max-age=31536000")
	}
}

func defaultMiddlewares() map[string]gin.HandlerFunc {
	return map[string]gin.HandlerFunc{
		"recovery":  gin.Recovery(),
		"secure":    Secure,
		"options":   Options,
		"nocache":   NoCache,
		"cors":      Cors(),
		"requestID": RequestID(),
		"logger":    Logger(),
		"dump":      gindump.Dump(), // Gin中间件/处理程序，用于转储请求和响应的标头/正文。对调试应用程序非常有帮助。
	}
}

package qy_middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"io"
	"time"
)

func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		rid := c.GetHeader(KeyXRequestId)

		if rid == "" {
			rid = uuid.Must(uuid.NewV4(), nil).String()
			c.Request.Header.Set(KeyXRequestId, rid)
			c.Set(KeyXRequestId, rid)
		}

		c.Writer.Header().Set(KeyXRequestId, rid)
		c.Next()
	}
}

func GetLoggerConfig(formatter gin.LogFormatter, output io.Writer, skipPaths []string) gin.LoggerConfig {
	if formatter == nil {
		formatter = GetDefaultLogFormatterWithRequestID()
	}

	return gin.LoggerConfig{
		Formatter: formatter,
		Output:    output,
		SkipPaths: skipPaths,
	}
}

func GetDefaultLogFormatterWithRequestID() gin.LogFormatter {
	return func(params gin.LogFormatterParams) string {
		var statusColor, methodColor, resetColor string
		if params.IsOutputColor() {
			statusColor = params.StatusCodeColor()
			methodColor = params.MethodColor()
			resetColor = params.ResetColor()
		}

		if params.Latency > time.Minute {
			params.Latency -= params.Latency % time.Second
		}
		return fmt.Sprintf("%s%3d%s - [%s] \"%v %s%s%s %s\" %s",
			statusColor, params.StatusCode,
			resetColor, params.ClientIP, params.Latency,
			methodColor, params.Method, resetColor, params.Path,
			params.ErrorMessage,
		)
	}
}

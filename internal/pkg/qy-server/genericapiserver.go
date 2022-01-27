package qy_server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type GenericAPIServer struct {
	*gin.Engine
	mode               string
	middlewares        []string
	InsecureSeringInfo *InsecureSeringInfo
	ShutdownTimeout    time.Duration
	healthz            bool
	enableMetrics      bool
	enableProfiling    bool
	insecureServer     *http.Server
}

func (s *GenericAPIServer) Setup() {
	gin.SetMode(s.mode)
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		//log.Inf
	}
}

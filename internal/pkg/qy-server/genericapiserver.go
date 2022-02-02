package qy_server

import (
	log "gitee.com/windcoder/qingyucms/internal/pkg/qy-log"
	middleware "gitee.com/windcoder/qingyucms/internal/pkg/qy-middleware"
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
		log.Infof("%-6s %-s --> %s (%d handlers)", httpMethod, absolutePath, handlerName, nuHandlers)
	}
}

func (s *GenericAPIServer) InstallMiddlewares() {
	s.Use(middleware.RequestID())
	s.Use(middleware.Context())

	for _, m := range s.middlewares {
		//mv, ok := middleware.
	}
}

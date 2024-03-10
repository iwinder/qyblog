package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/iwinder/qyblog/internal/pkg/core"
	log "github.com/iwinder/qyblog/internal/pkg/logger"
	"github.com/iwinder/qyblog/internal/pkg/middleware"
	"github.com/iwinder/qyblog/internal/pkg/version"
	ginprometheus "github.com/zsais/go-gin-prometheus"
	"golang.org/x/sync/errgroup"
	"net/http"
	"strings"
	"time"
)

type HttpServer struct {
	middlewares []string
	// SecureServingInfo holds configuration of the TLS server.
	httpServingInfo *HttpServingInfo

	// ShutdownTimeout is the timeout used for server shutdown. This specifies the timeout before server
	// gracefully shutdown returns.
	ShutdownTimeout time.Duration

	*gin.Engine
	healthz         bool
	enableMetrics   bool
	enableProfiling bool
	// wrapper for gin.Engine

	httpServer *http.Server
}

func initHttpServer(s *HttpServer) {
	// do some setup
	s.Setup()
	s.InstallMiddlewares()
	s.InstallAPIs()
}
func (s *HttpServer) Setup() {
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Infof("%-6s %-s --> %s (%d handlers)", httpMethod, absolutePath, handlerName, nuHandlers)
	}
}

func (s *HttpServer) InstallMiddlewares() {
	s.Use(middleware.RequestID())
	s.Use(middleware.TimeoutMiddleware(s.httpServingInfo.Timeout))
	for _, m := range s.middlewares {
		mw, ok := middleware.Middlewares[m]
		if !ok {
			log.Warnf("can not find middleware: %s", m)

			continue
		}

		log.Infof("install middleware: %s", m)
		s.Use(mw)
	}
}

func (s *HttpServer) InstallAPIs() {
	// install healthz handler
	if s.healthz {
		s.GET("/healthz", func(c *gin.Context) {
			core.WriteResponse(c, nil, map[string]string{"status": "ok"})
		})
	}

	// install metric handler
	if s.enableMetrics {
		prometheus := ginprometheus.NewPrometheus("gin")
		prometheus.Use(s.Engine)
	}

	// install pprof handler
	if s.enableProfiling {
		pprof.Register(s.Engine)
	}

	s.GET("/version", func(c *gin.Context) {
		core.WriteResponse(c, nil, version.Get())
	})
}
func (s *HttpServer) Run() error {
	s.httpServer = &http.Server{
		Addr:    s.httpServingInfo.Address,
		Handler: s,
	}
	var eg errgroup.Group

	eg.Go(func() error {
		log.Infof("Start to listening the incoming requests on http address: %s", s.httpServingInfo.Address)

		if err := s.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err.Error())

			return err
		}

		log.Infof("Server on %s stopped", s.httpServingInfo.Address)

		return nil
	})

	// Ping the server to make sure the router is working.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if s.healthz {
		if err := s.ping(ctx); err != nil {
			return err
		}
	}

	if err := eg.Wait(); err != nil {
		log.Fatal(err.Error())
	}

	return nil
}

func (s *HttpServer) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := s.httpServer.Shutdown(ctx); err != nil {
		log.Warnf("Shutdown HttpOptions server failed: %s", err.Error())
	}
	log.Infof("HttpOptions Server exiting")
}

func (s *HttpServer) ping(ctx context.Context) error {
	url := fmt.Sprintf("http://%s/healthz", s.httpServingInfo.Address)
	if strings.Contains(s.httpServingInfo.Address, "0.0.0.0") {
		url = fmt.Sprintf("http://127.0.0.1:%s/healthz", strings.Split(s.httpServingInfo.Address, ":")[1])
	}

	for {
		// Change NewRequest to NewRequestWithContext and pass context it
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err != nil {
			return err
		}
		// Ping the server by sending a GET request to `/healthz`.

		resp, err := http.DefaultClient.Do(req)
		if err == nil && resp.StatusCode == http.StatusOK {
			log.Info("The router has been deployed successfully.")

			resp.Body.Close()

			return nil
		}

		// Sleep for a second to continue the next ping.
		log.Info("Waiting for the router, retry in 1 second.")
		time.Sleep(1 * time.Second)

		select {
		case <-ctx.Done():
			log.Fatal("can not ping http server within the specified time interval.")
		default:
		}
	}
	// return fmt.Errorf("the router has no response, or it might took too long to start up")
}

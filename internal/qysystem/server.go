package qysystem

import (
	"fmt"
	"gitee.com/windcoder/qingyucms/internal/pkg/qy-common/shutdown"
	posixsignal "gitee.com/windcoder/qingyucms/internal/pkg/qy-common/shutdown/shutdownmanagers"
	log "gitee.com/windcoder/qingyucms/internal/pkg/qy-log"
	genericoptions "gitee.com/windcoder/qingyucms/internal/pkg/qy-options"
	genericapiserver "gitee.com/windcoder/qingyucms/internal/pkg/qy-server"
	"gitee.com/windcoder/qingyucms/internal/qysystem/config"
	"gitee.com/windcoder/qingyucms/internal/qysystem/store"
	"gitee.com/windcoder/qingyucms/internal/qysystem/store/mysql"
)

type apiServer struct {
	gs               *shutdown.GracefulShutdown
	genericAPIServer *genericapiserver.GenericAPIServer
}

type preparedAPIServer struct {
	*apiServer
}

type ExtraConfig struct {
	Addr         string
	MaxMsgSize   int
	mysqlOptions *genericoptions.MySQLOptions
}
type completedExtraConfig struct {
	*ExtraConfig
}

func (c *ExtraConfig) complete() *completedExtraConfig {
	if c.Addr == "" {
		c.Addr = "127.0.0.1:8081"
	}
	return &completedExtraConfig{c}
}
func createAPIServer(cfg *config.Config) (*apiServer, error) {
	gs := shutdown.New()
	gs.AddShutdownManagee(posixsignal.NewPosixSignalManager())

	genericConfig, err := buildGenericConfig(cfg)

	if err != nil {
		return nil, err
	}

	extraConfig, err := buildExtraConfig(cfg)
	if err != nil {
		return nil, err
	}

	genericServer, err := genericConfig.Complete().New()
	if err != nil {
		return nil, err
	}

	extraConfig.complete().New()
	//extraServer, err := extraConfig.complete().New()

	server := &apiServer{
		gs:               gs,
		genericAPIServer: genericServer,
	}
	return server, nil
}

func (s *apiServer) PrepareRun() preparedAPIServer {
	initRouter(s.genericAPIServer.Engine)
	s.gs.AddShutdownCallback(shutdown.ShutdownFunc(func(string) error {
		mysqlStore, _ := mysql.GetMySQLFactoryOr(nil)
		if mysqlStore != nil {
			return mysqlStore.Close()
		}
		s.genericAPIServer.Close()

		return nil
	}))
	return preparedAPIServer{s}
}
func buildGenericConfig(cfg *config.Config) (genericConfig *genericapiserver.Config, lastErr error) {
	genericConfig = genericapiserver.NewConfig()

	if lastErr = cfg.GenericServerRunOptions.ApplyTo(genericConfig); lastErr != nil {
		return
	}
	if lastErr = cfg.InsecureServing.ApplyTo(genericConfig); lastErr != nil {
		return
	}

	return
}

func buildExtraConfig(cfg *config.Config) (*ExtraConfig, error) {
	return &ExtraConfig{
		Addr:         fmt.Sprintf("%s:%d", cfg.InsecureServing.BindAddress, cfg.InsecureServing.BindPort),
		MaxMsgSize:   0,
		mysqlOptions: cfg.MySQLOptions,
	}, nil
}

func (s preparedAPIServer) Run() error {
	//go s.genericAPIServer.Run()
	if err := s.gs.Start(); err != nil {
		log.Fatalf("start shutdown manager failed: %s", err.Error())
	}
	return s.genericAPIServer.Run()
}

func (c *completedExtraConfig) New() {
	storeIns, _ := mysql.GetMySQLFactoryOr(c.mysqlOptions)
	store.SetClient(storeIns)
}

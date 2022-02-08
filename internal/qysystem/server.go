package qysystem

import (
	"context"
	"fmt"
	"gitee.com/windcoder/qingyucms/internal/pkg/qy-common/shutdown"
	posixsignal "gitee.com/windcoder/qingyucms/internal/pkg/qy-common/shutdown/shutdownmanagers"
	log "gitee.com/windcoder/qingyucms/internal/pkg/qy-log"
	genericoptions "gitee.com/windcoder/qingyucms/internal/pkg/qy-options"
	genericapiserver "gitee.com/windcoder/qingyucms/internal/pkg/qy-server"
	storage "gitee.com/windcoder/qingyucms/internal/pkg/qy-storage"
	"gitee.com/windcoder/qingyucms/internal/qysystem/config"
	"gitee.com/windcoder/qingyucms/internal/qysystem/store"
	"gitee.com/windcoder/qingyucms/internal/qysystem/store/mysql"
)

type apiServer struct {
	gs               *shutdown.GracefulShutdown
	redisOptions     *genericoptions.RedisOptions
	genericAPIServer *genericapiserver.GenericAPIServer
}

type preparedAPIServer struct {
	*apiServer
}

type ExtraConfig struct {
	Addr         string
	MaxMsgSize   int
	mysqlOptions *genericoptions.MySQLOptions
	qyOptions    *genericoptions.QyOptions
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

	erra := extraConfig.complete().New()
	if erra != nil {
		return nil, erra
	}

	genericServer, err := genericConfig.Complete().New()
	if err != nil {
		return nil, err
	}

	//extraServer, err := extraConfig.complete().New()

	server := &apiServer{
		gs:               gs,
		redisOptions:     cfg.RedisOptions,
		genericAPIServer: genericServer,
	}
	return server, nil
}

func (s *apiServer) PrepareRun() preparedAPIServer {
	initRouter(s.genericAPIServer.Engine)
	s.initRedisStore()
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
		qyOptions:    cfg.QyOptions,
	}, nil
}

func (s preparedAPIServer) Run() error {
	//go s.genericAPIServer.Run()
	if err := s.gs.Start(); err != nil {
		log.Fatalf("start shutdown manager failed: %s", err.Error())
	}
	return s.genericAPIServer.Run()
}

func (c *completedExtraConfig) New() error {
	storeIns, err := mysql.GetMySQLFactoryOr(c.mysqlOptions)
	if err != nil {
		return err
	}
	store.SetClient(storeIns)
	config.GetQyComConfigOr(c.qyOptions)
	return nil
}

func (s *apiServer) initRedisStore() {
	ctx, cancel := context.WithCancel(context.Background())
	s.gs.AddShutdownCallback(shutdown.ShutdownFunc(func(string) error {
		cancel()
		return nil
	}))

	config := &storage.Config{
		Host:                  s.redisOptions.Host,
		Port:                  s.redisOptions.Port,
		Addrs:                 s.redisOptions.Addrs,
		Username:              s.redisOptions.Username,
		Password:              s.redisOptions.Password,
		Database:              s.redisOptions.Database,
		MasterName:            s.redisOptions.MasterName,
		MaxIdle:               s.redisOptions.MaxIdle,
		MaxActive:             s.redisOptions.MaxActive,
		Timeout:               s.redisOptions.Timeout,
		EnableCluster:         s.redisOptions.EnableCluster,
		UseSSL:                s.redisOptions.UseSSL,
		SSLInsecureSkipVerify: s.redisOptions.SSLInsecureSkipVerify,
	}
	go storage.ConnectToRedis(ctx, config)
}

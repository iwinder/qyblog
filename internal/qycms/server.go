package qycms

import (
	"context"
	"fmt"
	genericoptions "gitee.com/windcoder/qingyucms/internal/pkg/qycms-options"
	genericapiserver "gitee.com/windcoder/qingyucms/internal/pkg/qycms-server"
	"gitee.com/windcoder/qingyucms/internal/pkg/qygo-common/shutdown"
	posixsignal "gitee.com/windcoder/qingyucms/internal/pkg/qygo-common/shutdown/shutdownmanagers"
	log "gitee.com/windcoder/qingyucms/internal/pkg/qygo-log"
	storage "gitee.com/windcoder/qingyucms/internal/pkg/qygo-storage"
	"gitee.com/windcoder/qingyucms/internal/qycms/config"
	cacheV1 "gitee.com/windcoder/qingyucms/internal/qycms/controller/v1/cache"
	pb "gitee.com/windcoder/qingyucms/internal/qycms/proto/v1"
	"gitee.com/windcoder/qingyucms/internal/qycms/store"
	"gitee.com/windcoder/qingyucms/internal/qycms/store/mysql"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type apiServer struct {
	gs               *shutdown.GracefulShutdown
	redisOptions     *genericoptions.RedisOptions
	grpcAPIServer    *grpcAPIServer
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

	extraServer, err := extraConfig.complete().New()
	if err != nil {
		return nil, err
	}

	genericServer, err := genericConfig.Complete().New()
	if err != nil {
		return nil, err
	}

	server := &apiServer{
		gs:               gs,
		redisOptions:     cfg.RedisOptions,
		genericAPIServer: genericServer,
		grpcAPIServer:    extraServer,
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
	if err := s.gs.Start(); err != nil {
		log.Fatalf("start shutdown manager failed: %s", err.Error())
	}
	return s.genericAPIServer.Run()
}

func (c *completedExtraConfig) New() (*grpcAPIServer, error) {
	opts := []grpc.ServerOption{
		grpc.MaxRecvMsgSize(c.MaxMsgSize),
		//grpc.Creds(c),
	}
	grpcServer := grpc.NewServer(opts...)

	storeIns, err := mysql.GetMySQLFactoryOr(c.mysqlOptions)
	if err != nil {
		log.Fatalf("Failed to get cache instance: %s", err.Error())
	}
	aerr := storeIns.InitTables()
	if aerr != nil {
		log.Fatalf("Failed to Init Tables: %s", err.Error())
	}
	store.SetClient(storeIns)

	config.GetQyComConfigOr(c.qyOptions)

	cacheIns, err := cacheV1.GetCacheInsOr(storeIns)
	if err != nil {
		log.Fatalf("Failed to get cache instance: %s", err.Error())
	}

	pb.RegisterCacheServer(grpcServer, cacheIns)

	reflection.Register(grpcServer)
	return &grpcAPIServer{grpcServer, c.Addr}, nil
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

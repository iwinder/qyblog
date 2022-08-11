package db

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	userv1 "github.com/iwinder/qingyucms/api/qycms_user/v1"
	"github.com/iwinder/qingyucms/app/qycms_blog/internal/conf"
	"github.com/iwinder/qingyucms/app/qycms_blog/internal/data/po"
	clientv3 "go.etcd.io/etcd/client/v3"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
	"sync"
	"time"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData,
	NewRegistrar,
	//NewDiscovery,

	NewArticleRepo,
	NewArticleContentRepo,
	NewUserRepo,
	NewUserServiceClient,
)

// Data .
type Data struct {
	// TODO wrapped database client
	db       *gorm.DB
	redisCli redis.Cmdable
	uc       userv1.UserClient
	log      *log.Helper
}

var (
	mysqlDb *Data
	once    sync.Once
)

// NewData .
func NewData(conf *conf.Data, logger log.Logger, uc userv1.UserClient) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}

	if strings.EqualFold(conf.Database.Source, "") && mysqlDb.db == nil {
		return &Data{}, cleanup, fmt.Errorf("MySql DB Open failed")
	}
	redisOpen := true
	if strings.EqualFold(conf.Redis.Addr, "") && mysqlDb.redisCli == nil {
		fmt.Errorf("Redis DB Open failed")
		redisOpen = false
	}

	var err error
	var dbIns *gorm.DB
	var redisCliDB redis.Cmdable
	l := log.NewHelper(log.With(logger, "module", "mysql/data"))
	once.Do(func() {
		dbIns, err = gorm.Open(mysql.Open(conf.Database.Source), &gorm.Config{})
		// redis
		if redisOpen {
			redisCliDB = redis.NewClient(&redis.Options{
				Addr:         conf.Redis.Addr,
				Password:     conf.Redis.Password,
				ReadTimeout:  conf.Redis.ReadTimeout.AsDuration(),
				WriteTimeout: conf.Redis.WriteTimeout.AsDuration(),
				DialTimeout:  time.Second * 2,
				PoolSize:     10,
			})
			timeout, cancelFunc := context.WithTimeout(context.Background(), time.Second*2)
			defer cancelFunc()
			err = redisCliDB.Ping(timeout).Err()
			if err != nil {
				log.Fatalf("redis connect error: %v", err)
			}
		}

		mysqlDb = &Data{
			db:       dbIns,
			log:      l,
			uc:       uc,
			redisCli: redisCliDB,
		}
		AutoMigrateTable(dbIns)
	})
	if mysqlDb.db == nil || err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}

	return mysqlDb, cleanup, nil
}

// AutoMigrateTable 初始化table
func AutoMigrateTable(dbIns *gorm.DB) {
	dbIns.AutoMigrate(&po.ArticlePO{}, &po.ArticleContentPO{})
}

func NewUserServiceClient(tp *tracesdk.TracerProvider) userv1.UserClient {
	//func NewUserServiceClient() userv1.UserClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("127.0.0.1:9000"),
		//grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			tracing.Client(tracing.WithTracerProvider(tp)),
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	c := userv1.NewUserClient(conn)
	return c
}

//
func NewRegistrar(conf *conf.Registry) registry.Registrar {
	// 注册服务
	// new etcd client
	client, err := clientv3.New(clientv3.Config{
		Endpoints: []string{conf.Etcd.Address},
	})
	if err != nil {
		panic(err)
	}
	// new reg with etcd client
	reg := etcd.New(client)
	return reg
}

//
//func NewDiscovery(conf *conf.Registry) registry.Discovery {
//
//	// new etcd client
//	client, err := clientv3.New(clientv3.Config{
//		Endpoints: []string{conf.Etcd.Address},
//	})
//	if err != nil {
//		panic(err)
//	}
//	// new dis with etcd client
//	dis := etcd.New(client)
//
//	//endpoint := "discovery:///qycms.user.server"
//	//conn, err := grpc.Dial(context.Background(), grpc.WithEndpoint(endpoint), grpc.WithDiscovery(dis))
//	//if err != nil {
//	//	panic(err)
//	//}
//
//	return dis
//}

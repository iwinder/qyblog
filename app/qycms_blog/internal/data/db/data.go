package db

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
	userv1 "github.com/iwinder/qingyucms/api/qycms_user/v1"
	"github.com/iwinder/qingyucms/app/qycms_blog/internal/conf"
	"github.com/iwinder/qingyucms/app/qycms_blog/internal/data/po"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
	"sync"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewArticleRepo, NewArticleContentRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	db  *gorm.DB
	log *log.Helper
}

var (
	mysqlDb *Data
	once    sync.Once
)

// NewData .
func NewData(conf *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}

	if strings.EqualFold(conf.Database.Source, "") && mysqlDb.db == nil {
		return &Data{}, cleanup, fmt.Errorf("MySql DB Open failed")
	}
	var err error
	var dbIns *gorm.DB
	l := log.NewHelper(log.With(logger, "module", "mysql/data"))
	once.Do(func() {
		dbIns, err = gorm.Open(mysql.Open(conf.Database.Source), &gorm.Config{})
		mysqlDb = &Data{
			db:  dbIns,
			log: l,
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
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("122.111.11.1:8080"),
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

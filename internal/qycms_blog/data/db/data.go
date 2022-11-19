package db

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"github.com/iwinder/qingyucms/internal/qycms_blog/conf"
	blogPo "github.com/iwinder/qingyucms/internal/qycms_blog/data/po"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	glogger "gorm.io/gorm/logger"
	"strings"
	"sync"
	"time"
)

// ProviderSet is cdata providers.
var ProviderSet = wire.NewSet(NewData, NewCasbinData, NewCasbinRuleRepo,
	NewUserRoleRepo, NewRoleMenusRepo, NewRoleApiRepo,
	NewUserRepo, NewRoleRepo, NewSiteConfigRepo,
	NewMenusAdminRepo, NewApiRepo, NewApiGroupRepo,
	NewFileLibTypeRepo, NewFileLibConfigRepo, NewFileLibRepo,
	NewArticleRepo, NewArticleContentRepo,
	NewCommentAgentRepo, NewCommentIndexRepo, NewCommentContentRepo,
	NewLinkRepo, NewShortLinkRepo, NewMenusRepo, NewMenusAgentRepo,
	NewTagsRepo, NewCategoryRepo, NewArticleTagsRepo,
)

// Data .
type Data struct {
	// TODO wrapped database client
	Db       *gorm.DB
	RedisCli redis.Cmdable
	log      *log.Helper
}

var (
	mysqlDb *Data
	once    sync.Once
)

// NewData .
func NewData(conf *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the cdata resources")
	}

	if strings.EqualFold(conf.Database.Source, "") && mysqlDb.Db == nil {
		return &Data{}, cleanup, fmt.Errorf("MySql DB Open failed")
	}
	redisOpen := true
	if strings.EqualFold(conf.Redis.Addr, "") && mysqlDb.RedisCli == nil {
		fmt.Errorf("Redis DB Open failed")
		redisOpen = false
	}

	var err error
	var dbIns *gorm.DB
	var redisCliDB redis.Cmdable
	l := log.NewHelper(log.With(logger, "module", "mysql/cdata"))
	once.Do(func() {
		dbIns, err = gorm.Open(mysql.Open(conf.Database.Source), &gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
			Logger:                                   glogger.Default.LogMode(glogger.Info),
		})
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
			Db:  dbIns,
			log: l,
			//Uc:       uc,
			RedisCli: redisCliDB,
		}
		AutoMigrateTable(dbIns)
	})
	if mysqlDb.Db == nil || err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}

	return mysqlDb, cleanup, nil
}

// AutoMigrateTable 初始化table
func AutoMigrateTable(dbIns *gorm.DB) {
	dbIns.AutoMigrate(
		&blogPo.UserPO{}, &blogPo.RolePO{}, &blogPo.SiteConfigPO{},
		&blogPo.ApiPO{}, &blogPo.ApiGroupPO{}, &blogPo.MenusAdminPO{},
		&blogPo.UserRolePO{}, &blogPo.RoleApiPO{}, &blogPo.RoleMenusPO{},
		&blogPo.FileLibTypePO{}, &blogPo.FileLibConfigPO{}, &blogPo.FileLibPO{},
		&blogPo.LinkPO{}, &blogPo.ShortLinkPO{}, &blogPo.MenusAgentPO{}, &blogPo.MenusPO{},
		&blogPo.TagsPO{}, &blogPo.CategoryPO{},
		&blogPo.CommentAgentPO{}, &blogPo.CommentIndexPO{}, &blogPo.CommentContentPO{},
		&blogPo.ArticlePO{}, &blogPo.ArticleContentPO{}, &blogPo.ArticleTagsPO{},
	)

}

func withFilterKeyLikeValue(key, value string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(key+" like ? ", value)
	}
}

func withFilterKeyEquarlsValue(key string, value interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(key+" = ? ", value)
	}
}

func withFilterKeyInValue(key string, value interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(key+" in ? ", value)
	}
}

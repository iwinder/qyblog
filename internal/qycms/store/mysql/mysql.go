package mysql

import (
	"fmt"
	db "gitee.com/windcoder/qingyucms/internal/pkg/qycms-db"
	logger "gitee.com/windcoder/qingyucms/internal/pkg/qycms-logger"
	genericoption "gitee.com/windcoder/qingyucms/internal/pkg/qycms-options"
	v1 "gitee.com/windcoder/qingyucms/internal/qycms/models/v1"
	"gitee.com/windcoder/qingyucms/internal/qycms/store"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"sync"
)

type datastore struct {
	db *gorm.DB
}

func (ds *datastore) Users() store.UserStore {
	return newUserStore(ds)
}

func (ds *datastore) CommonDB() store.CommonStore {
	return newCommonDB(ds)
}

func (ds *datastore) Close() error {
	db, err := ds.db.DB()
	if err != nil {
		return errors.Wrap(err, "get gorm db instance failed")
	}

	return db.Close()
}

func (ds *datastore) InitTables() error {
	return ds.db.AutoMigrate(
		v1.User{},
		v1.Role{},
		v1.Privilege{},
		v1.Menu{},
		v1.CasbinRule{},
	)
}

var (
	mysqlFactory store.Factory
	once         sync.Once
)

func GetMySQLFactoryOr(opts *genericoption.MySQLOptions) (store.Factory, error) {
	if opts == nil && mysqlFactory == nil {
		return nil, fmt.Errorf("failed to get mysql store fatory")
	}

	var err error
	var dbIns *gorm.DB
	once.Do(func() {
		options := &db.Optios{
			Host:                  opts.Host,
			Username:              opts.Username,
			Password:              opts.Password,
			Database:              opts.Database,
			MaxIdleConnections:    opts.MaxIdleConnections,
			MaxOpenConnections:    opts.MaxOpenConnections,
			MaxConnectionLifeTime: opts.MaxConnectionLifeTime,
			LogLevel:              opts.LogLevel,
			Logger:                logger.New(opts.LogLevel),
		}
		dbIns, err = db.New(options)
		mysqlFactory = &datastore{dbIns}
	})

	if mysqlFactory == nil || err != nil {
		return nil, fmt.Errorf("failed to get mysql store fatory, mysqlFactory: %+v, error: %w", mysqlFactory, err)
	}

	return mysqlFactory, nil
}

package base

import (
	"fmt"
	"github.com/iwinder/qyblog/internal/pkg/db"
	log "github.com/iwinder/qyblog/internal/pkg/logger"
	genericoptions "github.com/iwinder/qyblog/internal/pkg/options/base"
	"gorm.io/gorm"
	"sync"
)

type MySqlStore struct {
	Client *gorm.DB
}

var (
	mysqlStore *MySqlStore
	once       sync.Once
)

func (s *MySqlStore) Close() {

}
func GetMySQLFactoryOr(opts *genericoptions.MySQLOptions) (*MySqlStore, error) {
	if opts == nil && mysqlStore == nil {
		return nil, fmt.Errorf("failed to get mysql store fatory")
	}

	var err error
	var mysqlClient *gorm.DB
	once.Do(func() {
		options := &db.Options{
			Host:                  opts.Host,
			Username:              opts.Username,
			Password:              opts.Password,
			Database:              opts.Database,
			MaxIdleConnections:    opts.MaxIdleConnections,
			MaxOpenConnections:    opts.MaxOpenConnections,
			MaxConnectionLifeTime: opts.MaxConnectionLifeTime,
			LogLevel:              opts.LogLevel,
			Logger:                log.NewGormLogger(opts.LogLevel),
		}
		mysqlClient, err = db.New(options)

		mysqlStore = &MySqlStore{Client: mysqlClient}
	})

	if mysqlStore.Client == nil || err != nil {
		return nil, fmt.Errorf("failed to get mysql store fatory, mysqlFactory: %+v, error: %w", mysqlStore, err)
	}

	return mysqlStore, nil
}

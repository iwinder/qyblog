package mysql

import (
	v1 "gitee.com/windcoder/qingyucms/internal/qycms/models/v1"
	"gitee.com/windcoder/qingyucms/internal/qycms/store"
	"github.com/pkg/errors"
	"gorm.io/gorm"
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

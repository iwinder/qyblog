package mysql

import "gorm.io/gorm"

type commonDB struct {
	db *gorm.DB
}

func newCommonDB(ds *datastore) *commonDB {
	return &commonDB{
		db: ds.db,
	}
}

func (c *commonDB) GetCommonDB() *gorm.DB {
	return c.db
}

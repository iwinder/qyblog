package mysql

import "gorm.io/gorm"

type Privileges struct {
	db *gorm.DB
}

func NewPrivileges(da *datastore) *Privileges {
	return &Privileges{db: da.db}
}

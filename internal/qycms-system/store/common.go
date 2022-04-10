package store

import "gorm.io/gorm"

type CommonStore interface {
	GetcommonDB() *gorm.DB
}

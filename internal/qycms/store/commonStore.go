package store

import "gorm.io/gorm"

type CommonStore interface {
	GetCommonDB() *gorm.DB
}

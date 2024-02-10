package accountstorage

import (
	"gorm.io/gorm"
)

type accountStorage struct {
	db *gorm.DB
}

func NewAccountStorage(db *gorm.DB) *accountStorage {
	return &accountStorage{db: db}
}

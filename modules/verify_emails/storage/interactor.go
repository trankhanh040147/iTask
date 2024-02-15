package verifyemailsstorage

import (
	"gorm.io/gorm"
)

type verifyEmailsStorage struct {
	db *gorm.DB
}

func NewVerifyEmailsStorage(db *gorm.DB) *verifyEmailsStorage {
	return &verifyEmailsStorage{db: db}
}

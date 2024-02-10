package paymentstorage

import (
	"gorm.io/gorm"
)

type paymentStorage struct {
	db *gorm.DB
}

func NewPaymentStorage(db *gorm.DB) *paymentStorage {
	return &paymentStorage{db: db}
}

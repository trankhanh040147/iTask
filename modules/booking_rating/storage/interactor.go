package bookingratingstorage

import (
	"gorm.io/gorm"
)

type bookingratingstorage struct {
	db *gorm.DB
}

func Newbookingratingstorage(db *gorm.DB) *bookingratingstorage {
	return &bookingratingstorage{db: db}
}

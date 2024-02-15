package bookingstorage

import "gorm.io/gorm"

type bookingStorage struct {
	db *gorm.DB
}

func NewBookingStorage(db *gorm.DB) *bookingStorage {
	return &bookingStorage{db: db}
}

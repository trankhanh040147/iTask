package bookingdetailstorage

import "gorm.io/gorm"

type bookingDetailStorage struct {
	db *gorm.DB
}

func NewBookingDetailStorage(db *gorm.DB) *bookingDetailStorage {
	return &bookingDetailStorage{db: db}
}

package amenitystorage

import "gorm.io/gorm"

type amenityStorage struct {
	db *gorm.DB
}

func NewAmenityStorage(db *gorm.DB) *amenityStorage {
	return &amenityStorage{db: db}
}

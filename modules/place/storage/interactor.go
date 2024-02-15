package placestorage

import (
	"gorm.io/gorm"
)

type placeStorage struct {
	db *gorm.DB
}

func NewPlaceStorage(db *gorm.DB) *placeStorage {
	return &placeStorage{db: db}
}

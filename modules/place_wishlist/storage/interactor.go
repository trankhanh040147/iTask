package placewishliststorage

import "gorm.io/gorm"

type placeWishListStorage struct {
	db *gorm.DB
}

func NewPlaceWishListStorage(db *gorm.DB) *placeWishListStorage {
	return &placeWishListStorage{db: db}
}

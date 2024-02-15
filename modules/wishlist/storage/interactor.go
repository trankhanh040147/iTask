package wishliststorage

import (
	"gorm.io/gorm"
)

type wishListStorage struct {
	db *gorm.DB
}

func NewWishListStorage(db *gorm.DB) *wishListStorage {
	return &wishListStorage{db: db}
}

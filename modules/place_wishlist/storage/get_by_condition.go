package placewishliststorage

import (
	"context"
	"iTask/entities"
)

func (s *placeWishListStorage) GetByCondition(ctx context.Context, condition map[string]interface{}) ([]entities.PlaceWishList, error) {
	db := s.db

	var placeWishLists []entities.PlaceWishList
	if err := db.Where(condition).Find(&placeWishLists).Error; err != nil {
		return nil, err
	}

	return placeWishLists, nil
}

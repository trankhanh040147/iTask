package placestorage

import (
	"context"
	"iTask/common"
	"iTask/entities"
)

func (s *placeStorage) ListPlaceInIDs(ctx context.Context, placeIds []int) ([]entities.Place, error) {
	db := s.db

	var data []entities.Place

	if err := db.Where("id in (?)", placeIds).Find(&data).Error; err != nil {
		return nil, common.ErrorDB(err)
	}

	return data, nil
}

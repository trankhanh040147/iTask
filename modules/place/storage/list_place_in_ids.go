package placestorage

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/entities"
)

func (s *placeStorage) ListPlaceInIds(ctx context.Context, placeIds []int, vendorId int) ([]entities.Place, error) {
	db := s.db

	var data []entities.Place

	db = db.Where("vendor_id = ?", vendorId)
	if err := db.Where("id in (?)", placeIds).Find(&data).Error; err != nil {
		return nil, common.ErrorDB(err)
	}

	return data, nil
}

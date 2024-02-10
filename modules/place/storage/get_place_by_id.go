package placestorage

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/entities"

	"gorm.io/gorm"
)

func (s *placeStorage) GetPlaceByID(ctx context.Context, id int) (*entities.Place, error) {
	db := s.db

	var data entities.Place

	if err := db.Where("id = ?", id).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, common.ErrorDB(err)
	}

	return &data, nil
}

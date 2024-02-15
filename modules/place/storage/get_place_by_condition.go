package placestorage

import (
	"context"
	"paradise-booking/entities"

	"gorm.io/gorm"
)

func (s *placeStorage) GetPlaceByCondition(ctx context.Context, condition map[string]interface{}) ([]entities.Place, error) {
	db := s.db

	var data []entities.Place

	if err := db.Where(condition).Find(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return data, nil
}

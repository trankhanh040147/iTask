package bookingstorage

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/entities"

	"gorm.io/gorm"
)

func (s *bookingStorage) ListPlaceIds(ctx context.Context) ([]int, error) {
	db := s.db

	var data []int

	if err := db.Model(entities.Booking{}).Distinct().Pluck("place_id", &data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, common.ErrorDB(err)
	}

	return data, nil
}

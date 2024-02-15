package bookingratingstorage

import (
	"context"
	"iTask/common"
	"iTask/entities"
)

func (s *bookingratingstorage) GetByCondition(ctx context.Context, condition map[string]interface{}) ([]entities.BookingRating, error) {
	db := s.db

	var data []entities.BookingRating

	if err := db.Where(condition).Find(&data).Error; err != nil {
		return nil, common.ErrorDB(err)
	}

	return data, nil
}

// trigger deploy

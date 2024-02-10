package bookingratingstorage

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/entities"
)

func (s *bookingratingstorage) Create(ctx context.Context, data *entities.BookingRating) (*entities.BookingRating, error) {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return nil, common.ErrorDB(err)
	}

	return data, nil
}

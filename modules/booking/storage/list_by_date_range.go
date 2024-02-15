package bookingstorage

import (
	"context"
	"iTask/entities"
	"time"
)

func (s *bookingStorage) GetBookingsWithinDateRange(ctx context.Context, dateFrom, dateTo *time.Time) ([]entities.Booking, error) {
	db := s.db

	var data []entities.Booking

	if err := db.Raw("call GetBookingsWithinRange(?,?)", dateFrom, dateTo).Scan(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

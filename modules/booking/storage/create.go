package bookingstorage

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/entities"
)

func (s *bookingStorage) Create(ctx context.Context, data *entities.Booking) (err error) {
	db := s.db.Begin()

	if err := db.Create(data).Error; err != nil {
		db.Rollback()
		return common.ErrorDB(err)
	}

	db.Commit()
	return nil
}

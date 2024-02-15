package bookingstorage

import (
	"context"
	"iTask/common"
	"iTask/entities"
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

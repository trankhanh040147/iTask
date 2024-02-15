package bookingstorage

import (
	"context"
	"iTask/common"
	"iTask/entities"
)

func (s *bookingStorage) DeleteByID(ctx context.Context, id int) error {
	db := s.db

	if err := db.Where("id = ?", id).Delete(&entities.Booking{}).Error; err != nil {
		return common.ErrorDB(err)
	}

	return nil
}

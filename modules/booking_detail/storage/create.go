package bookingdetailstorage

import (
	"context"
	"iTask/common"
	"iTask/entities"
)

func (s *bookingDetailStorage) Create(ctx context.Context, data *entities.BookingDetail) (err error) {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrorDB(err)
	}

	return nil
}

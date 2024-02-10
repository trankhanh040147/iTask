package bookingstorage

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/entities"
)

func (s *bookingStorage) UpdateStatus(ctx context.Context, bookingID int, status int) error {
	db := s.db
	booking := entities.Booking{}
	if err := db.Table(booking.TableName()).Where("id = ?", bookingID).Update("status_id", status).Error; err != nil {
		return common.ErrorDB(err)
	}
	return nil
}

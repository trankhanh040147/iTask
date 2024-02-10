package bookingdetailstorage

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/entities"
)

func (s *bookingDetailStorage) GetByBookingID(ctx context.Context, bookingId int) (res *entities.BookingDetail, err error) {
	db := s.db

	var data entities.BookingDetail
	if err := db.Where("booking_id = ?", bookingId).First(&data).Error; err != nil {
		return nil, common.ErrorDB(err)
	}

	return &data, nil
}

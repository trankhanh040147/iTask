package bookingratingstorage

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/entities"
)

func (s *bookingratingstorage) GetByVendorID(ctx context.Context, vendorID int) ([]entities.BookingRating, error) {
	db := s.db

	var data []entities.BookingRating

	if err := db.Raw("call GetCommentsAndRatingsByVendorId(?)", vendorID).Scan(&data).Error; err != nil {
		return nil, common.ErrorDB(err)
	}

	return data, nil
}

package bookingstorage

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/entities"
)

func (s *bookingStorage) ListAllBookingWithCondition(ctx context.Context, condition []common.Condition) ([]entities.Booking, error) {
	var data []entities.Booking

	db := s.db
	db = db.Table(entities.Booking{}.TableName())

	for _, v := range condition {
		query := v.BuildQuery()
		db = db.Where(query+" ?", v.Value)
	}

	if err := db.Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

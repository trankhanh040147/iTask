package bookingstorage

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/entities"
)

func (s *bookingStorage) GetByPlaceID(ctx context.Context, placeId int, paging *common.Paging) ([]entities.Booking, error) {
	db := s.db

	var data []entities.Booking

	db = db.Table(entities.Booking{}.TableName()).Where("place_id = ?", placeId)

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrorDB(err)
	}

	if err := db.Offset((paging.Page - 1) * paging.Limit).Limit(paging.Limit).Find(&data).Error; err != nil {
		return nil, common.ErrorDB(err)
	}

	return data, nil
}

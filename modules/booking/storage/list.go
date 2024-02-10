package bookingstorage

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/entities"
	"paradise-booking/modules/booking/iomodel"
	"paradise-booking/utils"
)

func (s *bookingStorage) ListByFilter(ctx context.Context, filter *iomodel.FilterListBooking, paging *common.Paging, userId int) ([]entities.Booking, error) {
	db := s.db

	var data []entities.Booking

	db = db.Table(entities.Booking{}.TableName())

	db = db.Where("user_id = ?", userId)
	if v := filter; v != nil {
		if len(v.Statuses) > 0 && v.Statuses[0] != 0 {
			db = db.Where("status_id in (?) ", v.Statuses)
		}

		if v.DateFrom != "" {
			dateTime, _ := utils.ParseStringToTime(v.DateFrom)
			db = db.Where("created_at >= ?", dateTime)
		}
		if v.DateTo != "" {
			dateTime, _ := utils.ParseStringToTime(v.DateTo)
			db = db.Where("created_at <= ?", dateTime)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrorDB(err)
	}

	if err := db.Offset((paging.Page - 1) * paging.Limit).Limit(paging.Limit).Find(&data).Error; err != nil {
		return nil, common.ErrorDB(err)
	}

	return data, nil
}

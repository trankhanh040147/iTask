package bookingratingstorage

import (
	"context"
	"iTask/common"
	"iTask/entities"
)

func (s *bookingratingstorage) GetStatisticByPlaceID(ctx context.Context, placeId int64) ([]entities.StatisticResp, error) {
	db := s.db

	var data []entities.StatisticResp

	if err := db.Raw("call GetRatingStatisticByPlaceId(?)", placeId).Scan(&data).Error; err != nil {
		return nil, common.ErrorDB(err)
	}

	return data, nil
}

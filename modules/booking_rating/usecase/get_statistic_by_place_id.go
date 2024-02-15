package bookingratingusecase

import (
	"context"
	"iTask/entities"
)

func (uc *bookingRatingUsecase) GetStatisticByPlaceID(ctx context.Context, placeId int) ([]entities.StatisticResp, error) {
	res, err := uc.BookingRatingSto.GetStatisticByPlaceID(ctx, int64(placeId))
	if err != nil {
		return nil, err
	}

	return res, nil
}

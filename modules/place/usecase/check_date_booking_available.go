package placeusecase

import (
	"context"
	"paradise-booking/constant"
	"paradise-booking/utils"
)

func (uc *placeUseCase) CheckDateBookingAvailable(ctx context.Context, placeId int64, dateFrom string, dateTo string) (isValid *bool, err error) {

	place, err := uc.placeStorage.GetPlaceByID(ctx, int(placeId))
	if err != nil {
		return nil, err
	}

	timeFrom, _ := utils.ParseStringToTime(dateFrom)
	timeTo, _ := utils.ParseStringToTime(dateTo)
	bookeds, err := uc.bookingSto.GetBookingsWithinDateRange(ctx, timeFrom, timeTo)
	if err != nil {
		return nil, err
	}

	bookedCount := 0
	for _, booked := range bookeds {
		if booked.StatusId == constant.BookingStatusCancel || booked.StatusId == constant.BookingStatusCompleted {
			continue
		}
		if booked.PlaceId == int(placeId) {
			bookedCount += 1
		}
	}

	res := false
	if bookedCount >= place.NumPlaceOriginal {
		return &res, nil
	}

	res = true
	return &res, nil
}

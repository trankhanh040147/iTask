package placeusecase

import (
	"context"
	"paradise-booking/constant"
	"paradise-booking/modules/place/iomodel"
	"paradise-booking/utils"
)

func (uc *placeUseCase) GetStatusPlaceToBook(ctx context.Context, placeId int, dateFrom, dateTo string) (*iomodel.GetStatusPlaceToBookResp, error) {
	var result iomodel.GetStatusPlaceToBookResp

	place, err := uc.placeStorage.GetPlaceByID(ctx, placeId)
	if err != nil {
		return nil, err
	}

	timeFrom, _ := utils.ParseStringToTime(dateFrom)
	timeTo, _ := utils.ParseStringToTime(dateTo)
	bookeds, err := uc.bookingSto.GetBookingsWithinDateRange(ctx, timeFrom, timeTo)
	if err != nil {
		return nil, err
	}

	cntNumIsBooking := 0
	for _, booked := range bookeds {
		if booked.PlaceId != placeId {
			continue
		}

		if booked.StatusId != constant.BookingStatusCancel && booked.StatusId != constant.BookingStatusCompleted {
			cntNumIsBooking += 1
		}

		dateF := utils.ParseTimeToString(booked.CheckInDate)
		dateT := utils.ParseTimeToString(booked.ChekoutDate)
		result.BookingPlaceHistory = append(result.BookingPlaceHistory, iomodel.BookingPlaceHistoryResp{
			DateFrom:  dateF,
			DateTo:    dateT,
			BookingID: booked.Id,
		})
	}

	result.NumPlaceBooked = cntNumIsBooking
	result.NumPlaceOriginal = place.NumPlaceOriginal
	result.NumPlaceRemain = place.NumPlaceOriginal - cntNumIsBooking

	return &result, nil
}

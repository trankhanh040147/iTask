package bookingusecase

import (
	"context"
	"iTask/common"
	"iTask/entities"
	"iTask/modules/booking/convert"
	"iTask/modules/booking/iomodel"
)

func (uc *bookingUseCase) GetBookingByPlaceID(ctx context.Context, placeId int, paging *common.Paging) ([]iomodel.GetBookingByPlaceResp, error) {

	paging.Process()
	bookings, err := uc.bookingSto.GetByPlaceID(ctx, placeId, paging)
	if err != nil {
		return nil, err
	}

	var result []iomodel.GetBookingByPlaceResp

	// get place by id
	place, err := uc.PlaceSto.GetPlaceByID(ctx, placeId)
	if err != nil {
		return nil, common.ErrCannotGetEntity("place", err)
	}

	for _, booking := range bookings {

		bookingDetail := &entities.BookingDetail{}
		// get account by id
		bookingDetail, err = uc.bookingDetailSto.GetByBookingID(ctx, booking.Id)
		if err != nil {
			return nil, common.ErrCannotGetEntity("booking detail", err)
		}

		account, err := uc.AccountSto.GetProfileByID(ctx, booking.UserId)
		if err != nil {
			if err == common.RecordNotFound {
				// case: booking when user not login
				account = &entities.Account{
					Email:    bookingDetail.Email,
					Username: bookingDetail.FullName,
					Phone:    bookingDetail.Phone,
				}
			} else {
				return nil, common.ErrCannotGetEntity("account", err)
			}
		}

		result = append(result, *convert.ConvertBookingModelToGetByPlaceResp(account, &booking, place, bookingDetail))

	}

	return result, nil
}

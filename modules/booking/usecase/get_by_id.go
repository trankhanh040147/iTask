package bookingusecase

import (
	"context"
	"iTask/common"
	"iTask/entities"
	"iTask/modules/booking/convert"
	"iTask/modules/booking/iomodel"
)

func (uc *bookingUseCase) GetBookingByID(ctx context.Context, id int) (*iomodel.GetBookingResp, error) {
	booking, err := uc.bookingSto.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if booking == nil {
		return nil, common.ErrEntityNotFound("place", err)
	}

	// get booking detail by id
	bookingDetail, err := uc.bookingDetailSto.GetByBookingID(ctx, booking.Id)
	if err != nil {
		return nil, common.ErrCannotGetEntity("booking detail", err)
	}

	// get account by id
	account, err := uc.AccountSto.GetProfileByID(ctx, booking.UserId)
	if err != nil {
		if err == common.RecordNotFound {
			// case: booking when user not login
			account = &entities.Account{
				Email:    bookingDetail.Email,
				FullName: bookingDetail.FullName,
				Phone:    bookingDetail.Phone,
			}
		} else {
			return nil, common.ErrCannotGetEntity("account", err)
		}
	}

	// get place by id
	place, err := uc.PlaceSto.GetPlaceByID(ctx, booking.PlaceId)
	if err != nil {
		return nil, common.ErrCannotGetEntity("place", err)
	}

	result := convert.ConvertBookingModelToGetResp(account, booking, place, bookingDetail)
	return result, nil
}

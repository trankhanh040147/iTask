package bookingusecase

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/modules/booking/convert"
	"paradise-booking/modules/booking/iomodel"
)

func (uc *bookingUseCase) ListBooking(ctx context.Context, paging *common.Paging, filter *iomodel.FilterListBooking, userID int) (*iomodel.ListBookingResp, error) {

	// Assign permissions by userid
	if paging != nil {
		paging.Process()
	}
	result, err := uc.bookingSto.ListByFilter(ctx, filter, paging, userID)

	if err != nil {
		return nil, common.ErrCannotListEntity("booking", err)
	}

	// get user by id
	var res iomodel.ListBookingResp

	user, err := uc.AccountSto.GetProfileByID(ctx, userID)
	res.UserId = user.Id
	res.User = *user
	for _, booking := range result {
		// get user	by id
		if err != nil {
			return nil, common.ErrCannotGetEntity("user", err)
		}

		// get place by id
		place, err := uc.PlaceSto.GetPlaceByID(ctx, booking.PlaceId)
		if err != nil {
			return nil, common.ErrCannotGetEntity("place", err)
		}
		res.ListData = append(res.ListData, convert.ConvertBookingModelToListBooking(booking, place))
	}

	return &res, nil
}

package bookingusecase

import (
	"context"
	"errors"
	"iTask/common"
	"iTask/constant"
	"iTask/entities"
)

func (uc *bookingUseCase) CancelBooking(ctx context.Context, bookingID int) error {
	// get booking by id
	booking, err := uc.bookingSto.GetByID(ctx, bookingID)
	if err != nil {
		return common.ErrCannotGetEntity(entities.Booking{}.TableName(), err)
	}

	// check status must be pending
	if booking.StatusId != constant.BookingStatusPending {
		return errors.New("can not cancel booking because status is not pending")
	}
	// update status booking
	if err := uc.bookingSto.UpdateStatus(ctx, bookingID, constant.BookingStatusCancel); err != nil {
		return common.ErrCannotUpdateEntity(entities.Booking{}.TableName(), err)
	}
	return nil
}

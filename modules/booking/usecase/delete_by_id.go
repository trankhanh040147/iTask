package bookingusecase

import (
	"context"
	"errors"
	"paradise-booking/constant"
)

func (uc *bookingUseCase) DeleteBookingByID(ctx context.Context, id int) error {

	booking, err := uc.bookingSto.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if booking.StatusId != constant.BookingStatusCancel && booking.StatusId != constant.BookingStatusCompleted {
		return errors.New("booking is processing, can not delete")
	}

	err = uc.bookingSto.DeleteByID(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

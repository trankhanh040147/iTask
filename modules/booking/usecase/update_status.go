package bookingusecase

import (
	"context"
	"iTask/common"
	"iTask/entities"
)

func (uc *bookingUseCase) UpdateStatusBooking(ctx context.Context, bookingID, status int) error {
	// update status booking
	if err := uc.bookingSto.UpdateStatus(ctx, bookingID, status); err != nil {
		return common.ErrCannotUpdateEntity(entities.Booking{}.TableName(), err)
	}

	return nil
}

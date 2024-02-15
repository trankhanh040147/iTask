package bookingusecase

import (
	"context"
	"iTask/common"
	"iTask/constant"
	"iTask/entities"
	"time"
)

func (uc *bookingUseCase) ListBookingByCondition(ctx context.Context) ([]entities.Booking, error) {
	conditions := []common.Condition{}
	conditions = append(conditions, common.Condition{
		Field:    "created_at",
		Operator: "<=",
		Value:    time.Now().AddDate(0, 0, -1),
	})
	conditions = append(conditions, common.Condition{
		Field:    "created_at",
		Operator: ">=",
		Value:    time.Now().AddDate(0, 0, -2),
	})
	conditions = append(conditions, common.Condition{
		Field:    "status_id",
		Operator: "!=",
		Value:    constant.BookingStatusCancel,
	})

	result, err := uc.bookingSto.ListAllBookingWithCondition(ctx, conditions)
	if err != nil {
		return nil, common.ErrCannotListEntity("booking", err)
	}
	return result, nil
}

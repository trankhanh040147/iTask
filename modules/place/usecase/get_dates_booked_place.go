package placeusecase

import (
	"context"
	"iTask/common"
	"iTask/constant"
	"iTask/entities"
	"iTask/utils"
	"time"

	"github.com/samber/lo"
)

func (uc *placeUseCase) GetDatesBookedPlace(ctx context.Context, placeId int) ([][]string, error) {
	conditions := []common.Condition{}
	conditions = append(conditions, common.Condition{
		Field:    "place_id",
		Operator: common.OperatorEqual,
		Value:    placeId,
	})
	conditions = append(conditions, common.Condition{
		Field:    "status_id",
		Operator: common.OperatorNotEqual,
		Value:    constant.BookingStatusCancel,
	})
	conditions = append(conditions, common.Condition{
		Field:    "status_id",
		Operator: common.OperatorNotEqual,
		Value:    constant.BookingStatusCompleted,
	})

	bookings, err := uc.bookingSto.ListAllBookingWithCondition(ctx, conditions)
	if err != nil {
		return nil, err
	}

	timeNow := time.Now()
	dates := lo.FilterMap(bookings, func(item entities.Booking, _ int) ([]string, bool) {
		if item.CheckInDate.After(timeNow) {
			checkinDate := utils.ParseTimeToString(item.CheckInDate)
			checkoutDate := utils.ParseTimeToString(item.ChekoutDate)
			return []string{checkinDate, checkoutDate}, true
		}
		return []string{}, false
	})

	return dates, nil
}

package bookingusecase

import (
	"context"
	"errors"
	"iTask/common"
	"iTask/constant"
	"iTask/entities"
	"iTask/modules/booking/iomodel"

	"github.com/samber/lo"
)

func (uc *bookingUseCase) ListPlaceReservationByVendor(ctx context.Context, vendorId, placeId int) (*iomodel.ListBookingPlaceReservationResp, error) {

	var result iomodel.ListBookingPlaceReservationResp
	if placeId != 0 {
		place, err := uc.PlaceSto.GetPlaceByID(ctx, placeId)
		if err != nil {
			return nil, err
		}

		if place == nil {
			return nil, common.ErrEntityNotFound(entities.Place{}.TableName(), errors.New("place not found"))
		}

		// get place in booking
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

		isBooked := false
		if len(bookings) > 0 {
			isBooked = true
		}

		res := iomodel.BookingPlaceResp{
			place,
			isBooked,
		}
		return &iomodel.ListBookingPlaceReservationResp{Data: []iomodel.BookingPlaceResp{res}}, nil
	}

	conditions := []common.Condition{}
	conditions = append(conditions, common.Condition{
		Field:    "vendor_id",
		Operator: common.OperatorEqual,
		Value:    vendorId,
	})

	places, err := uc.PlaceSto.ListPlaceByCondition(ctx, conditions)
	if err != nil {
		return nil, err
	}

	placeIds := lo.Map(places, func(data entities.Place, _ int) int {
		return data.Id
	})

	if len(placeIds) == 0 {
		return &result, nil
	}

	conditions = []common.Condition{}
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
	conditions = append(conditions, common.Condition{
		Field:    "place_id",
		Operator: common.OperatorIn,
		Value:    placeIds,
	})

	bookings, err := uc.bookingSto.ListAllBookingWithCondition(ctx, conditions)
	if err != nil {
		return nil, err
	}

	mapPlaceBooked := map[int]bool{}
	for _, item := range bookings {
		mapPlaceBooked[item.PlaceId] = true
	}

	lo.ForEach(places, func(item entities.Place, _ int) {
		isBooked := false
		if _, ok := mapPlaceBooked[item.Id]; ok {
			isBooked = true
		}

		res := iomodel.BookingPlaceResp{
			&item,
			isBooked,
		}
		result.Data = append(result.Data, res)
	})

	return &result, nil
}

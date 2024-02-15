package bookingratingusecase

import (
	"context"
	"log"
	"paradise-booking/common"
	"paradise-booking/entities"
	"paradise-booking/modules/booking_rating/iomodel"
)

func (uc *bookingRatingUsecase) GetCommentByBookingID(ctx context.Context, bookingID int) ([]iomodel.GetCommentResp, error) {
	res, err := uc.BookingRatingSto.GetByCondition(ctx, map[string]interface{}{"booking_id": bookingID})
	if err != nil {
		return nil, common.ErrCannotGetEntity(entities.BookingRating{}.TableName(), err)
	}

	var result []iomodel.GetCommentResp

	for _, bookingRate := range res {
		user, err := uc.AccountSto.GetProfileByID(ctx, bookingRate.UserId)
		if err != nil {
			log.Printf("Error when get user profile by id: %v\n", err)
			continue
		}

		place, err := uc.PlaceSto.GetPlaceByID(ctx, bookingRate.PlaceId)
		if err != nil {
			log.Printf("Error when get place by id: %v\n", err)
			continue
		}
		result = append(result, iomodel.GetCommentResp{
			DataRating: bookingRate,
			DataUser:   *user,
			DataPlace:  *place,
		})
	}

	return result, nil
}

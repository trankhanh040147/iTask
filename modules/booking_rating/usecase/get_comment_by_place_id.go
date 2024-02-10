package bookingratingusecase

import (
	"context"
	"log"
	"paradise-booking/common"
	"paradise-booking/entities"
	"paradise-booking/modules/booking_rating/iomodel"
)

func (uc *bookingRatingUsecase) GetCommentByPlaceID(ctx context.Context, placeID int) (*iomodel.GetCommentByPlaceResp, error) {
	res, err := uc.BookingRatingSto.GetByCondition(ctx, map[string]interface{}{"place_id": placeID})
	if err != nil {
		return nil, common.ErrCannotGetEntity(entities.BookingRating{}.TableName(), err)
	}

	var result iomodel.GetCommentByPlaceResp
	var listRating []iomodel.GetCommentRespByPlace
	for _, bookingRate := range res {
		user, err := uc.AccountSto.GetProfileByID(ctx, bookingRate.UserId)
		if err != nil {
			log.Printf("Error when get user profile by id: %v\n", err)
			continue
		}

		listRating = append(listRating, iomodel.GetCommentRespByPlace{
			DataRating: bookingRate,
			DataUser:   *user,
		})
	}
	result.ListRating = listRating

	place, err := uc.PlaceSto.GetPlaceByID(ctx, placeID)
	if err != nil {
		log.Printf("Error when get place by id: %v\n", err)
		return nil, common.ErrCannotGetEntity(entities.Place{}.TableName(), err)
	}
	result.DataPlace = *place

	return &result, nil
}

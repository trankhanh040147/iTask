package bookingratingusecase

import (
	"context"
	"log"
	"iTask/common"
	"iTask/entities"
	"iTask/modules/booking_rating/iomodel"
)

func (uc *bookingRatingUsecase) GetCommentByVendorID(ctx context.Context, vendorID int) (*iomodel.GetCommentByVendorResp, error) {
	res, err := uc.BookingRatingSto.GetByVendorID(ctx, vendorID)
	if err != nil {
		return nil, common.ErrCannotGetEntity(entities.BookingRating{}.TableName(), err)
	}

	var result iomodel.GetCommentByVendorResp
	for _, bookingRate := range res {
		place, err := uc.PlaceSto.GetPlaceByID(ctx, bookingRate.PlaceId)
		if err != nil {
			log.Printf("Error when get place by id: %v\n", err)
			continue
		}

		user, err := uc.AccountSto.GetProfileByID(ctx, bookingRate.UserId)
		if err != nil {
			log.Printf("Error when get user profile by id: %v\n", err)
			return nil, common.ErrCannotGetEntity(entities.Account{}.TableName(), err)
		}

		result.ListRating = append(result.ListRating, iomodel.GetCommentUserByVendor{
			DataRating: bookingRate,
			DataPlace:  *place,
			DataUser:   *user,
		})
	}

	// dataVendor, err := uc.AccountSto.GetProfileByID(ctx, vendorID)
	// if err != nil {
	// 	return nil, common.ErrCannotGetEntity(entities.Account{}.TableName(), err)
	// }

	// result.DataVendor = dataVendor

	return &result, nil
}

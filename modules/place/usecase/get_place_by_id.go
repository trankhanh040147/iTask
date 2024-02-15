package placeusecase

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/modules/place/convert"
	"paradise-booking/modules/place/iomodel"
)

func (uc *placeUseCase) GetPlaceByID(ctx context.Context, placeID int, userEmail string) (result *iomodel.GetPlaceResp, err error) {
	place, err := uc.placeStorage.GetPlaceByID(ctx, placeID)
	if err != nil {
		return nil, err
	}

	if place == nil {
		return nil, common.ErrEntityNotFound("place", err)
	}

	isFree := true

	userID := 0
	if userEmail != "" {
		user, err := uc.accountSto.GetAccountByEmail(ctx, userEmail)
		if err != nil {
			return nil, err
		}
		userID = user.Id
		// user, err := uc.accountSto.GetAccountByEmail(ctx, userEmail)
		// if err != nil {
		// 	return nil, err
		// }

		placeWishList, err := uc.placeWishSto.GetByCondition(ctx, map[string]interface{}{"user_id": userID, "place_id": place.Id})
		if err != nil {
			return nil, err
		}

		if len(placeWishList) > 0 {
			isFree = false
		}
	}

	// get rating average
	ratingAverage, err := uc.placeStorage.GetRatingAverageByPlaceId(ctx, int64(placeID))
	if err != nil {
		return nil, err
	}

	if ratingAverage == nil {
		defaulRating := 0.0
		ratingAverage = &defaulRating
	}

	result = convert.ConvertPlaceEntityToGetModel(place, isFree, ratingAverage)
	return result, nil
}

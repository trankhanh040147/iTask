package placeusecase

import (
	"context"
	"log"
	"iTask/common"
	"iTask/modules/place/convert"
	"iTask/modules/place/iomodel"
)

func (uc *placeUseCase) ListPlaceByVendorByID(ctx context.Context, vendorID int, paging *common.Paging) (result []iomodel.GetPlaceResp, err error) {

	paging.Process()
	// get places by vendorID
	places, err := uc.placeStorage.ListPlaceByVendorID(ctx, vendorID, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity("place", err)
	}

	if len(places) == 0 {
		log.Printf("Not found any place by vendorID: %d", vendorID)
	}

	// convert data to iomodel
	if len(places) == 0 {
		return []iomodel.GetPlaceResp{}, nil
	}

	for _, place := range places {
		ratingAverage, err := uc.placeStoCache.GetRatingAverageByPlaceId(ctx, int64(place.Id))
		if err != nil {
			return nil, err
		}

		if ratingAverage == nil {
			defaulRating := 0.0
			ratingAverage = &defaulRating
		}

		result = append(result, *convert.ConvertPlaceEntityToGetModel(&place, false, ratingAverage))
	}
	return result, nil
}

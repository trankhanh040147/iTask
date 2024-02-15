package placeusecase

import (
	"context"
	"fmt"
	"log"
	"iTask/common"
	"iTask/constant"
	"iTask/entities"
	"iTask/modules/place/convert"
	"iTask/modules/place/iomodel"
	googlemapprovider "iTask/provider/googlemap"
	"iTask/utils"

	"github.com/samber/lo"
)

func (uc *placeUseCase) ListAllPlace(ctx context.Context, paging *common.Paging, filter *iomodel.Filter, userEmail string) (result []iomodel.GetPlaceResp, err error) {

	address := googlemapprovider.GoogleMapAddress{}

	// get geocode to fill country, state, district
	if filter.Lat != nil && filter.Lng != nil {
		lat := *filter.Lat
		lng := *filter.Lng
		address1, err := uc.googleMap.GetAddressFromLatLng(ctx, lat, lng)
		if err != nil {
			log.Printf("Error when get address from lat lng: %v", err)
			addr, err := uc.placeStorage.GetPlaceByCondition(ctx, map[string]interface{}{"lat": lat, "lng": lng})
			if err != nil {
				return nil, err
			}

			if len(addr) == 0 {
				return nil, fmt.Errorf("cannot get address from lat %v lng %v", lat, lng)
			}

			if len(addr) > 0 {
				address.Country = addr[0].Country
				address.State = addr[0].State
				address.District = addr[0].District
			}
		} else {
			address = *address1
		}
	}

	paging.Process()
	places, err := uc.placeStorage.ListPlaces(ctx, paging, filter, &address)
	if err != nil {
		return nil, common.ErrCannotListEntity("place", err)
	}

	if filter.DateFrom != nil && filter.DateTo != nil {
		// filter by date range
		timeFrom, _ := utils.ParseStringToTime(*filter.DateFrom)
		timeTo, _ := utils.ParseStringToTime(*filter.DateTo)
		bookeds, err := uc.bookingSto.GetBookingsWithinDateRange(ctx, timeFrom, timeTo)
		if err != nil {
			return nil, err
		}

		mapNumPlaceWithPlaceID := make(map[int]int)
		for _, booked := range bookeds {
			if booked.StatusId == constant.BookingStatusCancel || booked.StatusId == constant.BookingStatusCompleted {
				continue
			}
			mapNumPlaceWithPlaceID[booked.PlaceId] += 1
		}

		places = lo.Filter(places, func(item entities.Place, index int) bool {
			return item.NumPlaceOriginal-mapNumPlaceWithPlaceID[item.Id] > 0
		})

		paging.Total = int64(len(places))
	}

	userID := 0
	if userEmail != "" {
		user, err := uc.accountSto.GetAccountByEmail(ctx, userEmail)
		if err != nil {
			return nil, err
		}
		userID = user.Id
	}

	// convert data to iomodel
	for _, place := range places {
		isFree := true

		if userID != 0 {
			placeWishList, err := uc.placeWishSto.GetByCondition(ctx, map[string]interface{}{"user_id": userID, "place_id": place.Id})
			if err != nil {
				return nil, err
			}

			if len(placeWishList) > 0 {
				isFree = false
			}
		}

		// get rating average
		ratingAverage, err := uc.placeStoCache.GetRatingAverageByPlaceId(ctx, int64(place.Id))
		if err != nil {
			return nil, err
		}

		if ratingAverage == nil {
			defaulRating := 0.0
			ratingAverage = &defaulRating
		}

		result = append(result, *convert.ConvertPlaceEntityToGetModel(&place, isFree, ratingAverage))
	}
	return result, nil
}

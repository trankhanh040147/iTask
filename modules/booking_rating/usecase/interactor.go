package bookingratingusecase

import (
	"context"
	"paradise-booking/entities"
)

type BookingRatingSto interface {
	Create(ctx context.Context, data *entities.BookingRating) (*entities.BookingRating, error)
	GetByCondition(ctx context.Context, condition map[string]interface{}) ([]entities.BookingRating, error)
	GetByVendorID(ctx context.Context, vendorID int) ([]entities.BookingRating, error)
	GetStatisticByPlaceID(ctx context.Context, placeId int64) ([]entities.StatisticResp, error)
}

type AccountSto interface {
	GetProfileByID(ctx context.Context, id int) (*entities.Account, error)
}

type PlaceSto interface {
	GetPlaceByID(ctx context.Context, id int) (*entities.Place, error)
}

type Cache interface {
	Delete(ctx context.Context, key string)
}

type bookingRatingUsecase struct {
	BookingRatingSto BookingRatingSto
	AccountSto       AccountSto
	PlaceSto         PlaceSto
	cache            Cache
}

func Newbookingratingusecase(BookingRatingSto BookingRatingSto, accountSto AccountSto, placeSto PlaceSto, cache Cache) *bookingRatingUsecase {
	return &bookingRatingUsecase{BookingRatingSto, accountSto, placeSto, cache}
}

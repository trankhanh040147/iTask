package placewishlistusecase

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/entities"
	"time"
)

type PlaceWishListSto interface {
	Create(ctx context.Context, data *entities.PlaceWishList) error
	Delete(ctx context.Context, place_id, wish_list_id int) error
	GetPlaceIDs(ctx context.Context, wish_list_id int, paging *common.Paging, userID int) ([]int, error)
	GetByCondition(ctx context.Context, condition map[string]interface{}) ([]entities.PlaceWishList, error)
}

type PlaceSto interface {
	ListPlaceInIDs(ctx context.Context, placeIds []int) ([]entities.Place, error)
}

type Cache interface {
	Set(ctx context.Context, key string, value any, ttl time.Duration) error
	Get(ctx context.Context, key string, value any) error
	Delete(ctx context.Context, key string)
}

type placeWishListUsecase struct {
	placeWishListSto PlaceWishListSto
	placeSto         PlaceSto
	cache            Cache
}

func NewPlaceWishListUseCase(PlaceWishListSto PlaceWishListSto, placeSto PlaceSto, cache Cache) *placeWishListUsecase {
	return &placeWishListUsecase{
		placeWishListSto: PlaceWishListSto,
		placeSto:         placeSto,
		cache:            cache,
	}
}

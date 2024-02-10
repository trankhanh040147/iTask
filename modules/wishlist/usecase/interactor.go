package wishlistusecase

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/entities"
	"time"
)

type WishListSto interface {
	Create(ctx context.Context, data *entities.WishList) error
	GetByID(ctx context.Context, id int) (*entities.WishList, error)
	GetByUserID(ctx context.Context, userId int, paging *common.Paging) ([]entities.WishList, error)
	UpdateByID(ctx context.Context, id int, data *entities.WishList) error
	DeleteByID(ctx context.Context, id int) error
}
type Cache interface {
	Set(ctx context.Context, key string, value any, ttl time.Duration) error
	Get(ctx context.Context, key string, value any) error
	Delete(ctx context.Context, key string)
}

type PlaceWishListSto interface {
	GetByCondition(ctx context.Context, condition map[string]interface{}) ([]entities.PlaceWishList, error)
	DeleteByCondition(ctx context.Context, condition []common.Condition) error
}

type wishListUsecase struct {
	wishListSto      WishListSto
	placeWishListSto PlaceWishListSto
	cacheStore       Cache
}

func NewWishListUseCase(wishListSto WishListSto, placeWishListSto PlaceWishListSto, cache Cache) *wishListUsecase {
	return &wishListUsecase{wishListSto: wishListSto, placeWishListSto: placeWishListSto, cacheStore: cache}
}

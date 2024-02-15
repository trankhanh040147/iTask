package placewishlistusecase

import (
	"context"
	"paradise-booking/entities"
	"paradise-booking/modules/place_wishlist/iomodel"
	"time"
)

func (placeWishListUsecase *placeWishListUsecase) CreatePlaceWishList(ctx context.Context, data *iomodel.CreatePlaceWishListReq, userID int) (*entities.PlaceWishList, error) {
	model := entities.PlaceWishList{
		PlaceId:    data.PlaceID,
		WishListId: data.WishListID,
		UserId:     userID,
	}

	if err := placeWishListUsecase.placeWishListSto.Create(ctx, &model); err != nil {
		return nil, err
	}

	// set cache
	key := model.CacheKey()
	err := placeWishListUsecase.cache.Set(ctx, key, model, 24*time.Hour)
	if err != nil {
		return nil, err
	}

	return &model, nil
}

package wishlistusecase

import (
	"context"
	"paradise-booking/entities"
	wishlistiomodel "paradise-booking/modules/wishlist/iomodel"
)

func (wishListUsecase *wishListUsecase) CreateWishList(ctx context.Context, data *wishlistiomodel.CreateWishListReq) (*entities.WishList, error) {
	model := entities.WishList{
		UserID: data.UserID,
		Title:  data.Title,
	}
	if err := wishListUsecase.wishListSto.Create(ctx, &model); err != nil {
		return nil, err
	}

	return &model, nil
}

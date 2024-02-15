package wishlisthandler

import (
	"context"
	"iTask/common"
	"iTask/entities"
	wishlistiomodel "iTask/modules/wishlist/iomodel"
)

type wishListUseCase interface {
	CreateWishList(ctx context.Context, data *wishlistiomodel.CreateWishListReq) (*entities.WishList, error)
	GetWishListByID(ctx context.Context, id int) (*entities.WishList, error)
	GetWishListByUserID(ctx context.Context, userId int, paging *common.Paging) ([]entities.WishList, error)
	UpdateByID(ctx context.Context, id int, title string) error
	DeleteByID(ctx context.Context, id int) error
}

type wishListHandler struct {
	wishListUC wishListUseCase
}

func NewWishListHandler(wishListUC wishListUseCase) *wishListHandler {
	return &wishListHandler{wishListUC: wishListUC}
}

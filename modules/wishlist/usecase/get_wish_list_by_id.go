package wishlistusecase

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/entities"
)

func (uc *wishListUsecase) GetWishListByID(ctx context.Context, id int) (*entities.WishList, error) {
	res, err := uc.wishListSto.GetByID(ctx, id)
	if err != nil {
		return nil, common.ErrCannotGetEntity(entities.WishList{}.TableName(), err)
	}

	return res, nil
}

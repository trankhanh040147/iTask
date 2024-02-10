package wishlistusecase

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/entities"
)

func (uc *wishListUsecase) GetWishListByUserID(ctx context.Context, userId int, paging *common.Paging) ([]entities.WishList, error) {
	paging.Process()
	res, err := uc.wishListSto.GetByUserID(ctx, userId, paging)
	if err != nil {
		return nil, common.ErrCannotGetEntity(entities.WishList{}.TableName(), err)
	}

	return res, nil
}

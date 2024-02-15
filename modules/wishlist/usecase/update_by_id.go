package wishlistusecase

import (
	"context"
	"iTask/common"
	"iTask/entities"
)

func (uc *wishListUsecase) UpdateByID(ctx context.Context, id int, title string) error {
	res, err := uc.wishListSto.GetByID(ctx, id)
	if err != nil {
		return common.ErrCannotGetEntity(entities.WishList{}.TableName(), err)
	}

	// check exist by id
	if res == nil {
		return common.ErrEntityNotFound(entities.WishList{}.TableName(), nil)
	}

	// update
	res.Title = title
	if err := uc.wishListSto.UpdateByID(ctx, id, res); err != nil {
		return common.ErrCannotUpdateEntity(entities.WishList{}.TableName(), err)
	}

	return nil
}

package placewishlistusecase

import (
	"context"
	"iTask/common"
)

func (uc *placeWishListUsecase) DeletePlaceWishList(ctx context.Context, placeId, wishListID int) error {

	// get wish list by wish list id and place id
	data, err := uc.placeWishListSto.GetByCondition(ctx, map[string]interface{}{
		"place_id":    placeId,
		"wishlist_id": wishListID,
	})

	if err != nil {
		return err
	}

	if err := uc.placeWishListSto.Delete(ctx, placeId, wishListID); err != nil {
		return common.ErrCannotDeleteEntity("place_wish_list", err)
	}

	// delete cache
	key := data[0].CacheKey()
	uc.cache.Delete(ctx, key)

	return nil
}

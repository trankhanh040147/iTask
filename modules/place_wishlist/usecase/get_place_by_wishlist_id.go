package placewishlistusecase

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/entities"
)

func (uc *placeWishListUsecase) GetPlaceByWishListID(ctx context.Context, wishListID int, paging *common.Paging, userID int) ([]entities.Place, error) {
	paging.Process()

	// get list placeIDS by wishListID
	placeIDs, err := uc.placeWishListSto.GetPlaceIDs(ctx, wishListID, paging, userID)
	if err != nil {
		return nil, err
	}

	places, err := uc.placeSto.ListPlaceInIDs(ctx, placeIDs)
	if err != nil {
		return nil, err
	}
	return places, nil

}

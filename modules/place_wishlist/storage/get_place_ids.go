package placewishliststorage

import (
	"context"
	"iTask/common"
	"iTask/entities"
)

func (s *placeWishListStorage) GetPlaceIDs(ctx context.Context, wish_list_id int, paging *common.Paging, userId int) ([]int, error) {
	db := s.db
	var res []int

	db = db.Model(entities.PlaceWishList{})

	db = db.Where("wishlist_id = ? AND user_id = ?", wish_list_id, userId)

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrorDB(err)
	}

	if err := db.Offset((paging.Page-1)*paging.Limit).Limit(paging.Limit).Pluck("place_id", &res).Error; err != nil {
		return nil, common.ErrorDB(err)
	}

	return res, nil
}

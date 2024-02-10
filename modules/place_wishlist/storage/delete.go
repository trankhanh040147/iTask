package placewishliststorage

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/entities"
)

func (s *placeWishListStorage) Delete(ctx context.Context, place_id, wish_list_id int) error {
	db := s.db
	if err := db.Where("place_id = ? and wishlist_id = ?", place_id, wish_list_id).Delete(&entities.PlaceWishList{}).Error; err != nil {
		return common.ErrorDB(err)
	}
	return nil
}

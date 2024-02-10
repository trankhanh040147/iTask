package placewishliststorage

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/entities"
)

func (s *placeWishListStorage) DeleteByCondition(ctx context.Context, condition []common.Condition) error {
	db := s.db

	for _, v := range condition {
		query := v.BuildQuery()
		db = db.Where(query+" ?", v.Value)
	}

	if err := db.Delete(&entities.PlaceWishList{}).Error; err != nil {
		return err
	}

	return nil
}

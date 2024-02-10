package wishliststorage

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/entities"
)

func (s *wishListStorage) DeleteByID(ctx context.Context, id int) error {
	db := s.db

	if err := db.Where("id = ?", id).Delete(entities.WishList{}).Error; err != nil {
		return common.ErrorDB(err)
	}

	return nil
}

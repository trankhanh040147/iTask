package wishliststorage

import (
	"context"
	"iTask/common"
	"iTask/entities"
)

func (s *wishListStorage) UpdateByID(ctx context.Context, id int, data *entities.WishList) error {
	db := s.db

	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrorDB(err)
	}

	return nil
}

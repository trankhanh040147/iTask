package wishliststorage

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/entities"

	"gorm.io/gorm"
)

func (s *wishListStorage) GetByID(ctx context.Context, id int) (*entities.WishList, error) {
	db := s.db

	var data entities.WishList

	if err := db.Where("id = ?", id).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, common.ErrorDB(err)
	}

	return &data, nil
}

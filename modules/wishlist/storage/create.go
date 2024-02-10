package wishliststorage

import (
	"context"
	"paradise-booking/entities"
)

func (s *wishListStorage) Create(ctx context.Context, data *entities.WishList) error {
	db := s.db

	if err := db.Create(&data).Error; err != nil {
		return err
	}

	return nil
}

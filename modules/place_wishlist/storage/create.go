package placewishliststorage

import (
	"context"
	"paradise-booking/entities"
)

func (s *placeWishListStorage) Create(ctx context.Context, data *entities.PlaceWishList) error {

	if err := s.db.Create(data).Error; err != nil {
		return err
	}
	return nil
}

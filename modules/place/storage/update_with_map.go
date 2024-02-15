package placestorage

import (
	"context"
	"iTask/entities"
)

func (s *placeStorage) UpdateWithMap(ctx context.Context, place *entities.Place, props map[string]interface{}) error {
	db := s.db

	if err := db.Model(place).Updates(props).Error; err != nil {
		return err
	}
	return nil
}

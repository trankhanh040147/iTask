package placestorage

import (
	"context"
	"iTask/common"
	"iTask/entities"
)

func (s *placeStorage) UpdateByID(ctx context.Context, id int, data *entities.Place) error {
	db := s.db

	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrorDB(err)
	}

	return nil
}

package placestorage

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/entities"
)

func (s *placeStorage) DeleteByID(ctx context.Context, id int) error {
	db := s.db

	if err := db.Where("id = ?", id).Delete(&entities.Place{}).Error; err != nil {
		return common.ErrorDB(err)
	}

	return nil
}

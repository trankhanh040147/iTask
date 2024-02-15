package placestorage

import (
	"context"
	"iTask/common"
	"iTask/entities"
)

func (s *placeStorage) Create(ctx context.Context, data *entities.Place) (err error) {
	db := s.db.Begin()

	if err := db.Create(data).Error; err != nil {
		db.Rollback()
		return common.ErrorDB(err)
	}

	db.Commit()
	return nil
}

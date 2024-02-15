package placestorage

import (
	"context"
	"iTask/common"
	"iTask/entities"
)

func (s *placeStorage) ListPlaceByCondition(ctx context.Context, condition []common.Condition) ([]entities.Place, error) {
	var data []entities.Place

	db := s.db
	db = db.Table(entities.Place{}.TableName())

	for _, v := range condition {
		query := v.BuildQuery()
		db = db.Where(query+" ?", v.Value)
	}

	if err := db.Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

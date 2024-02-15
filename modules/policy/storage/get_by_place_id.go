package policiesstorage

import (
	"context"
	"iTask/entities"
)

func (s *policyStorage) GetByPlaceID(ctx context.Context, placeId int) ([]entities.Policy, error) {
	db := s.db.WithContext(ctx)

	var data []entities.Policy
	err := db.Table(entities.Policy{}.TableName()).Where("place_id = ?", placeId).Find(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

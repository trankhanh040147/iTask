package policiesstorage

import (
	"context"
	"iTask/entities"
)

func (s *policyStorage) GetByCondition(ctx context.Context, condition map[string]any) ([]entities.Policy, error) {
	db := s.db.WithContext(ctx)

	var data []entities.Policy
	err := db.Table(entities.Policy{}.TableName()).Where(condition).Find(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

package policiesstorage

import (
	"context"
	"paradise-booking/entities"
)

func (s *policyStorage) Update(ctx context.Context, data *entities.Policy) error {
	db := s.db.WithContext(ctx)
	err := db.Model(data).Updates(data).Error
	if err != nil {
		return err
	}
	return nil
}

package policiesstorage

import (
	"context"
	"paradise-booking/entities"
)

func (s *policyStorage) Create(ctx context.Context, data *entities.Policy) error {
	db := s.db.WithContext(ctx).Create(data)
	if err := db.Error; err != nil {
		return err
	}
	return nil
}

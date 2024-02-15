package amenitystorage

import (
	"context"
	"paradise-booking/entities"
)

func (s *amenityStorage) DeleteByCondition(ctx context.Context, condition map[string]any) error {
	db := s.db.Model(&entities.Amenity{}).Where(condition)

	err := db.Delete(&entities.Amenity{}).Error
	if err != nil {
		return err
	}

	return nil
}

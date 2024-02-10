package amenitystorage

import (
	"context"
	"paradise-booking/entities"
)

func (s *amenityStorage) GetAllAmenityConfig(ctx context.Context) ([]entities.ConfigAmenity, error) {
	var res []entities.ConfigAmenity
	db := s.db.Table(entities.ConfigAmenity{}.TableName())
	if err := db.Where("1=1").Find(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}

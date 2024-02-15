package amenitystorage

import (
	"context"
	"iTask/entities"
)

func (s *amenityStorage) Create(ctx context.Context, data *entities.Amenity) (res *entities.Amenity, err error) {
	db := s.db.Table(data.TableName())
	if err = db.Create(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

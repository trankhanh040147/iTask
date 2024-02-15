package amenitystorage

import (
	"context"
	"iTask/entities"
)

func (s *amenityStorage) ListByPlaceID(ctx context.Context, placeID int) ([]entities.Amenity, error) {
	var res []entities.Amenity
	db := s.db.Table(entities.Amenity{}.TableName())
	if err := db.Where("place_id = ?", placeID).Find(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}

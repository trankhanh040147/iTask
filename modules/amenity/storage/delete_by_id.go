package amenitystorage

import (
	"context"
	"iTask/common"
	"iTask/entities"
)

func (s *amenityStorage) Delete(ctx context.Context, id int) error {
	db := s.db.Table(entities.Amenity{}.TableName()).Where("id = ?", id)
	if err := db.Delete(&entities.Amenity{}).Error; err != nil {
		return common.ErrorDB(err)
	}

	return nil
}

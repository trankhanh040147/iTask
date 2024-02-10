package placestorage

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/entities"
)

func (s *placeStorage) ListPlaceByVendorID(ctx context.Context, vendorID int, paging *common.Paging) ([]entities.Place, error) {
	db := s.db

	var data []entities.Place

	db = db.Table(entities.Place{}.TableName()).Where("vendor_id = ?", vendorID)
	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrorDB(err)
	}

	if err := db.Offset((paging.Page - 1) * paging.Limit).Limit(paging.Limit).Find(&data).Error; err != nil {
		return nil, common.ErrorDB(err)
	}

	return data, nil
}

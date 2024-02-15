package placestorage

import (
	"context"
	"iTask/common"
	"iTask/entities"
	"iTask/modules/place/iomodel"
	googlemapprovider "iTask/provider/googlemap"
)

func (s *placeStorage) ListPlaces(ctx context.Context, paging *common.Paging, filter *iomodel.Filter, address *googlemapprovider.GoogleMapAddress) ([]entities.Place, error) {
	db := s.db

	var data []entities.Place

	db = db.Table(entities.Place{}.TableName())

	if v := filter; v != nil {
		if v.VendorID != nil {
			db = db.Where("vendor_id = ?", v.VendorID)
		}

		if v.Lat != nil && v.Lng != nil {
			if address.Country != "" {
				db = db.Where("country = ?", address.Country)
			}
			if address.State != "" {
				db = db.Where("state = ?", address.State)
			}
			if address.District != "" {
				db = db.Where("district = ?", address.District)
			}
		}

		if v.Guest != nil {
			db = db.Where("max_guest >= ?", v.Guest)
		}

		if v.PriceFrom != nil {
			db = db.Where("price_per_night >= ?", v.PriceFrom)
		}
		if v.PriceTo != nil {
			db = db.Where("price_per_night <= ?", v.PriceTo)
		}

		if v.NumBed != nil {
			db = db.Where("num_bed >= ?", v.NumBed)
		}

		if v.Bedroom != nil {
			db = db.Where("bed_room >= ?", v.Bedroom)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrorDB(err)
	}

	if err := db.Offset((paging.Page - 1) * paging.Limit).Limit(paging.Limit).Find(&data).Error; err != nil {
		return nil, common.ErrorDB(err)
	}

	return data, nil
}

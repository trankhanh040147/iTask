package convert

import (
	"iTask/entities"
	"iTask/modules/place/iomodel"
)

func ConvertPlaceEntityToGetModel(data *entities.Place, isFree bool, ratingAverage *float64) *iomodel.GetPlaceResp {
	return &iomodel.GetPlaceResp{
		ID:            data.Id,
		VendorID:      data.VendorID,
		Name:          data.Name,
		Description:   data.Description,
		PricePerNight: data.PricePerNight,
		Address:       data.Address,
		Cover:         data.Cover,
		Lat:           data.Lat,
		Lng:           data.Lng,
		Country:       data.Country,
		State:         data.State,
		District:      data.District,
		MaxGuest:      data.MaxGuest,
		Numbed:        data.NumBed,
		IsFree:        isFree,
		RatingAverage: *ratingAverage,
		BedRoom:       data.BedRoom,
	}
}

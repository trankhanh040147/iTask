package convert

import (
	"iTask/entities"
	"iTask/modules/place/iomodel"
)

func ConvertPlaceCreateModelToEntity(data *iomodel.CreatePlaceReq) *entities.Place {
	return &entities.Place{
		Name:             data.Name,
		Description:      data.Description,
		PricePerNight:    data.PricePerNight,
		Address:          data.Address,
		Cover:            data.Cover,
		Lat:              data.Lat,
		Lng:              data.Lng,
		MaxGuest:         data.MaxGuest,
		NumBed:           data.NumBed,
		BedRoom:          data.BedRoom,
		NumPlaceOriginal: data.NumPlaceOriginal,
	}
}

func ConvertPlaceUpdateModelToEntity(data *iomodel.UpdatePlaceReq) *entities.Place {
	return &entities.Place{
		Name:             data.Name,
		Description:      data.Description,
		PricePerNight:    data.PricePerNight,
		Address:          data.Address,
		Cover:            data.Cover,
		Lat:              data.Lat,
		Lng:              data.Lng,
		Country:          data.Country,
		State:            data.State,
		District:         data.District,
		MaxGuest:         data.MaxGuest,
		NumBed:           data.NumBed,
		BedRoom:          data.BedRoom,
		NumPlaceOriginal: data.NumPlaceOriginal,
	}
}

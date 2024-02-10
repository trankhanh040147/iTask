package placeusecase

import (
	"context"
	"paradise-booking/modules/place/convert"
	"paradise-booking/modules/place/iomodel"
)

func (uc *placeUseCase) CreatePlace(ctx context.Context, data *iomodel.CreatePlaceReq, emailVendor string) error {
	// convert iomodel to entity
	placeEntity := convert.ConvertPlaceCreateModelToEntity(data)

	// check vendor exist
	vendor, err := uc.accountSto.GetAccountByEmail(ctx, emailVendor)
	if err != nil {
		return err
	}

	placeEntity.VendorID = vendor.Id

	// get geocode to fill country, state, district
	adress, err := uc.googleMap.GetAddressFromLatLng(ctx, data.Lat, data.Lng)
	if err != nil {
		return err
	}

	placeEntity.Country = adress.Country
	placeEntity.State = adress.State
	placeEntity.District = adress.District
	// create place
	if err := uc.placeStorage.Create(ctx, placeEntity); err != nil {
		return err
	}
	return nil
}

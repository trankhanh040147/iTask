package placeusecase

import (
	"context"
	"iTask/common"
	"iTask/modules/place/convert"
	"iTask/modules/place/iomodel"
)

func (uc *placeUseCase) ListPlaceByVendor(ctx context.Context, vendorEmail string) (result []iomodel.GetPlaceResp, err error) {

	// get vendorID from vendorEmail
	vendor, err := uc.accountSto.GetAccountByEmail(ctx, vendorEmail)
	if err != nil {
		return nil, common.ErrCannotGetEntity("account", err)
	}

	var paging common.Paging
	paging.Process()
	// get places by vendorID
	places, err := uc.placeStorage.ListPlaceByVendorID(ctx, vendor.Id, &paging)
	if err != nil {
		return nil, common.ErrCannotListEntity("place", err)
	}

	if len(places) == 0 {
		return []iomodel.GetPlaceResp{}, nil
	}

	// convert data to iomodel
	defaulRating := 0.0

	for _, place := range places {
		result = append(result, *convert.ConvertPlaceEntityToGetModel(&place, false, &defaulRating))
	}
	return result, nil
}

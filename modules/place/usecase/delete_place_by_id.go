package placeusecase

import (
	"context"
	"errors"
	"iTask/common"
)

func (uc *placeUseCase) DeletePlaceByID(ctx context.Context, placeID int, vendorEmail string) (err error) {

	// get place to check owner
	place, err := uc.placeStorage.GetPlaceByID(ctx, placeID)
	if err != nil {
		return common.ErrCannotGetEntity("Place", err)
	}

	// get vendor to check owner
	vendor, err := uc.accountSto.GetAccountByEmail(ctx, vendorEmail)
	if err != nil {
		return common.ErrCannotGetEntity("Vendor", err)
	}

	// check owner
	if place.VendorID != vendor.Id {
		return common.ErrNotOwner(errors.New("you re not owner of this place"))
	}

	err = uc.placeStorage.DeleteByID(ctx, placeID)
	if err != nil {
		return common.ErrCannotDeleteEntity("Place", err)
	}
	return nil
}

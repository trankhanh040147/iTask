package placeusecase

import (
	"context"
	"errors"
	"iTask/common"
	"iTask/modules/place/convert"
	"iTask/modules/place/iomodel"
)

func (uc *placeUseCase) UpdatePlace(ctx context.Context, data *iomodel.UpdatePlaceReq, placeID int, vendorEmail string) error {
	// convert iomodel to entity
	placeEntity := convert.ConvertPlaceUpdateModelToEntity(data)

	// check vendor exist
	vendor, err := uc.accountSto.GetAccountByEmail(ctx, vendorEmail)
	if err != nil {
		return err
	}

	// get place by id to check if vendor is owner of this place
	place, err := uc.placeStorage.GetPlaceByID(ctx, placeID)

	// check exist before update
	if place == nil && err == nil {
		return common.ErrEntityNotFound(placeEntity.TableName(), err)
	}

	if err != nil {
		return common.ErrCannotGetEntity(placeEntity.TableName(), err)
	}

	if place.VendorID != vendor.Id {
		return common.ErrNotOwner(errors.New("you are not owner of this place"))
	}

	// update place
	if err := uc.placeStorage.UpdateByID(ctx, place.Id, placeEntity); err != nil {
		return err
	}

	return nil
}

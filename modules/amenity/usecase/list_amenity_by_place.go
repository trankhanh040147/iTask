package amenityusecase

import (
	"context"
	"iTask/entities"
)

func (uc *amenityUseCase) ListAmenityByPlaceID(ctx context.Context, placeID int) (res []entities.Amenity, err error) {

	res, err = uc.amenitySto.ListByPlaceID(ctx, placeID)
	if err != nil {
		return nil, err
	}
	return res, nil

}

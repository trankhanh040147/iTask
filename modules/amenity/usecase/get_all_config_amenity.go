package amenityusecase

import (
	"context"
	"iTask/entities"
)

func (uc *amenityUseCase) GetAllConfigAmenity(ctx context.Context) (res []entities.ConfigAmenity, err error) {

	res, err = uc.amenitySto.GetAllAmenityConfig(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil

}

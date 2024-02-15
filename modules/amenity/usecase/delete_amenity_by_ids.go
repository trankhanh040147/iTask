package amenityusecase

import (
	"context"
	"iTask/modules/amenity/iomodel"
)

func (u *amenityUseCase) DeleteAmenityByListId(ctx context.Context, req *iomodel.DeleteAmenityReq) error {

	for _, id := range req.ListConfigAmenityId {
		condition := map[string]any{
			"place_id":          req.IDPlace,
			"config_amenity_id": id,
		}
		err := u.amenitySto.DeleteByCondition(ctx, condition)
		if err != nil {
			return err
		}
	}

	return nil
}

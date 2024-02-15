package amenityusecase

import (
	"context"
	"iTask/entities"
	"iTask/modules/amenity/iomodel"
)

func (u *amenityUseCase) CreateAmenity(ctx context.Context, data *iomodel.CreateAmenityReq) (err error) {
	for _, v := range data.ListDetailAmenity {
		data := &entities.Amenity{
			PlaceId:         data.PlaceId,
			Description:     v.Description,
			ConfigAmenityId: v.ConfigAmenityId,
		}
		_, err := u.amenitySto.Create(ctx, data)
		if err != nil {
			return err
		}
	}

	return nil
}

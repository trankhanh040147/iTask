package amenityusecase

import (
	"context"
)

func (u *amenityUseCase) DeleteAmenityById(ctx context.Context, id int) error {
	err := u.amenitySto.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

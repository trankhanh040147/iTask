package policiesusecase

import (
	"context"
	"paradise-booking/entities"
)

func (uc *policyUsecase) GetPolicyByPlaceID(ctx context.Context, placeId int) ([]entities.Policy, error) {

	data, err := uc.PolicyStore.GetByPlaceID(ctx, placeId)
	if err != nil {
		return nil, err
	}

	return data, nil
}

package policiesusecase

import (
	"context"
	"paradise-booking/entities"
	"paradise-booking/modules/policy/iomodel"

	"gorm.io/gorm"
)

func (uc *policyUsecase) UpSearchPolicy(ctx context.Context, dataReq *iomodel.CreatePolicyReq) error {

	for _, policy := range dataReq.Data.ListPolicy {
		data, err := uc.PolicyStore.GetByCondition(ctx, map[string]any{
			"place_id":        dataReq.Data.PlaceID,
			"group_policy_id": policy.GroupPolicyID,
		})

		if err != nil && err != gorm.ErrRecordNotFound {
			return err
		}

		// because place_id and group_policy_id is unique => len(data) must be 0 or 1
		if len(data) == 0 {
			// create new
			record := &entities.Policy{
				PlaceId:       dataReq.Data.PlaceID,
				Name:          policy.Name,
				GroupPolicyId: policy.GroupPolicyID,
			}

			err = uc.PolicyStore.Create(ctx, record)
			if err != nil {
				return err
			}

		} else {
			if data[0].Name == policy.Name {
				continue
			}
			// update
			data[0].Name = policy.Name
			err := uc.PolicyStore.Update(ctx, &data[0])
			if err != nil {
				return err
			}
		}

	}
	return nil
}

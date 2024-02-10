package policiesusecase

import (
	"context"
	"paradise-booking/entities"
)

type PolicyStore interface {
	Create(ctx context.Context, data *entities.Policy) error
	GetByPlaceID(ctx context.Context, placeId int) ([]entities.Policy, error)
	Update(ctx context.Context, data *entities.Policy) error
	GetByCondition(ctx context.Context, condition map[string]any) ([]entities.Policy, error)
}

type policyUsecase struct {
	PolicyStore PolicyStore
}

func NewPolicyUseCase(PolicyStore PolicyStore) *policyUsecase {
	return &policyUsecase{PolicyStore: PolicyStore}
}

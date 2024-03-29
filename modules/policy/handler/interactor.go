package policieshandler

import (
	"context"
	"iTask/entities"
	"iTask/modules/policy/iomodel"
)

type PolicyUseCase interface {
	UpSearchPolicy(ctx context.Context, dataReq *iomodel.CreatePolicyReq) error
	GetPolicyByPlaceID(ctx context.Context, placeId int) ([]entities.Policy, error)
}

type policyHandler struct {
	policyUC PolicyUseCase
}

func NewPolicyHandler(policy PolicyUseCase) *policyHandler {
	return &policyHandler{policyUC: policy}
}

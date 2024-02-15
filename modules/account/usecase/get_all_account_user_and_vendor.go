package accountusecase

import (
	"context"
	"iTask/common"
	"iTask/entities"
)

func (a *accountUseCase) GetAllAccountUserAndVendor(ctx context.Context, paging *common.Paging) ([]entities.Account, error) {
	paging.Process()
	result, err := a.accountStorage.GetAllAccountUserAndVendor(ctx, paging)
	if err != nil {
		return nil, err
	}
	return result, nil
}

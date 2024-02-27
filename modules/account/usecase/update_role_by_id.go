package accountusecase

import (
	"context"
	"iTask/common"
	"iTask/entities"
	"iTask/modules/account/iomodel"
)

func (uc *accountUseCase) UpdateAccountRoleByID(ctx context.Context, accountModel *iomodel.AccountChangeRole, id int) (err error) {

	model := entities.Account{
		Role: entities.UserRole(accountModel.Role),
	}
	err = uc.accountStorage.UpdateAccountById(ctx, id, &model)
	if err != nil {
		return common.ErrInternal(err)
	}
	return nil
}

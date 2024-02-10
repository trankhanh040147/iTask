package accountusecase

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/entities"
	"paradise-booking/modules/account/iomodel"
)

func (uc *accountUseCase) UpdateAccountRoleByID(ctx context.Context, accountModel *iomodel.AccountChangeRole, id int) (err error) {

	model := entities.Account{
		Role: accountModel.Role,
	}
	err = uc.accountStorage.UpdateAccountById(ctx, id, &model)
	if err != nil {
		return common.ErrInternal(err)
	}
	return nil
}

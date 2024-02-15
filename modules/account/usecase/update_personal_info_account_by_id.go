package accountusecase

import (
	"context"
	"iTask/common"
	"iTask/modules/account/convert"
	"iTask/modules/account/iomodel"
)

func (uc *accountUseCase) UpdatePersonalInforAccountById(ctx context.Context, accountModel *iomodel.AccountUpdatePersonalInfo, id int) (err error) {
	// convert data
	accountEntity := convert.ConvertAccountUpdatePersonalInfoModelToEntity(accountModel)

	err = uc.accountStorage.UpdateAccountById(ctx, id, &accountEntity)
	if err != nil {
		return common.ErrInternal(err)
	}
	return nil
}

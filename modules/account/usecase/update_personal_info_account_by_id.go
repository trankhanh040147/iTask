package accountusecase

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/modules/account/convert"
	"paradise-booking/modules/account/iomodel"
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

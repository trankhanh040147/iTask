package accountusecase

import (
	"context"
	"iTask/modules/account/convert"
	"iTask/modules/account/iomodel"
)

func (uc *accountUseCase) GetAccountByID(ctx context.Context, id int) (account *iomodel.AccountInfoResp, err error) {
	accountEntity, err := uc.accountStorage.GetProfileByID(ctx, id)
	if err != nil {
		return nil, err
	}

	account = convert.ConvertAccountEntityToInfoResp(accountEntity)
	return account, nil
}

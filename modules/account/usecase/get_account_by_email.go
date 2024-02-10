package accountusecase

import (
	"context"
	"paradise-booking/modules/account/convert"
	"paradise-booking/modules/account/iomodel"
)

func (uc *accountUseCase) GetAccountByEmail(ctx context.Context, email string) (account *iomodel.AccountInfoResp, err error) {
	accountEntity, err := uc.accountStorage.GetAccountByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	account = convert.ConvertAccountEntityToInfoResp(accountEntity)
	return account, nil
}

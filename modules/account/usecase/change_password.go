package accountusecase

import (
	"context"
	"iTask/common"
	"iTask/entities"
	"iTask/modules/account/iomodel"
	"iTask/utils"
)

func (uc *accountUseCase) ChangePassword(ctx context.Context, email string, changePassModel *iomodel.ChangePassword) error {
	account, err := uc.accountStorage.GetAccountByEmail(ctx, email)
	if err != nil {
		return err
	}

	// check old password is correct
	err = utils.Compare(account.Password, changePassModel.OldPassword)
	if err != nil {
		return common.ErrOldPasswordInvalid(account.TableName(), err)
	}

	// hash password before store in db
	hashedPassword, err := utils.HashPassword(changePassModel.NewPassword)
	if err != nil {
		return common.ErrInternal(err)
	}

	updateAccount := entities.Account{
		Password: hashedPassword,
	}

	if err := uc.accountStorage.UpdateAccountById(ctx, account.Id, &updateAccount); err != nil {
		return err
	}
	return nil
}

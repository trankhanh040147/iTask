package accountusecase

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/entities"
	"paradise-booking/utils"
)

func (uc *accountUseCase) UpdatePassword(ctx context.Context, email string, newPassword string) error {
	account, err := uc.accountStorage.GetAccountByEmail(ctx, email)
	if err != nil {
		return err
	}

	// hash password before store in db
	hashedPassword, err := utils.HashPassword(newPassword)
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

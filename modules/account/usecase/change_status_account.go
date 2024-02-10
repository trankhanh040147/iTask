package accountusecase

import (
	"context"
	"paradise-booking/entities"
)

func (uc *accountUseCase) ChangeStatusAccount(ctx context.Context, accountID int, status int) error {
	updateStatusAccount := entities.Account{
		Status: status,
	}

	err := uc.accountStorage.UpdateAccountById(ctx, accountID, &updateStatusAccount)
	if err != nil {
		return err
	}
	return nil
}

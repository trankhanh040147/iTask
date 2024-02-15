package accountstorage

import (
	"context"
	"iTask/entities"
)

type CreateUserTxParam struct {
	Data        *entities.Account
	AfterCreate func(data *entities.Account) error
}

func (s *accountStorage) CreateTx(ctx context.Context, createUserTxParam CreateUserTxParam) error {

	err := s.execTx(ctx, func(store *accountStorage) error {
		if err := store.Create(ctx, createUserTxParam.Data); err != nil {
			return err
		}

		if createUserTxParam.AfterCreate != nil {
			if err := createUserTxParam.AfterCreate(createUserTxParam.Data); err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

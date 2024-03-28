package storage

import (
	"context"
	"iTask/entities"
)

type InviteMemberTxParam struct {
	Data        *entities.Account
	AfterCreate func(data *entities.Account) error
}

func (s *sqlStore) InviteMemberTx(ctx context.Context, createUserTxParam InviteMemberTxParam) error {

	err := s.execTx(ctx, func(store *sqlStore) error {
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

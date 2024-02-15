package accountstorage

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/entities"
)

func (s *accountStorage) Create(ctx context.Context, account *entities.Account) (err error) {
	db := s.db

	if err = db.Create(account).Error; err != nil {
		return common.ErrorDB(err)
	}
	return nil
}

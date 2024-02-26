package accountstorage

import (
	"context"
	"iTask/common"
	"iTask/entities"
)

func (s *accountStorage) Create(ctx context.Context, account *entities.Account) (err error) {
	db := s.db

	if err = db.Create(account).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}

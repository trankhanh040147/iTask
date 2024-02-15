package accountstorage

import (
	"context"
	"iTask/common"
	"iTask/entities"
)

func (s *accountStorage) UpdateAccountById(ctx context.Context, id int, accountUpdate *entities.Account) error {
	db := s.db
	var account entities.Account
	if err := db.Table(account.TableName()).Where("id = ?", id).Updates(&accountUpdate).Error; err != nil {
		return common.ErrorDB(err)
	}
	return nil
}

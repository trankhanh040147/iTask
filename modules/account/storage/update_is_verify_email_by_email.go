package accountstorage

import (
	"context"
	"iTask/common"
	"iTask/entities"
)

func (s *accountStorage) UpdateIsVerifyEmailByEmail(ctx context.Context, email string) error {
	db := s.db
	account := entities.Account{}
	if err := db.Table(account.TableName()).Where("email = ?", email).Update("is_email_verified", 1).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}

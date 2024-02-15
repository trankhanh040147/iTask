package accountstorage

import (
	"context"
	"iTask/common"
	"iTask/entities"

	"gorm.io/gorm"
)

func (s *accountStorage) GetAccountByEmail(ctx context.Context, email string) (*entities.Account, error) {
	db := s.db
	var account entities.Account
	if err := db.Table(account.TableName()).Where("email = ?", email).First(&account).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, common.ErrorDB(err)
	}
	return &account, nil
}

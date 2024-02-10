package accountstorage

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/entities"

	"gorm.io/gorm"
)

func (s *accountStorage) GetProfileByID(ctx context.Context, id int) (*entities.Account, error) {

	var account entities.Account
	db := s.db.Table(account.TableName())

	if err := db.Where("id = ?", id).First(&account).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}

		return nil, common.ErrorDB(err)
	}

	return &account, nil
}

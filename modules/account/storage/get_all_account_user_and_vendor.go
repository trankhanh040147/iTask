package accountstorage

import (
	"context"
	"iTask/common"
	"iTask/constant"
	"iTask/entities"
)

func (s *accountStorage) GetAllAccountUserAndVendor(ctx context.Context, paging *common.Paging) ([]entities.Account, error) {
	var result []entities.Account
	db := s.db.Table(entities.Account{}.TableName())
	db = db.Where("role = ? OR role = ?", constant.UserRole, constant.VendorRole)

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if err := db.Offset((paging.Page - 1) * paging.Limit).Limit(paging.Limit).Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil

}

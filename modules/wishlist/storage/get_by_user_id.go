package wishliststorage

import (
	"context"
	"iTask/common"
	"iTask/entities"
)

func (s *wishListStorage) GetByUserID(ctx context.Context, userId int, paging *common.Paging) ([]entities.WishList, error) {
	db := s.db

	var data []entities.WishList

	db = db.Table(entities.WishList{}.TableName()).Where("user_id = ?", userId)

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrorDB(err)
	}

	if err := db.Offset((paging.Page - 1) * paging.Limit).Limit(paging.Limit).Find(&data).Error; err != nil {
		return nil, common.ErrorDB(err)
	}

	return data, nil
}

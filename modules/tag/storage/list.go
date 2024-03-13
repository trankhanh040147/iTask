package storage

import (
	"context"
	"iTask/common"
	"iTask/modules/tag/model"
)

func (s *sqlStore) ListTag(ctx context.Context, filter *model.Filter, paging *common.Paging, moreKeys ...string) ([]model.Tag, error) {
	var result []model.Tag

	db := s.db.Table(model.Tag{}.TableName())

	if err := db.Select("id").Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if err := db.
		Select("*").
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}

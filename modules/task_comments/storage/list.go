package storage

import (
	"context"
	"iTask/common"
	"iTask/modules/task_comments/model"
)

func (s *SQLStore) ListComments(ctx context.Context, cond map[string]interface{}, paging *common.Paging, moreKeys ...string) ([]model.TaskComment, error) {
	s.db = s.db.Table(model.TaskComment{}.TableName())

	if err := s.db.Where(cond).Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for _, key := range moreKeys {
		s.db = s.db.Preload(key)
	}

	var result []model.TaskComment
	if err := s.db.Select("*").
		Where(cond).
		Offset((paging.Page - 1) * paging.Limit).
		Order("id").
		Limit(paging.Limit).
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}

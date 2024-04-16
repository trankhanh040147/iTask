package storage

import (
	"context"
	"iTask/common"
	"iTask/modules/task_comments/model"
)

func (s *SQLStore) GetComment(ctx context.Context, cond map[string]interface{}) (*model.TaskComment, error) {
	var data model.TaskComment

	if err := s.db.Model(&model.TaskComment{}).Where(cond).First(&data).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return &data, nil
}

package storage

import (
	"context"
	"iTask/common"
	"iTask/modules/task_comments/model"
)

func (s *SQLStore) CreateTaskComment(ctx context.Context, data *model.TaskCommentCreation) error {
	if err := s.db.Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

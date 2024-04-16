package storage

import (
	"context"
	"iTask/common"
	"iTask/modules/task_comments/model"
)

func (s *SQLStore) UpdateComment(ctx context.Context, cond map[string]interface{}, dataUpdate *model.TaskCommentUpdate) error {
	if err := s.db.Where(cond).Updates(dataUpdate).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

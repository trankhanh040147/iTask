package storage

import (
	"context"
	"iTask/common"
	"iTask/modules/task_assignees/model"
)

func (s *sqlStore) DeleteAssignee(ctx context.Context, cond map[string]interface{}) error {
	if err := s.db.Where(cond).Delete(model.TaskAssignee{}).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

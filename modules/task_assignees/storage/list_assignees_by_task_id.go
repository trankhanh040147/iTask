package storage

import (
	"context"
	"iTask/common"
	"iTask/modules/task_assignees/model"
)

func (s *sqlStore) ListAssigneeByTaskId(ctx context.Context, taskId int, moreKeys ...string) ([]model.TaskAssignee, error) {
	var result []model.TaskAssignee

	db := s.db.Table(model.TaskAssignee{}.TableName())

	for key := range moreKeys {
		db = db.Preload(moreKeys[key])
	}

	db = db.Where("task_id = ?", taskId)

	if err := db.
		Select("*").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}

package storage

import (
	"context"
	"iTask/common"
	"iTask/modules/task_assignees/model"
	"log"
)

func (s *sqlStore) ListAssignee(ctx context.Context, filter *model.Filter, moreKeys ...string) ([]model.TaskAssignee, error) {
	var result []model.TaskAssignee

	db := s.db.Table(model.TaskAssignee{}.TableName())

	for key := range moreKeys {
		db = db.Preload(moreKeys[key])
	}

	log.Println("---->filter", filter)

	if filter != nil {
		if filter.TaskID != 0 {
			db = db.Where("task_id = ?", filter.TaskID)
		}
	}

	if err := db.
		Select("*").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}

package storage

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"iTask/common"
	"iTask/modules/task_assignees/model"
)

func (s *sqlStore) GetAssignee(ctx context.Context, cond map[string]interface{}, moreKeys ...string) (*model.TaskAssignee, error) {
	var data *model.TaskAssignee

	if err := s.db.Where(cond).First(&data).Error; err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return data, nil
}

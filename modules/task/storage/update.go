package storage

import (
	"context"
	"iTask/modules/task/model"
)

func (s *sqlStore) UpdateTask(ctx context.Context, cond map[string]interface{}, dataUpdate *model.TaskUpdate) error {
	if err := s.db.Where(cond).Updates(dataUpdate).Error; err != nil {
		return err
	}
	return nil
}

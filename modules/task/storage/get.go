package storage

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"iTask/common"
	"iTask/modules/task/model"
)

func (s *sqlStore) GetTask(ctx context.Context, cond map[string]interface{}, moreKeys ...string) (*model.Task, error) {
	var data *model.Task

	if err := s.db.Where(cond).First(&data).Error; err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			return nil, common.RecordNotFound
		}

		return nil, common.ErrDB(err)
	}

	return data, nil
}

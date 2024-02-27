package storage

import (
	"context"
	"errors"
	"iTask/common"
	"iTask/modules/project/model"

	"gorm.io/gorm"
)

func (s *sqlStore) GetProject(ctx context.Context, cond map[string]interface{}) (*model.Project, error) {
	var data *model.Project

	if err := s.db.Where(cond).First(&data).Error; err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			return nil, common.RecordNotFound
		}

		return nil, common.ErrDB(err)
	}

	return data, nil
}

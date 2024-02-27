package storage

import (
	"context"
	"iTask/modules/project/model"
)

func (s *sqlStore) UpdateProject(ctx context.Context, cond map[string]interface{}, dataUpdate *model.ProjectUpdate) error {
	if err := s.db.Where(cond).Updates(dataUpdate).Error; err != nil {
		return err
	}
	return nil
}

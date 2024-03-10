package storage

import (
	"context"
	"iTask/common"
	"iTask/modules/project/model"
)

func (s *sqlStore) DeleteProject(ctx context.Context, cond map[string]interface{}) error {
	deletedStatus := common.StatusDeleted

	if err := s.db.Table(model.Project{}.TableName()).
		Where(cond).
		Updates(map[string]interface{}{
			"status": deletedStatus,
		}).Error; err != nil {
		return err
	}

	return nil
}

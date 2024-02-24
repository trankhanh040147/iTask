package storage

import (
	"context"
	"iTask/modules/project/model"
)

func (s *sqlStore) CreateProject(ctx context.Context, data *model.ProjectCreation) error {
	if err := s.db.Create(data).Error; err != nil {
		return err
	}

	return nil
}

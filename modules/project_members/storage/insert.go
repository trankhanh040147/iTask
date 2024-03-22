package storage

import (
	"context"
	"iTask/modules/project_members/model"
)

func (store *sqlStore) CreateProjectMember(ctx context.Context, projectId, userId int) error {
	data := &model.ProjectMember{
		UserId:    userId,
		ProjectId: projectId,
	}

	if err := store.db.Create(data).Error; err != nil {
		return err
	}

	return nil
}

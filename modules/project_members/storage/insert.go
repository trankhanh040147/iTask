package storage

import (
	"context"
	"iTask/modules/project_members/model"
)

func (store *sqlStore) CreateProjectMember(ctx context.Context, data *model.ProjectMemberCreation) error {
	//*(data.AddedAt) = time.Now()
	if err := store.db.Create(data).Error; err != nil {
		return err
	}

	return nil
}

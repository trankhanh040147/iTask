package storage

import (
	"context"
	"iTask/common"
	"iTask/modules/project_members/model"
)

func (store *sqlStore) GetProjectIdsByUserId(ctx context.Context, userId int) ([]int, error) {
	var result []int

	if err := store.db.
		Table(model.ProjectMember{}.TableName()).
		Where("user_id = ?", userId).
		Pluck("project_id", &result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}

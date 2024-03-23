package storage

import (
	"context"
	"iTask/modules/task/model"
)

func (sql *sqlStore) CreateTask(ctx context.Context, data *model.TaskCreation) error {
	if err := sql.db.Create(data).Error; err != nil {
		return err
	}

	return nil
}

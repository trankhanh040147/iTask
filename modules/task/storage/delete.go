package storage

import (
	"context"
	"iTask/modules/task/model"
)

func (sql *sqlStore) DeleteTask(ctx context.Context, cond map[string]interface{}) error {
	deletedStatus := int(model.StatusDeleted)

	if err := sql.db.Table(model.Task{}.TableName()).
		Where(cond).
		Updates(map[string]interface{}{"status": deletedStatus}).Error; err != nil {
		return err
	}

	return nil
}

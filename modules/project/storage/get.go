package storage

import (
	"context"
	"errors"
	"iTask/common"
	"iTask/modules/project/model"

	"gorm.io/gorm"
)

func (store *sqlStore) GetProject(ctx context.Context, cond map[string]interface{}, moreKeys ...string) (*model.Project, error) {
	var data *model.Project

	db := store.db.
		Table(model.Project{}.TableName()).
		Where("status <> ?", common.StatusDeleted)

	for _, value := range moreKeys {
		db = db.Preload(value)
	}

	db = db.Preload("Members").Preload("Members.AccountInfo")

	if err := db.Where(cond).First(&data).Error; err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			return nil, common.ErrEntityNotFound(model.EntityName)
		}

		return nil, common.ErrDB(err)
	}

	return data, nil
}

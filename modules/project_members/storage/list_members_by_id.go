package storage

import (
	"golang.org/x/net/context"
	"iTask/common"
	"iTask/modules/project_members/model"
)

func (store *sqlStore) ListMembersById(
	ctx context.Context,
	id int,
	moreKeys ...string,
) ([]model.SimpleMember, error) {
	var result []model.SimpleMember

	db := store.db

	db = db.Preload("AccountInfo")

	if err := db.Select("*").Where("project_id = ?", id).Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}

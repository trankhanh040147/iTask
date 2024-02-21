package storage

import (
	"golang.org/x/net/context"
	"iTask/common"
	"iTask/modules/project/model"
)

func (store *sqlStore) ListProject(
	ctx context.Context,
	filter *model.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]model.Project, error) {
	var result []model.Project

	db := store.db.
		Table(model.Project{}.TableName()).
		Where("status <> ?", model.StatusDeleted)

	// Get items of requester only
	// requester := ctx.Value(common.CurrentUser).(common.Requester)
	// db = db.Where("user_id = ?", requester.GetUserId())

	//if f := filter; f != nil {
	//	if v := f.Status; v != "" {
	//		db = db.Where("status = ?", v)
	//	}
	//}

	if err := db.Select("id").Count(&paging.Total).Error; err != nil {
		return nil, common.ErrorDB(err)
	}

	//for _, value := range moreKeys {
	//	db = db.Preload(value)
	//}

	//if cursor := strings.TrimSpace(paging.FakeCursor); cursor != "" {
	//	id, err := common.UIDFromBase58(cursor)
	//	if err != nil {
	//		return nil, common.ErrorDB(err)
	//	}
	//
	//	db = db.Where("id < ?", id())
	//} else {
	//	db = db.Offset((paging.Page - 1) * paging.Limit)
	//}

	if err := db.
		Select("*").
		Offset((paging.Page - 1) * paging.Limit).
		Order("id desc").
		Limit(paging.Limit).
		Find(&result).Error; err != nil {
		return nil, common.ErrorDB(err)
	}

	//size := len(result)
	//if size > 0 {
	//	paging.NextCursor = result[size-1].SQLModel.Id
	//}

	return result, nil
}

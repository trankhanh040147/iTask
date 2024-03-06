package storage

import (
	"golang.org/x/net/context"
	"iTask/common"
	"iTask/modules/task/model"
	"log"
)

func (store *sqlStore) ListTask(
	ctx context.Context,
	filter *model.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]model.Task, error) {
	var result []model.Task

	db := store.db.
		Table(model.Task{}.TableName()).
		Where("status <> ?", common.StatusDeleted)

	// Get items of requester only
	// requester := ctx.Value(common.CurrentUser).(common.Requester)
	// db = db.Where("user_id = ?", requester.GetUserId())

	log.Println("filter: ", filter)

	//if f := filter; f != nil {
	//	if v := f.Keyword; v != "" {
	//		db = db.Where("name LIKE ? OR description LIKE ?", "%"+v+"%", "%"+v+"%")
	//	}
	//
	//	// if DateRangeFrom = 0, list all projects today
	//	// if DateRangeFrom = 1, list all projects since yesterday
	//	// if DateRangeFrom = 2, list all projects since 2 days ago
	//	today := time.Now()
	//	if v := f.CreatedDayRange; v >= 0 {
	//		dateRange := today.AddDate(0, 0, -v)
	//		db = db.Where("created_at >= ?", dateRange)
	//	}
	//}

	if err := db.Select("id").Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	//for _, value := range moreKeys {
	//	db = db.Preload(value)
	//}

	//db = db.Preload("Members").Preload("Members.AccountInfo")

	//if cursor := strings.TrimSpace(paging.FakeCursor); cursor != "" {
	//	id, err := common.UIDFromBase58(cursor)
	//	if err != nil {
	//		return nil, common.ErrDB(err)
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
		return nil, common.ErrDB(err)
	}

	//size := len(result)
	//if size > 0 {
	//	paging.NextCursor = result[size-1].SQLModel.Id
	//}

	return result, nil
}

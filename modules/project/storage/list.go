package storage

import (
	"golang.org/x/net/context"
	"iTask/common"
	"iTask/modules/project/model"
	"log"
	"time"
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
		Where("status <> ?", common.StatusDeleted)

	// Get items of requester only
	// requester := ctx.Value(common.CurrentUser).(common.Requester)
	// db = db.Where("user_id = ?", requester.GetUserId())

	log.Println("filter: ", filter)

	if f := filter; f != nil {
		if v := f.Keyword; v != "" {
			db = db.Where("name LIKE ? OR description LIKE ?", "%"+v+"%", "%"+v+"%")
		}
		// if DateRangeFrom = 0, list all projects today
		// if DateRangeFrom = 1, list all projects since yesterday
		// if DateRangeFrom = 2, list all projects since 2 days ago

		today := time.Now()
		if v := f.CreatedDateRange; v >= 0 {
			dateRange := today.AddDate(0, 0, -v)
			db = db.Where("created_at >= ?", dateRange)
		}
	}

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

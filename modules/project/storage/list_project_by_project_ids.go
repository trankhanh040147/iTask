package storage

import (
	"golang.org/x/net/context"
	"iTask/common"
	"iTask/modules/project/model"
	"time"
)

func (store *sqlStore) ListProjectByProjectIds(
	ctx context.Context,
	projectIds []int,
	filter *model.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]model.Project, error) {
	var result []model.Project

	db := store.db.
		Table(model.Project{}.TableName()).
		Where("status <> ?", common.StatusDeleted).
		Where("id IN (?)", projectIds)

	if f := filter; f != nil {
		if v := f.Keyword; v != "" {
			db = db.Where("name LIKE ? OR description LIKE ?", "%"+v+"%", "%"+v+"%")
		}

		// if DateRangeFrom = 0, list all projects today
		// if DateRangeFrom = 1, list all projects since yesterday
		// if DateRangeFrom = 2, list all projects since 2 days ago
		today := time.Now()
		if v := f.CreatedDayRange; v >= 0 {
			dateRange := today.AddDate(0, 0, -v)
			db = db.Where("created_at >= ?", dateRange)
		}
	}

	if err := db.Select("id").Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for _, value := range moreKeys {
		db = db.Preload(value)
	}

	db = db.Preload("Members").Preload("Members.AccountInfo")

	// ? testing projectIds

	if err := db.
		Select("*").
		Offset((paging.Page - 1) * paging.Limit).
		Order("id").
		Limit(paging.Limit).
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}

package storage

import (
	"context"
	"social-todo-list/common"
	"social-todo-list/module/userlikeitem/model"
	"strings"
	"time"

	"github.com/btcsuite/btcutil/base58"
)

const (
	timeLayout   = "2006-01-02T15:04:05.999999"
	timeDBLayout = "2006-01-02 15:04:05.999999"
)

func (store *sqlStore) ListUsers(ctx context.Context, itemId int, paging *common.Paging) ([]common.SimpleUser, error) {
	var result []model.Like
	db := store.db.Table(model.Like{}.TableName()).Where("item_id = ?", itemId)

	if err := db.Select("user_id").Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if cursor := strings.TrimSpace(paging.FakeCursor); cursor != "" {
		createdTime, err := time.Parse(timeLayout, string(base58.Decode(cursor)))
		if err != nil {
			return nil, common.ErrDB(err)
		}

		db = db.Where("created_at < ?", createdTime.Format(timeDBLayout))
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}

	if err := db.Select("*").
		Order("created_at desc").
		Limit(paging.Limit).
		Preload("User").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	size := len(result)
	users := make([]common.SimpleUser, size)
	for index := range users {
		users[index] = *result[index].User
		users[index].UpdatedAt = nil
		users[index].CreatedAt = result[index].CreatedAt
	}

	if size > 0 {
		paging.NextCursor = base58.Encode([]byte(users[size-1].CreatedAt.Format(timeLayout)))
	}

	return users, nil
}

func (store *sqlStore) GetItemLikes(ctx context.Context, ids []int) (map[int]int, error) {
	result := make(map[int]int)

	type sqlData struct {
		ItemId int `gorm:"column:item_id"`
		Count  int `gorm:"column:count"`
	}

	var likes []sqlData
	if err := store.db.Table(model.Like{}.TableName()).
		Select("item_id, Count(item_id) as `count`").
		Where("item_id in (?)", ids).
		Group("item_id").
		Find(&likes).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for _, item := range likes {
		result[item.ItemId] = item.Count
	}

	return result, nil
}

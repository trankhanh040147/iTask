package biz

import (
	"context"
	"social-todo-list/common"
	"social-todo-list/module/user/model"
)

type ListUsersLikedItemStorage interface {
	ListUsers(ctx context.Context, itemId int, paging *common.Paging) ([]common.SimpleUser, error)
}

type listUsersLikedItemBiz struct {
	store ListUsersLikedItemStorage
}

func NewListUsersLikedItemBiz(store ListUsersLikedItemStorage) *listUsersLikedItemBiz {
	return &listUsersLikedItemBiz{store: store}
}

func (biz *listUsersLikedItemBiz) ListUsersLikedItem(ctx context.Context, itemId int, paging *common.Paging) ([]common.SimpleUser, error) {
	result, err := biz.store.ListUsers(ctx, itemId, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity(model.EntityName, err)
	}

	return result, nil
}

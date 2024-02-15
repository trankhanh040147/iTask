package biz

import (
	"context"
	"log"
	"social-todo-list/common"
	"social-todo-list/module/userlikeitem/model"
)

type UserUnlikeItemStorage interface {
	Find(ctx context.Context, userId, itemId int) (*model.Like, error)
	Delete(ctx context.Context, userId, itemId int) error
}

type DecreaseLikedCountStorage interface {
	DecreaseLikedCount(ctx context.Context, id int) error
}

type userUnlikeItemBiz struct {
	store     UserUnlikeItemStorage
	itemStore DecreaseLikedCountStorage
}

func NewUserUnlikeItemBiz(store UserUnlikeItemStorage, itemStore DecreaseLikedCountStorage) *userUnlikeItemBiz {
	return &userUnlikeItemBiz{store: store, itemStore: itemStore}
}

func (biz *userUnlikeItemBiz) UnlikeItem(ctx context.Context, userId, itemId int) error {
	_, err := biz.store.Find(ctx, userId, itemId)
	if err == common.RecordNotFound {
		return model.ErrDidNotLikeItem(err)
	}

	if err != nil {
		return model.ErrCannotUnlikeItem(err)
	}

	if err := biz.store.Delete(ctx, userId, itemId); err != nil {
		return model.ErrCannotUnlikeItem(err)
	}

	go func() {
		defer common.Recovery()

		if err := biz.itemStore.DecreaseLikedCount(ctx, itemId); err != nil {
			log.Println(err)
		}
	}()

	return nil
}

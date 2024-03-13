package biz

import (
	"context"
	"iTask/common"
	"iTask/modules/tag/model"
)

type ListTagStorage interface {
	ListTag(ctx context.Context, filter *model.Filter, paging *common.Paging, moreKeys ...string) ([]model.Tag, error)
}

type listTagBiz struct {
	store ListTagStorage
}

func NewListTagBiz(store ListTagStorage) *listTagBiz {
	return &listTagBiz{store: store}
}

func (biz *listTagBiz) ListTag(ctx context.Context, filter *model.Filter, paging *common.Paging) ([]model.Tag, error) {
	data, err := biz.store.ListTag(ctx, filter, paging)

	if err != nil {
		return nil, common.ErrCannotListEntity(model.EntityName, err)
	}

	return data, nil
}

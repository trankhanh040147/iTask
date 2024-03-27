package biz

import (
	"context"
	"iTask/modules/task_assignees/model"
)

type ListAssigneeStorage interface {
	ListAssignee(ctx context.Context, filter *model.Filter, moreKeys ...string) ([]model.TaskAssignee, error)
}

type listAssigneeBiz struct {
	store ListAssigneeStorage
}

func NewListAssigneeBiz(store ListAssigneeStorage) *listAssigneeBiz {
	return &listAssigneeBiz{store: store}
}

func (biz *listAssigneeBiz) ListAssignee(ctx context.Context, filter *model.Filter) ([]model.TaskAssignee, error) {
	// task_id is required, and will be stored in filter
	data, err := biz.store.ListAssignee(ctx, filter, "UserInfo")

	if err != nil {
		return nil, err
	}

	return data, nil
}

package biz

import (
	"context"
	"iTask/common"
	"iTask/modules/task/model"
)

type DeleteTaskStorage interface {
	GetTask(ctx context.Context, cond map[string]interface{}, moreKeys ...string) (*model.Task, error)
	DeleteTask(ctx context.Context, cond map[string]interface{}) error
}

type deleteTaskBiz struct {
	store DeleteTaskStorage
}

func NewDeleteTaskBiz(store DeleteTaskStorage) *deleteTaskBiz {
	return &deleteTaskBiz{store: store}
}

func (biz *deleteTaskBiz) DeleteTaskById(ctx context.Context, id int) error {
	data, err := biz.store.GetTask(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return common.ErrCannotGetEntity(model.EntityName, err)
	}

	if data.Status == model.StatusDeleted {
		return common.ErrEntityDeleted(model.EntityName)
	}

	if err := biz.store.DeleteTask(ctx, map[string]interface{}{"id": id}); err != nil {
		return common.ErrCannotDeleteEntity(model.EntityName, err)
	}

	return nil
}

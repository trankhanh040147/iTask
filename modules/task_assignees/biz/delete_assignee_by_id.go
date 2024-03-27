package biz

import (
	"context"
	"iTask/common"
	model "iTask/modules/task_assignees/model"
)

type DeleteAssigneeStorage interface {
	GetAssignee(ctx context.Context, cond map[string]interface{}, moreKeys ...string) (*model.TaskAssignee, error)
	DeleteAssignee(ctx context.Context, cond map[string]interface{}) error
}

type deleteAssigneeBiz struct {
	store DeleteAssigneeStorage
}

func NewDeleteAssigneeBiz(store DeleteAssigneeStorage) *deleteAssigneeBiz {
	return &deleteAssigneeBiz{store: store}
}

func (biz *deleteAssigneeBiz) DeleteAssignee(ctx context.Context, userId, taskId int) error {
	_, err := biz.store.GetAssignee(ctx, map[string]interface{}{"user_id": userId, "task_id": taskId})

	if err != nil {
		return common.ErrCannotGetEntity(model.EntityName, err)
	}

	if err := biz.store.DeleteAssignee(ctx, map[string]interface{}{"user_id": userId, "task_id": taskId}); err != nil {
		return common.ErrCannotDeleteEntity(model.EntityName, err)
	}

	return nil
}

package biz

import (
	"context"
	"iTask/common"
	"iTask/modules/task_assignees/model"
)

type CreateAssigneeStorage interface {
	CreateAssignee(ctx context.Context, data *model.TaskAssigneeCreation) error
}

type createAssigneeBiz struct {
	store     CreateAssigneeStorage
	requester common.Requester
}

func NewCreateAssigneeBiz(store CreateAssigneeStorage, requester common.Requester) *createAssigneeBiz {
	return &createAssigneeBiz{
		store:     store,
		requester: requester,
	}
}

func (biz *createAssigneeBiz) CreateNewAssignee(ctx context.Context, data *model.TaskAssigneeCreation) error {
	if err := data.Validate(); err != nil {
		return common.ErrValidation(err)
	}

	if err := biz.store.CreateAssignee(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(model.EntityName, err)
	}

	return nil
}

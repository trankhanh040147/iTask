package biz

import (
	"context"
	"iTask/common"
	"iTask/modules/task_comments/model"
)

type CreateTaskCommentStorage interface {
	CreateTaskComment(ctx context.Context, data *model.TaskCommentCreation) error
}

type createTaskCommentBiz struct {
	store     CreateTaskCommentStorage
	requester common.Requester
}

func NewCreateTaskCommentBiz(store CreateTaskCommentStorage, requester common.Requester) *createTaskCommentBiz {
	return &createTaskCommentBiz{
		store:     store,
		requester: requester,
	}
}

func (biz *createTaskCommentBiz) CreateNewTaskComment(ctx context.Context, data *model.TaskCommentCreation) error {
	if err := data.Validate(); err != nil {
		return common.ErrValidation(err)
	}

	if err := biz.store.CreateTaskComment(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(model.EntityName, err)
	}

	return nil
}

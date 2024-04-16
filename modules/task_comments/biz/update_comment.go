package biz

import (
	"context"
	"iTask/common"
	"iTask/modules/task_comments/model"
)

type UpdateCommentStorage interface {
	UpdateComment(ctx context.Context, cond map[string]interface{}, dataUpdate *model.TaskCommentUpdate) error
	GetComment(ctx context.Context, cond map[string]interface{}) (*model.TaskComment, error)
}

type updateCommentBiz struct {
	store UpdateCommentStorage
}

func NewUpdateCommentBiz(store UpdateCommentStorage) *updateCommentBiz {
	return &updateCommentBiz{store: store}
}

func (biz *updateCommentBiz) UpdateComment(ctx context.Context, taskId int, dataUpdate *model.TaskCommentUpdate) error {
	// * check existence
	if _, err := biz.store.GetComment(ctx, map[string]interface{}{"id": taskId}); err != nil {
		//return common.ErrEntityNotFound(model.EntityName)
		return common.ErrCannotGetEntity(model.EntityName, err)
	}

	// * check validation
	if err := dataUpdate.Validate(); err != nil {
		return common.ErrValidation(err)
	}

	// todo: check ErrNoPermission

	if err := biz.store.UpdateComment(ctx, map[string]interface{}{"id": taskId}, dataUpdate); err != nil {
		return common.ErrCannotUpdateEntity(model.EntityName, err)
	}

	return nil
}

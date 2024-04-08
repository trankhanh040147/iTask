package biz

import (
	"context"
	"iTask/common"
	"iTask/modules/task_comments/model"
)

type ListTaskCommentStorage interface {
	ListComments(ctx context.Context, cond map[string]interface{}, paging *common.Paging, moreKeys ...string) ([]model.TaskComments, error)
}

type listTaskCommentBiz struct {
	store ListTaskCommentStorage
}

func NewListTaskCommentBiz(store ListTaskCommentStorage) *listTaskCommentBiz {
	return &listTaskCommentBiz{store: store}
}

func (biz *listTaskCommentBiz) ListTaskCommentsByTaskId(ctx context.Context, taskId int, paging *common.Paging) ([]model.TaskComments, error) {

	result, err := biz.store.ListComments(ctx, map[string]interface{}{"id": taskId}, paging)
	if err != nil {
		return nil, err
	}

	return result, nil

}

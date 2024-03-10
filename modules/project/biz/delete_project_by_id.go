package biz

import (
	"context"
	"iTask/common"
	"iTask/modules/project/model"
)

type DeleteProjectStorage interface {
	GetProject(ctx context.Context, cond map[string]interface{}, moreKeys ...string) (*model.Project, error)
	DeleteProject(ctx context.Context, cond map[string]interface{}) error
}

type deleteProjectBiz struct {
	store DeleteProjectStorage
}

func NewDeleteProjectBiz(store DeleteProjectStorage) *deleteProjectBiz {
	return &deleteProjectBiz{store: store}
}

func (biz *deleteProjectBiz) DeleteProjectById(ctx context.Context, id int) error {
	data, err := biz.store.GetProject(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return common.ErrCannotGetEntity(model.EntityName, err)
	}

	if data.Status.String() == "Deleted" {
		return common.ErrEntityDeleted(model.EntityName)
	}

	if err := biz.store.DeleteProject(ctx, map[string]interface{}{"id": id}); err != nil {
		return common.ErrCannotDeleteEntity(model.EntityName, err)
	}

	return nil
}

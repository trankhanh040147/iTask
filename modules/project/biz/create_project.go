package biz

import (
	"context"
	"iTask/common"
	"iTask/modules/project/model"
)

type CreateProjectStorage interface {
	CreateProject(ctx context.Context, data *model.ProjectCreation) error
}

type createProjectBiz struct {
	store CreateProjectStorage
}

func NewCreateProjectBiz(store CreateProjectStorage) *createProjectBiz {
	return &createProjectBiz{store: store}
}

func (biz *createProjectBiz) CreateNewProject(ctx context.Context, data *model.ProjectCreation) error {
	if err := data.Validate(); err != nil {
		return common.ErrValidation(err)
	}

	if err := biz.store.CreateProject(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(model.EntityName, err)
	}

	return nil
}

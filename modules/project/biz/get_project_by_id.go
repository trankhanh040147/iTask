package biz

import (
	"context"
	"iTask/common"
	"iTask/modules/project/model"
)

type GetProjectStorage interface {
	GetProject(ctx context.Context, cond map[string]interface{}) (*model.Project, error)
}

type getProjectBiz struct {
	store GetProjectStorage
}

func NewGetProjectBiz(store GetProjectStorage) *getProjectBiz {
	return &getProjectBiz{store: store}
}

func (biz *getProjectBiz) GetProjectById(ctx context.Context, id int) (*model.Project, error) {
	data, err := biz.store.GetProject(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return nil, common.ErrCannotGetEntity(model.EntityName, err)
	}

	return data, nil
}

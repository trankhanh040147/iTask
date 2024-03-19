package biz

import (
	"context"
	"iTask/common"
	"iTask/modules/project/model"
)

type GetProjectStorage interface {
	GetProject(ctx context.Context, cond map[string]interface{}, moreKeys ...string) (*model.Project, error)
}

type getProjectBiz struct {
	store             GetProjectStorage
	ProjectTagStorage ProjectTagStorage
}

func NewGetProjectBiz(store GetProjectStorage, projectTagStorage ProjectTagStorage) *getProjectBiz {
	return &getProjectBiz{store: store, ProjectTagStorage: projectTagStorage}
}

func (biz *getProjectBiz) GetProjectById(ctx context.Context, id int) (*model.Project, error) {
	data, err := biz.store.GetProject(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return nil, common.ErrCannotGetEntity(model.EntityName, err)
	}

	cond := map[string]interface{}{"project_id": id}

	// get tags
	ProjectTagsMap, err := biz.ProjectTagStorage.GetProjectTagsByProjectId(ctx, cond)
	if err != nil {
		return data, nil
	}

	data.Tags = ProjectTagsMap[data.Id]

	return data, nil
}
